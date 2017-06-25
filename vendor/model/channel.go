package model

import (
	"ffmpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
	"youtube"

	"encoding/json"

	"github.com/gorilla/websocket"
)

const (
	ChannelCommandPause = iota
	ChannelCommandResume
	ChannelCommandSkip
)

const (
	ChannelStatePlaying = iota
	ChannelStatePaused
)

type Channel struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CreatedBy   int     `json:"created_by"`
	I           I       `json:"-"`
	UsersArray  []*User `json:"users"` // UsersArray is only for display
}

type BroadcastItem struct {
	MessageType int
	Data        []byte
}

type BroadcastResult struct {
	User  *User
	Error error
}

// I is the internal variable struct
type I struct {
	CommandStream       chan int                        `json:"-"`
	DataBroadcastStream chan BroadcastItem              `json:"-"`
	State               int                             `json:"-"`
	VideoResultsCache   map[string]youtube.YoutubeVideo `json:"-"`
	Queue               *Queue                          `json:"-"`
	Users               map[int]*User                   `json:"-"`
}

func NewChannel(ID string, creator User, name, description string) Channel {
	return Channel{
		ID:          ID,
		CreatedBy:   creator.ID,
		Name:        name,
		Description: description,
		I: I{
			CommandStream:       make(chan int, 10),
			DataBroadcastStream: make(chan BroadcastItem, 100),
			Users:               map[int]*User{creator.ID: &creator},
			State:               ChannelStatePlaying,
			VideoResultsCache:   make(map[string]youtube.YoutubeVideo),
			Queue:               &Queue{Q: make([]QueueItem, 0)},
		},
	}
}

func (c *Channel) BroadcastMessage(messageType int, data []byte) []BroadcastResult {
	var errors = make([]BroadcastResult, 0)
	var err error
	for _, user := range c.I.Users {
		if user.WSConn == nil || user.WSConn.Conn == nil {
			continue
		}
		err = user.WSConn.WriteMessage(messageType, data)
		if err != nil {
			errors = append(errors, BroadcastResult{
				user,
				err,
			})
		}
	}
	return errors
}

func (c *Channel) BroadcastTextMessage(data []byte) {
	c.I.DataBroadcastStream <- BroadcastItem{
		websocket.TextMessage,
		data,
	}
}

func (c *Channel) BroadcastBinaryMessage(data []byte) {
	c.I.DataBroadcastStream <- BroadcastItem{
		websocket.BinaryMessage,
		data,
	}
}

func (c *Channel) BroadcastJSON(data interface{}) {
	jsonData, _ := json.Marshal(data)
	c.I.DataBroadcastStream <- BroadcastItem{
		websocket.TextMessage,
		jsonData,
	}
}

func (c *Channel) IssueCommand(command int) {
	c.I.CommandStream <- command
}

func (c *Channel) BroadcastManager() {
	for broadcastItem := range c.I.DataBroadcastStream {
		results := c.BroadcastMessage(broadcastItem.MessageType, broadcastItem.Data)
		for _, result := range results {
			log.Println(result.Error)
			log.Printf("[CM 2.0] '%s' left\n", result.User.Nickname)
			c.RemoveUser(result.User.ID)

			// Broadcast
			jsonData, _ := json.Marshal(map[string]interface{}{
				"command": userLeft,
				"user":    *result.User,
			})
			c.BroadcastTextMessage(jsonData)
		}
	}
}

// Manager the manager loop
func (c *Channel) Manager() {
	log.Printf("[CM 2.0] started for Channel '%s'\n", c.Name)

	// Broadcast run
	go c.BroadcastManager()

	// Periodically ping all users
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			jsonData, _ := json.Marshal(map[string]string{"command": ping})
			c.BroadcastTextMessage(jsonData)
		}
	}()

	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		if c.I.Queue.Count() == 0 || len(c.I.Users) == 0 || c.I.State == ChannelStatePaused {
			continue
		}

		// No command
		queueItem := c.I.Queue.First()
		result := queueItem.Video

		log.Printf("[CM 2.0] Downloading file '%s' via URL '%s'\n", result.Fulltitle, result.Formats[0].URL)

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

		log.Printf("[CM 2.0] Segmenting '%s'\n", downloadFile.Name())

		encoded, err := ffmpeg.Segment(downloadFile.Name())
		if err != nil {
			log.Println(err)
			continue
		}

	segmentLoop:
		for _, segmentFileName := range encoded.SegmentFileNames {
			log.Printf("[CM 2.0] Feeding segment '%s'\n", segmentFileName)

			data, err := ioutil.ReadFile(path.Join(encoded.ContainerDir, segmentFileName))
			if err != nil {
				log.Println(err)
				continue
			}

			// Distribute
			c.I.DataBroadcastStream <- BroadcastItem{
				MessageType: websocket.BinaryMessage,
				Data:        data,
			}

			select {
			case cmd := <-c.I.CommandStream:
				// There is a command
				switch cmd {
				case ChannelCommandPause:
					log.Printf("[CM 2.0] Paused\n")
					c.I.State = ChannelStatePaused
					continue
				case ChannelCommandResume:
					log.Printf("[CM 2.0] Resume\n")
					c.I.State = ChannelStatePlaying
					continue
				case ChannelCommandSkip:
					log.Printf("[CM 2.0] Skip requested\n")
					os.RemoveAll(encoded.ContainerDir)
					break segmentLoop
				}
				break
			default:
				break
			}

			time.Sleep(time.Millisecond * 4000)
		}

		// Remove the item from the queue
		c.I.Queue.PopFirst()

		// Check if channel died
		if c.I.Queue.Count() == 0 && len(c.I.Users) == 0 {
			close(c.I.CommandStream)
			close(c.I.DataBroadcastStream)
			ticker.Stop()
		}
	}
}

func (c Channel) AddUser(user *User) {
	log.Printf("[CM 2.0] '%s' joined\n", user.Nickname)
	c.I.Users[user.ID] = user
}

func (c Channel) RemoveUser(id int) {
	delete(c.I.Users, id)
}
