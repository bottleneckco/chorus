package web

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"model"
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

	users := make(map[int]model.User)
	newUserID := 1
	createdByUser := model.User{
		ID:       newUserID,
		Nickname: payload.CreatedBy,
	}

	users[newUserID] = createdByUser

	channel := model.Channel{
		ID:          generateAccessCode(),
		CreatedBy:   newUserID,
		Name:        payload.Name,
		Description: payload.Description,
		I: model.I{
			Users:             users,
			VideoResultsCache: make(map[string]youtube.YoutubeVideo),
			Queue:             make([]youtube.YoutubeVideo, 0),
			SkipCurrent:       false,
		},
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
		for numUsers != 0 || len(channel.I.Queue) > 0 {
			if len(channel.I.Queue) == 0 {
				time.Sleep(time.Second * 2)
				usersWhoLeft := channel.CheckUsersStillAlive()
				if len(usersWhoLeft) > 0 {
					jsonData, _ := json.Marshal(websocketDataCommand{
						websocketCommand{
							Command: commandUsersLeft,
						},
						usersWhoLeft,
					})
					channel.BroadcastMessage(websocket.TextMessage, jsonData)
				}
				continue
			}

			result := channel.I.Queue[0]

			log.Printf("Downloading file '%s' via URL '%s'\n", queueItem.Video.Fulltitle, queueItem.Video.Formats[0].URL)

			// Download file
			downloadFile, err := ioutil.TempFile(os.TempDir(), "audio")
			if err != nil {
				log.Println(err)
				continue
			}
			resp, err := http.Get(queueItem.Video.Formats[0].URL)
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
				if channel.I.SkipCurrent {
					channel.I.SkipCurrent = false
					jsonData, _ := json.Marshal(websocketCommand{
						Command: commandSkipCurrent,
					})
					channel.BroadcastMessage(websocket.TextMessage, jsonData)
					break
				}

				time.Sleep(time.Millisecond * 4000)
			}

			numUsers = len(channel.I.Users)
			if len(channel.I.Queue) == 1 {
				channel.I.Queue = make([]youtube.YoutubeVideo, 0)
			} else {
				channel.I.Queue = channel.I.Queue[1:]
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

	channel.UsersArray = formatUsersForJson(channel.I.Users)

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

	users := channel.I.Users

	newUserID := len(users) + 1
	newUser := model.User{
		ID:       newUserID,
		Nickname: json.Nickname,
	}

	users[newUserID] = newUser

	channel.I.Users = users
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
		Users:    formatUsersForJson(channel.I.Users),
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
	for _, result := range channel.I.Queue {
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

	user, err := getUserByCookies(c)
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

	channel.I.Queue = append(channel.I.Queue, channel.I.VideoResultsCache[payload.URL])
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
	if err != nil || len(indexStr) == 0 || index < 0 || index > len(channel.I.Queue)-1 {
		c.JSON(http.StatusNotAcceptable, response{
			Status: statusError,
			Error:  "Invalid index",
		})
		return
	}

	// I know
	if len(channel.I.Queue) > 1 && index != 0 {
		channel.I.Queue = append(channel.I.Queue[:index], channel.I.Queue[index+1:]...)
	} else {
		channel.I.SkipCurrent = true
	}

	c.JSON(http.StatusOK, response{Status: statusOK})
}
