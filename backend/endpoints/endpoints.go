package ends

import (
	"net/http"

	"github.com/anubhavitis/Xmeme/backend/db"
	"github.com/anubhavitis/Xmeme/backend/util"
	"github.com/gin-gonic/gin"
)

//NewMemes handler
func NewMemes(c *gin.Context) {
	var body util.Meme
	c.BindJSON(&body)
	if id, e := db.AddMeme(body); e != nil {
		c.JSON(http.StatusConflict, gin.H{"Error": e})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

//AllMemes handler
func AllMemes(c *gin.Context) {
	memes, e := db.ShowAllMeme()
	if e != nil {
		c.JSON(http.StatusConflict, gin.H{"error": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": memes})
}

//DellAll handler
func DellAll(c *gin.Context) {
	if e := db.ResetTable(); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": true})
	return
}
