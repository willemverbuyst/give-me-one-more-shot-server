package message

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willemverbuyst/give-me-one-more-shot/server/model"
)

type Message struct {
	Content string `json:"content" binding:"required"`
}

func GetMessages(c *gin.Context) {

	var messages []model.Message

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&messages).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, messages)

}
