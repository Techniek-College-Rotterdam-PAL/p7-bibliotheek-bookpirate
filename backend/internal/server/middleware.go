package server

import (
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
	//var data models.Book
	//if data, ok = db.Get(searchRequest.Name); !ok {
	//	c.JSON(http.StatusInternalServerError, Message{
	//		Code:    DatabaseQueryError,
	//		Message: messages[DatabaseQueryError],
	//	})
	//}
}
