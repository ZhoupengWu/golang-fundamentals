package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"urlShortnerPg/dtos"
	"urlShortnerPg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UrlHandler struct {
	store *gorm.DB
}

func (u *UrlHandler) Shorten (c *gin.Context) {
	var short_url dtos.UrlWrapper

	var long_url dtos.UrlWrapper
	c.BindJSON(&long_url)

	var converted_url models.ConvertedUrl
	err := u.store.First(&converted_url, "long_url = ?", long_url.Url).Error

	// Se ha trovato un record ritorna nil
	if err == nil {
		short_url.Url = converted_url.Short_url
		c.IndentedJSON(http.StatusOK, short_url)

		return
	}

	short_url.Url = fmt.Sprintf("%s%v",  models.ShortUrlPrefix, generateRandomHexChars(3))
	u.store.Create(models.NewConvertedUrl(long_url.Url, short_url.Url))

	c.IndentedJSON(http.StatusOK, short_url)
}

func (u *UrlHandler) Unshorten (c *gin.Context) {
	var short_url dtos.UrlWrapper
	c.BindJSON(&short_url)

	var converted_url models.ConvertedUrl
	err := u.store.Where("short_url = ?", short_url.Url).First(&converted_url).Error

	if err != nil {
		c.String(http.StatusNotFound, fmt.Sprintf("The short url %s is not found\n", short_url.Url))

		return
	}

	var long_url dtos.UrlWrapper
	long_url.Url = converted_url.Long_url
	c.IndentedJSON(http.StatusOK, long_url)
}

func NewUrlHandler (db *gorm.DB) *UrlHandler {
	return &UrlHandler{store: db}
}

func generateRandomHexChars (number_of_char int) string {
	rst := rand.New(rand.NewSource(time.Now().Unix()))

	b := make([]byte, number_of_char)
	rst.Read(b)

	return fmt.Sprintf("%x", b[:number_of_char])
}