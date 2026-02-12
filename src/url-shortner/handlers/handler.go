package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const shortUrlPrefix = "http://wu.x/"

type urlWrapper struct {
	Url string `json:"url"`
}

type UrlHandler struct {
	store map[string]string
}

func (u *UrlHandler) Shorten (c *gin.Context) {
	var short_url urlWrapper

	var long_url urlWrapper
	c.BindJSON(&long_url)

	var value string
	var is_found bool

	if value, is_found = u.store[long_url.Url]; is_found {
		short_url.Url = value
		c.IndentedJSON(http.StatusOK, short_url)

		return
	}

	short_url.Url = fmt.Sprintf("%s%v", shortUrlPrefix, generateRandomHexChars(3))
	u.store[long_url.Url] = short_url.Url

	// c.JSON(http.StatusOK, long_url)
	c.IndentedJSON(http.StatusOK, short_url)
}

func (u *UrlHandler) Unshorten (c *gin.Context) {
	var short_url urlWrapper
	c.BindJSON(&short_url)

	for k, v := range u.store {
		if v == short_url.Url {
			var long_url urlWrapper
			long_url.Url = k
			c.IndentedJSON(http.StatusOK, long_url)

			return
		}
	}

	c.String(http.StatusNotFound, fmt.Sprintf("The short url %s is not found\n", short_url.Url))
}

func NewUrlHandler (store map[string]string) *UrlHandler {
	return &UrlHandler{store: store}
}

func generateRandomHexChars (number_of_char int) string {
	rst := rand.New(rand.NewSource(time.Now().Unix()))

	b := make([]byte, number_of_char)
	rst.Read(b)

	return fmt.Sprintf("%x", b[:number_of_char])
}