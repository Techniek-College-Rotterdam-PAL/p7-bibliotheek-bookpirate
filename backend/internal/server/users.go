package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"server/internal/models"
)

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

func Login(c *gin.Context) {
	var authRequest models.AuthenticationRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Message{
			Code:    InvalidAuthenticationRequest,
			Message: messages[InvalidAuthenticationRequest],
		})
		return
	}

	if err := validator.New().Struct(authRequest).(validator.ValidationErrors); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Message{
			Code:    InvalidAuthenticationRequest,
			Message: messages[InvalidAuthenticationRequest],
		})
		return
	}

	var user models.User
	//if err := db.Site().First(&user, &models.User{Email: authRequest.Email}).Error; err != nil {
	//	switch {
	//	case errors.Is(err, gorm.ErrRecordNotFound):
	//		c.JSON(http.StatusNotFound, Message{
	//			Code:    UserNotFound,
	//			Message: messages[UserNotFound],
	//		})
	//	default:
	//		c.JSON(http.StatusInternalServerError, Message{
	//			Code:    DatabaseQueryError,
	//			Message: messages[DatabaseQueryError],
	//		})
	//	}
	//	return
	//}

	if err := bcrypt.CompareHashAndPassword([]byte(authRequest.Password), []byte(user.HashedPassword)); err != nil {
		c.JSON(http.StatusForbidden, Message{
			Code:    IncorrectPassword,
			Message: messages[IncorrectPassword],
		})
		return
	}
}
