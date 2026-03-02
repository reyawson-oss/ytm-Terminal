package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/youtube/v3"
)


type ID struct {
    Kind			string`json:"kind"`
    VideoId			string`json:"videoId"`
    ChannelId		string`json:"channelId"`
    PlaylistId		string`json:"playlistId"`
  }

type Snippet struct {
	PublishedAt		string
    ChannelId 		string
    Title 			string
    Description 	string
}

type Search struct {
	Id 			ID`json:"id"`
	Snippet 	Snippet`json:"snippet"`
}

func getSong(context *gin.Context) {
	context.IndentedJSON(http.StatusOK(), )
}

func main() {
	router := gin.Default()
	router.GET("/search", getSong)
	router.Run("localhost:8080")
}
