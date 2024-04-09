package server

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"server/internal/models"
	"server/internal/util"
	"strconv"
	"strings"
	"time"
)

func Register(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

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

	if err := validator.New().Struct(&regRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Message{
			Code:    InvalidEmail,
			Message: messages[InvalidEmail],
		})
		return
	}

	email := strings.Split(regRequest.Email, "@")
	if email[1] == TCRDocentDomain {
		c.JSON(http.StatusNotAcceptable, Message{
			Code:    AdminNeeded,
			Message: messages[AdminNeeded],
		})
		return
	}
	if len(email) <= 1 || email[1] != TCRStudentDomain || len(email[0]) == 0 {
		c.JSON(http.StatusNotAcceptable, Message{
			Code:    InvalidEmail,
			Message: messages[InvalidEmail],
		})
		return
	}
	if len(regRequest.Password) < 6 || len(regRequest.Password) > defaultAuthLength {
		c.JSON(http.StatusNotAcceptable, Message{
			Code:    InvalidPassword,
			Message: messages[InvalidPassword],
		})
		return
	}

	var user models.User
	if err := db.Where("email = ?", regRequest.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusNotAcceptable, Message{
			Code:    EmailAlreadyUsed,
			Message: messages[EmailAlreadyUsed],
		})
		return
	}
	if err := db.Where("username = ?", regRequest.Username).First(&user).Error; err == nil {
		c.JSON(http.StatusNotAcceptable, Message{
			Code:    UsernameAlreadyTaken,
			Message: messages[UsernameAlreadyTaken],
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message{
			Code:    InternalServerError,
			Message: messages[InternalServerError],
		})
		return
	}
	user = models.User{ // only for now.
		ID:             uuid.New().ID(),
		Email:          regRequest.Email,
		Username:       regRequest.Username,
		HashedPassword: string(hashedPassword),
	}

	currentTime := time.Now()
	sha1Hasher.Write([]byte(strconv.Itoa(int(user.ID))))

	user.Token = fmt.Sprintf("%v.%v.%v",
		base64.RawStdEncoding.EncodeToString([]byte(strconv.Itoa(int(user.ID)))),
		base64.RawStdEncoding.EncodeToString([]byte(strconv.Itoa(int(util.GenerateSnowflake(currentTime))))),
		util.GenerateRandomString(defaultAuthLength)+hex.EncodeToString(sha256Hasher.Sum(nil)),
	)
	data, err := json.Marshal(models.RegistrationResponse{
		Token:     user.Token,
		TimeStamp: currentTime.UnixNano(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message{
			Code:    DatabaseQueryError,
			Message: messages[DatabaseQueryError],
		})
		return
	}
	if err = db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Message{
			Code:    DatabaseQueryError,
			Message: messages[DatabaseQueryError],
		})
		return
	}

	c.JSON(http.StatusCreated, Message{
		Code:    SuccessfulRegistration,
		Message: messages[SuccessfulRegistration],
		Data:    data,
	})
}

func Login(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	c.Header(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, Message{
			Code:    MethodNotAllowed,
			Message: messages[MethodNotAllowed],
		})
		return
	}

	var authRequest models.AuthenticationRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Message{
			Code:    InvalidAuthenticationRequest,
			Message: messages[InvalidAuthenticationRequest],
		})
		return
	}

	if err := validator.New().Struct(&authRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Message{
			Code:    InvalidEmail,
			Message: messages[InvalidEmail],
		})
		return
	}

	var user models.User
	if err := db.First(&user, &models.User{Email: authRequest.Email, Username: authRequest.Username}).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.JSON(http.StatusNotFound, Message{
				Code:    UserNotFound,
				Message: messages[UserNotFound],
			})
		default:
			c.JSON(http.StatusInternalServerError, Message{
				Code:    DatabaseQueryError,
				Message: messages[DatabaseQueryError],
			})
		}
		return
	}
	if user.Token == c.GetHeader(authorizationHeader) && len(user.Token) >= defaultAuthLength {
		c.JSON(http.StatusOK, Message{
			Code:    AlreadyLoggedIn,
			Message: messages[AlreadyLoggedIn],
		})
		return
	} else if c.GetHeader(authorizationHeader) != "" {
		if err := db.Where("token = ?", c.GetHeader(authorizationHeader)).First(&user).Error; err == nil {
			c.JSON(http.StatusUnauthorized, Message{
				Code:    InvalidSession,
				Message: messages[InvalidSession],
			})
		} else {
			c.JSON(http.StatusOK, Message{
				Code:    AlreadyLoggedInDifferentAccount,
				Message: messages[AlreadyLoggedInDifferentAccount],
			})
		}
		return

	}

	if !(user.Email == authRequest.Email && user.Username == authRequest.Username) {
		c.JSON(http.StatusUnauthorized, Message{
			Code:    InvalidAuthenticationRequest,
			Message: messages[InvalidAuthenticationRequest],
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(authRequest.Password)); err != nil {
		c.JSON(http.StatusForbidden, Message{
			Code:    IncorrectPassword,
			Message: messages[IncorrectPassword],
		})
		return
	}

	data, err := json.Marshal(models.AuthenticationResponse{
		Token: user.Token,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message{
			Code:    InternalServerError,
			Message: messages[InternalServerError],
		})
		return
	}
	c.JSON(http.StatusOK, Message{
		Code:    SuccessfulAuthentication,
		Message: messages[SuccessfulAuthentication],
		Data:    data,
	})
}
