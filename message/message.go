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

func PostMessage(c *gin.Context) {
	var message Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMessage := model.Message{Content: message.Content}

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newMessage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newMessage)
}
