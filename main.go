package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//NewMemes handler
func NewMemes(c *gin.Context) {
	var body utility.Meme
	c.BindJSON(&body)
	fmt.Println(body)
	c.JSON(http.StatusOK, gin.H{"Name": body.Name, "URL": body.URL})
	// c.String(http.StatusOK, "Hello for GET request, %s", name)
}

func main() {
	r := gin.Default()
	r.POST("/memes", NewMemes)
	// r.GET("/memes", AllMemes)
	r.Run()
}
