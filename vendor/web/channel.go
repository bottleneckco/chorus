package web

import (
	"errors"
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
)

func createChannel(c *gin.Context) {
	var json createChannelPayload
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, response{
			Status: statusError,
			Error:  errors.New("Invalid payload"),
		})
		return
	}

	users := make(map[int]User)
	newUserID := 1
	createdByUser := User{
		ID:       newUserID,
		Nickname: json.CreatedBy,
	}

	users[newUserID] = createdByUser

	channel := Channel{
		ID:                getNextChannelID(),
		CreatedBy:         newUserID,
		Name:              json.Name,
		Description:       json.Description,
		AccessCode:        generateAccessCode(),
		Users:             users,
		Stream:            make(chan []byte, 10),
		VideoResultsCache: make(map[string]youtube.YoutubeVideo),
		Queue:             make([]youtube.YoutubeVideo, 0),
	}

	setUserCookie(newUserID, c)
	Channels[getNextChannelID()] = channel

	// Populate usersArr for view
	channel.UsersArray = formatUsersForJson(users)

	c.JSON(http.StatusOK, channelCreateResponse{
		response{Status: statusOK},
		channel,
	})

	// Channel Manager
	go func() {
		log.Printf("Channel manager started for Channel '%s'\n", channel.Name)
		numUsers := 1
		for numUsers != 0 {
			channel = Channels[channel.ID]
			if len(channel.Queue) == 0 {
				time.Sleep(time.Second * 2)
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
				data, err := ioutil.ReadFile(path.Join(encode.ContainerDir, segmentFileName))
				if err != nil {
					log.Println(err)
					continue
				}
				channel.Stream <- data
				time.Sleep(time.Millisecond * 2000)
			}

			channel = Channels[channel.ID]
			numUsers = len(channel.Users)
			if len(channel.Queue) == 1 {
				channel.Queue = make([]youtube.YoutubeVideo, 0)
			} else {
				channel.Queue = channel.Queue[1:]
			}
			Channels[channel.ID] = channel

			log.Println("Job complete")
			os.Remove(encode.ContainerDir)
		}
	}()

}

func getChannelIDFromParam(c *gin.Context) ChannelID {
	id, _ := strconv.Atoi(c.Param("id"))
	return ChannelID(id)
}

func getChannel(c *gin.Context) {
	channelID := getChannelIDFromParam(c)
	channel := Channels[channelID]
	channel.UsersArray = formatUsersForJson(channel.Users)

	c.JSON(http.StatusOK, channel)
}

func addUserToChannel(c *gin.Context) {
	channelID := getChannelIDFromParam(c)

	var json createUserPayload
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, response{
			Status: statusError,
			Error:  errors.New("Invalid payload"),
		})
		return
	}

	channel := Channels[channelID]
	users := channel.Users

	newUserID := len(users) + 1
	newUser := User{
		ID:       newUserID,
		Nickname: json.Nickname,
	}

	users[newUserID] = newUser

	channel.Users = users
	Channels[channelID] = channel

	// Populate usersArr for view
	channel.UsersArray = formatUsersForJson(users)

	c.JSON(http.StatusOK, response{Status: statusOK})
}

func getUsersInChannel(c *gin.Context) {
	channelID := getChannelIDFromParam(c)
	c.JSON(http.StatusOK, channelListUsersResponse{
		response: response{Status: statusOK},
		Users:    formatUsersForJson(Channels[channelID].Users),
	})
}

func getChannelQueue(c *gin.Context) {
	channelID := getChannelIDFromParam(c)
	channel, isChannelExists := Channels[channelID]
	if !isChannelExists {
		c.JSON(http.StatusInternalServerError, response{
			Status: statusError,
			Error:  errors.New("Channel does not exist"),
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
			Error:  errors.New("Invalid payload"),
		})
		return
	}

	channelID := getChannelIDFromParam(c)
	channel, isChannelExists := Channels[channelID]
	if !isChannelExists {
		c.JSON(http.StatusInternalServerError, response{
			Status: statusError,
			Error:  errors.New("Channel does not exist"),
		})
		return
	}

	channel.Queue = append(channel.Queue, channel.VideoResultsCache[payload.URL])
	Channels[channelID] = channel

	c.JSON(http.StatusOK, response{Status: statusOK})
}

func skipInChannelQueue(c *gin.Context) {
	channelID := getChannelIDFromParam(c)
	channel, isChannelExists := Channels[channelID]
	if !isChannelExists {
		c.JSON(http.StatusInternalServerError, response{
			Status: statusError,
			Error:  errors.New("Channel does not exist"),
		})
		return
	}

	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil || len(indexStr) == 0 || index < 0 || index > len(channel.Queue)-1 {
		c.JSON(http.StatusNotAcceptable, response{
			Status: statusError,
			Error:  errors.New("Invalid index"),
		})
		return
	}

	// I know
	channel.Queue = append(channel.Queue[:index], channel.Queue[index+1:]...)
	Channels[channelID] = channel

	c.JSON(http.StatusOK, response{Status: statusOK})
}
