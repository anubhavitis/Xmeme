package main

import (
	"fmt"
	"os"

	"github.com/anubhavitis/Xmeme/backend/db"
	ends "github.com/anubhavitis/Xmeme/backend/endpoints"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	if dab, err := db.InitDb(); err != nil {
		fmt.Println(err)
		return
	} else {
		db.Mydb = dab
	}

	rou := gin.Default()
	rou.Use(cors.Default())
	rou.POST("/memes", ends.NewMemes)
	rou.GET("/memes", ends.AllMemes)
	rou.DELETE("/memes/:id", ends.DelMeme)
	rou.PATCH("/memes/:id", ends.EditMeme)

	// rou.Run()
	port := os.Getenv("PORT")
	if port == "" {
		rou.Run(":8081")
	} else {
		rou.Run()
	}
}
