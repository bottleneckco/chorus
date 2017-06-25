package web

import (
	"errors"
	"log"
	"net/http"
	"youtube"

	"github.com/gin-gonic/gin"
)

type videoResult struct {
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnail"`
	Duration     int    `json:"duration"`
	URL          string `json:"url"`
}

func searchMusic(c *gin.Context) {
	query := c.Query("term")

	results, err := youtube.Search(query, 10)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response{
			Status: statusError,
			Error:  errors.New("Error with service provider"),
		})
		return
	}

	channel, isChannelExists := channelMap[c.Param("id")]
	if !isChannelExists {
		c.JSON(http.StatusInternalServerError, response{
			Status: statusError,
			Error:  errors.New("Channel does not exist"),
		})
		return
	}

	// Add to the cache
	var jsonArray []videoResult
	for _, result := range results {
		jsonArray = append(jsonArray, videoResult{
			Name:         result.Fulltitle,
			ThumbnailURL: result.Thumbnail,
			Duration:     result.Duration,
			URL:          result.WebpageURL,
		})
		channel.VideoResultsCache[result.WebpageURL] = result
	}

	c.JSON(http.StatusOK, searchResponse{
		response: response{Status: statusOK},
		Count:    len(jsonArray),
		Results:  jsonArray,
	})
}
