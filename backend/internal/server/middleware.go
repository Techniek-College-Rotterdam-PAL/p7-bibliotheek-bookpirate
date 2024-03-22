package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/models"
)

func SearchBooks(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, Message{
			Code:    MethodNotAllowed,
			Message: messages[MethodNotAllowed],
		})
		return
	}
	var searchRequest models.SearchRequest
	if err := c.ShouldBindJSON(&searchRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Message{
			Code:    InvalidAuthenticationRequest,
			Message: messages[InvalidAuthenticationRequest],
		})
		return
	}

	searchData, ok := db.Get(searchRequest.Name)
	if !ok {
		c.JSON(http.StatusInternalServerError, Message{
			Code:    DatabaseQueryError,
			Message: messages[DatabaseQueryError],
		})
	}
	data, err := json.Marshal(searchData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message{
			Code:    InternalServerError,
			Message: messages[InternalServerError],
		})
		return
	}
	c.JSON(http.StatusOK, Message{
		Data: data,
	})
}

func FeedBooks(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, Message{
			Code:    MethodNotAllowed,
			Message: messages[MethodNotAllowed],
		})
		return
	}
	var feedRequest models.FeedRequest
	if err := c.ShouldBindJSON(&feedRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Message{
			Code:    InvalidAuthenticationRequest,
			Message: messages[InvalidAuthenticationRequest],
		})
		return
	}

	var results []models.Book
	if err := db.Order("RAND()").Limit(feedRequest.Limit).Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Message{
			Code:    DatabaseQueryError,
			Message: messages[DatabaseQueryError],
		})
	}
	data, err := json.Marshal(results)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message{
			Code:    InternalServerError,
			Message: messages[InternalServerError],
		})
		return
	}
	fmt.Print(data)
	c.JSON(http.StatusOK, Message{
		Data: data,
	})
}
