package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"server/internal/models"
)

var db *gorm.DB

func Login(c *gin.Context) {

}

func Home(c *gin.Context) {

}

func Register(c *gin.Context) {
	c.Header(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, Message{
			Code:    MethodNotAllowed,
			Message: messages[MethodNotAllowed],
		})
		return
	}
	var regRequest models.RegistrationRequest
	if err := c.ShouldBindJSON(&regRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Message{
			Code:    InvalidAuthenticationRequest,
			Message: messages[InvalidAuthenticationRequest],
		})
		return
	}
	var user models.User
	if err := db.Where("email = ?", regRequest.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusNotAcceptable, Message{
			Code:    UsernameAlreadyTaken,
			Message: messages[UsernameAlreadyTaken],
		})
		return
	}
	//hmac := hmac.New(sha256.New, secret)
	//hmac.Write([]byte(data))
	//dataHmac := hmac.Sum(nil)
	//
	//cipherText := hex.EncodeToString(dataHmac)
	//key := hex.EncodeToString(secret)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Message{
			Code:    DatabaseQueryError,
			Message: messages[DatabaseQueryError],
		})
		return
	}

	c.JSON(http.StatusOK, Message{
		Code:    SuccessfulRegistration,
		Message: messages[SuccessfulRegistration],
	})

}
