package main

import (
	"urlShortner/handlers"

	"github.com/gin-gonic/gin"
)

func main () {
	url_stored := make(map[string]string, 0)
	handler := handlers.NewUrlHandler(url_stored)

	router := gin.Default()

	router.POST("/shorten", handler.Shorten)
	router.POST("/unshorten", handler.Unshorten)

	router.Run(":8000")
}