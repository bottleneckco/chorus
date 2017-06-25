package web

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"youtube"

	"ffmpeg"

	"path"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func createChannel(c *gin.Context) {
	var payload createChannelPayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, response{
			Status: statusError,
			Error:  "Invalid payload",
		})
		return
	}

	users := make(map[int]User)
	newUserID := 1
	createdByUser := User{
		ID:       newUserID,
		Nickname: payload.CreatedBy,
	}

	users[newUserID] = createdByUser

	channel := Channel{
		ID:                generateAccessCode(),
		CreatedBy:         newUserID,
		Name:              payload.Name,
		Description:       payload.Description,
		Users:             users,
		VideoResultsCache: make(map[string]youtube.YoutubeVideo),
		Queue:             make([]youtube.YoutubeVideo, 0),
		SkipCurrent:       false,
	}

	setUserCookie(createdByUser, c)
	channelMap[channel.ID] = &channel

	// Populate usersArr for view
	channel.UsersArray = formatUsersForJson(users)

	c.JSON(http.StatusOK, channelResponse{
		response{Status: statusOK},
		channel,
	})

	// Channel Manager
	go func() {
		log.Printf("Channel manager started for Channel '%s'\n", channel.Name)
		numUsers := 1
		for numUsers != 0 || len(channel.Queue) > 0 {
			if len(channel.Queue) == 0 {
				time.Sleep(time.Second * 2)
				channel.CheckUsersStillAlive()
				continue
			}

			result := channel.Queue[0]

			log.Printf("Downloading file '%s' via URL '%s'\n", result.Fulltitle, result.Formats[0].URL)

			// Download file
			downloadFile, err := ioutil.TempFile(os.TempDir(), "audio")
			if err != nil {
				log.Println(err)
				continue
			}
			resp, err := http.Get(result.Formats[0].URL)
			if err != nil {
				log.Println(err)
				continue
			}

			defer resp.Body.Close()
			io.Copy(downloadFile, resp.Body)

			log.Printf("Segmenting '%s'\n", downloadFile.Name())

			encode, err := ffmpeg.Segment(downloadFile.Name())
			if err != nil {
				log.Println(err)
				continue
			}

			for _, segmentFileName := range encode.SegmentFileNames {
				log.Printf("Feeding segment '%s'\n", segmentFileName)

				channel.CheckUsersStillAlive()

				data, err := ioutil.ReadFile(path.Join(encode.ContainerDir, segmentFileName))
				if err != nil {
					log.Println(err)
					continue
				}

				// Distribute
				channel.BroadcastMessage(websocket.BinaryMessage, data)

				// Check if we should abort distribution of the current song
				if channel.SkipCurrent {
					channel.SkipCurrent = false
					jsonData, _ := json.Marshal(websocketCommand{
						Command: commandSkipCurrent,
					})
					channel.BroadcastMessage(websocket.TextMessage, jsonData)
					break
				}

				time.Sleep(time.Millisecond * 4000)
			}

			numUsers = len(channel.Users)
			if len(channel.Queue) == 1 {
				channel.Queue = make([]youtube.YoutubeVideo, 0)
			} else {
				channel.Queue = channel.Queue[1:]
			}

			log.Println("Job complete")
			os.Remove(encode.ContainerDir)
		}
	}()
}

func getChannel(c *gin.Context) {
	channel, isChannelExists := channelMap[c.Param("id")]
	if !isChannelExists {
		c.JSON(http.StatusInternalServerError, response{
			Status: statusError,
			Error:  "Channel does not exist",
		})
		return
	}

	channel.UsersArray = formatUsersForJson(channel.Users)

	c.JSON(http.StatusOK, channelResponse{
		response: response{Status: statusOK},
		Channel:  *channel,
	})
}

func addUserToChannel(c *gin.Context) {
	var json createUserPayload
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, response{
			Status: statusError,
			Error:  "Invalid payload",
		})
		return
	}

	channel, isChannelExists := channelMap[c.Param("id")]
	if !isChannelExists {
		c.JSON(http.StatusInternalServerError, response{
			Status: statusError,
			Error:  "Channel does not exist",
		})
		return
	}

	users := channel.Users

	newUserID := len(users) + 1
	newUser := User{
		ID:       newUserID,
		Nickname: json.Nickname,
	}

	users[newUserID] = newUser

	channel.Users = users
	channelMap[c.Param("id")] = channel
	setUserCookie(newUser, c)

	// Populate usersArr for view
	channel.UsersArray = formatUsersForJson(users)

	c.JSON(http.StatusOK, channelAddUserResponse{
		response: response{Status: statusOK},
		User:     newUser,
	})
}

func getUsersInChannel(c *gin.Context) {
	channel, isChannelExists := channelMap[c.Param("id")]
	if !isChannelExists {
		c.JSON(http.StatusInternalServerError, response{
			Status: statusError,
			Error:  "Channel does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, channelListUsersResponse{
		response: response{Status: statusOK},
		Users:    formatUsersForJson(channel.Users),
	})
}

func getChannelQueue(c *gin.Context) {
	channel, isChannelExists := channelMap[c.Param("id")]
	if !isChannelExists {
		c.JSON(http.StatusInternalServerError, response{
			Status: statusError,
			Error:  "Channel does not exist",
		})
		return
	}

	var jsonArray = make([]videoResult, 0)
	for _, result := range channel.Queue {
		jsonArray = append(jsonArray, videoResult{
			Name:         result.Fulltitle,
			ThumbnailURL: result.Thumbnail,
			Duration:     result.Duration,
			URL:          result.WebpageURL,
		})
	}

	c.JSON(http.StatusOK, channelListQueueResponse{
		response: response{Status: statusOK},
		Count:    len(jsonArray),
		Queue:    jsonArray,
	})
}

func addToChannelQueue(c *gin.Context) {
	var payload addToChannelQueuePayload
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, response{
			Status: statusError,
			Error:  "Invalid payload",
		})
		return
	}

	channel, isChannelExists := channelMap[c.Param("id")]
	if !isChannelExists {
		c.JSON(http.StatusInternalServerError, response{
			Status: statusError,
			Error:  "Channel does not exist",
		})
		return
	}

	channel.Queue = append(channel.Queue, channel.VideoResultsCache[payload.URL])
	channelMap[c.Param("id")] = channel

	c.JSON(http.StatusOK, response{Status: statusOK})
}

func skipInChannelQueue(c *gin.Context) {
	channel, isChannelExists := channelMap[c.Param("id")]
	if !isChannelExists {
		c.JSON(http.StatusInternalServerError, response{
			Status: statusError,
			Error:  "Channel does not exist",
		})
		return
	}

	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil || len(indexStr) == 0 || index < 0 || index > len(channel.Queue)-1 {
		c.JSON(http.StatusNotAcceptable, response{
			Status: statusError,
			Error:  "Invalid index",
		})
		return
	}

	// I know
	if len(channel.Queue) > 1 && index != 0 {
		channel.Queue = append(channel.Queue[:index], channel.Queue[index+1:]...)
	} else {
		channel.SkipCurrent = true
	}

	c.JSON(http.StatusOK, response{Status: statusOK})
}
