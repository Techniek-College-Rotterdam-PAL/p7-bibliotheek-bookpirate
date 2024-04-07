package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/models"
)

// SearchBooks returns list of books of requested name from database
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

// FeedBooks returns random list of books from database given the amount requested
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
		return
	}
	data, err := json.Marshal(results)
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

func ReserveBook(c *gin.Context) {

}

// AddBook adds book to database
func AddBook(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, Message{
			Code:    MethodNotAllowed,
			Message: messages[MethodNotAllowed],
		})
		return
	}
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Message{
			Code:    MalformedContent,
			Message: messages[MalformedContent],
		})
		return
	}
	var user models.User
	if err := db.Where("token = ?", c.GetHeader(authorizationHeader)).First(&user).Error; err == nil {
		c.JSON(http.StatusUnauthorized, Message{
			Code:    InvalidSession,
			Message: messages[InvalidSession],
		})
		return
	}

	if err := db.Where("isbn = ?", book.Isbn).First(&book).Error; err == nil {
		book.Stock++
		if err = db.Where("isbn = ?", book.Isbn).Save(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, Message{
				Code:    DatabaseQueryError,
				Message: messages[DatabaseQueryError],
			})
			return
		}
	} else {
		if err = db.Create(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, Message{
				Code:    DatabaseQueryError,
				Message: messages[DatabaseQueryError],
			})
			return
		}
	}

	c.JSON(http.StatusOK, Message{
		Code:    SuccessfulInsert,
		Message: messages[SuccessfulInsert],
	})
}

// RemoveBook removes a book from the database if user is admin
func RemoveBook(c *gin.Context) {
	if c.Request.Method != http.MethodDelete {
		c.JSON(http.StatusMethodNotAllowed, Message{
			Code:    MethodNotAllowed,
			Message: messages[MethodNotAllowed],
		})
		return
	}
	var removeRequest models.DeleteBook
	if err := c.ShouldBindJSON(&removeRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Message{
			Code:    MalformedContent,
			Message: messages[MalformedContent],
		})
		return
	}

	var user models.User
	if err := db.Where("token = ?", c.GetHeader(authorizationHeader)).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, Message{
			Code:    InvalidSession,
			Message: messages[InvalidSession],
		})
		return
	}
	var admin models.Admin
	if err := db.Where("ID = ?", user.ID).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, Message{
			Code:    InsufficientPermissions,
			Message: messages[InsufficientPermissions],
		})
		return
	}
	var book models.Book
	if err := db.Where("isbn = ?", removeRequest.Isbn).Delete(&book).Error; err == nil {
		c.JSON(http.StatusInternalServerError, Message{
			Code:    IsbnNotFound,
			Message: messages[IsbnNotFound],
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func countBooks(isbn string) (int64, error) {
	var books int64
	if err := db.Where("isbn = ?").Count(&books).Error; err != nil {
		return 0, err
	}
	return books, nil
}
