package main

import (
	"urlShortnerPg/handlers"
	"urlShortnerPg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main () {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=54322 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.ConvertedUrl{})

	handler := handlers.NewUrlHandler(db)

	router := gin.Default()

	router.POST("/shorten", handler.Shorten)
	router.POST("/unshorten", handler.Unshorten)

	router.Run(":8000")
}