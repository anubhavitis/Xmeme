package ends

import (
	"net/http"
	"strconv"

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

//DelMeme handler
func DelMeme(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if e := db.RemRec(id); e != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	return
}

//EditMeme handler
func EditMeme(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var body struct {
		URL     string `json:"url"`
		Caption string `json:"caption"`
	}
	c.BindJSON(&body)

	if e := db.EditRec(id, body.URL, body.Caption); e != nil {
		c.Writer.WriteHeader(http.StatusConflict)
		return
	}

	c.Writer.WriteHeader(http.StatusAccepted)
	return

}
