package main

import (
	"fmt"

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
	rou.GET("/delall", ends.DellAll)
	rou.Run()
}
