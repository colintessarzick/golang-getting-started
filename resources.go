package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}

func getAlbums(context *gin.Context) {
	context.JSON(http.StatusOK, albums)
}

func getAlbum(context *gin.Context) {
	id := context.Param("id")

	for index := 0; index < len(albums); index++ {
		if albums[index].ID == id {
			context.JSON(http.StatusOK, albums[index])
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func postAlbum(context *gin.Context) {
	var newAlbum album

	if err := context.BindJSON(&newAlbum); err != nil {
		fmt.Println("Unexpected Error:", err)
		return
	}

	albums = append(albums, newAlbum)
	checkNewAlbumWasAdded()
	context.JSON(http.StatusCreated, newAlbum)
}

func checkNewAlbumWasAdded() {
	if len(albums) == 4 {
		fmt.Println("New album was added!")
	} else {
		fmt.Printf("Albums slice still contains %v elements", len(albums))
	}

	for index := 0; index < len(albums); index++ {
		fmt.Println(albums[index])
	}
}