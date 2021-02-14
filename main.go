package main

import (
	"fmt"

	"github.com/anubhavitis/Xmeme/backend/db"
	ends "github.com/anubhavitis/Xmeme/backend/endpoints"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// _ "github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/docs"
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
	rou.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	rou.POST("/memes", ends.NewMemes)
	rou.GET("/memes", ends.AllMemes)
	rou.DELETE("/memes/:id", ends.DelMeme)
	rou.PATCH("/memes/:id", ends.EditMeme)
	rou.Run()
}
