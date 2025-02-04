package server

import (
	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"server/internal/util"
	"strings"

	"github.com/gin-gonic/gin"
)

// TODO: add more db instances
var db = &gorm.DB{}
var (
	authDb = func() *gorm.DB { return &gorm.DB{} }
	mainDb = func() *gorm.DB { return &gorm.DB{} }
)

func Run(router *gin.Engine) {
	config := util.LoadConfigFile()

	router.NoRoute(func(c *gin.Context) {
		if strings.Contains(c.Request.RequestURI, "api") || c.Request.Method != http.MethodGet {
			c.JSON(http.StatusBadRequest, Message{})
		} else {
			c.File("../../../static/error-page.html")
		}
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, Message{
			Code:    MethodNotAllowed,
			Message: messages[MethodNotAllowed],
		})
	})
	router.Static("../../../static", "../../.././static")
	StaticRoutes(router)

	for s, middleware := range middleWares {
		info := strings.Split(s, ":")
		path := info[0]

		switch info[1] {
		case http.MethodGet:
			router.GET(path, middleware)
		case http.MethodPost:
			router.POST(path, middleware)
		case http.MethodDelete:
			router.DELETE(path, middleware)
		case http.MethodPut:
			router.PUT(path, middleware)
		case http.MethodPatch:
			router.PATCH(path, middleware)
		}
	}

	var err error
	if db, err = gorm.Open(mysql.Open(config.Database.Dsn), &gorm.Config{}); err != nil {
		log.Fatal(err)
	}
	//if err = router.RunTLS(config.Server.IP+":"+cast.ToString(config.Server.Port), "../../../certificate.pem", "../../../key.pem"); err != nil {
	//	log.Fatal(err)
	//}
	HostBooks(router)
	if err = router.Run(config.Server.IP + ":" + cast.ToString(config.Server.Port)); err != nil {
		log.Println(err)
	}
}

func StaticRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.File("../../../static/index.html")
	})
	router.GET("/books", func(c *gin.Context) {
		c.File("../../../static/books.html")
	})
	router.GET("/login", func(c *gin.Context) {
		c.File("../../../static/login.html")
	})
	router.GET("/register", func(c *gin.Context) {
		c.File("../../../static/register.html")
	})
	router.GET("/contact-owner", func(c *gin.Context) {
		c.File("../../../static/contact-owner.html")
	})
	router.GET("/dashboard", func(c *gin.Context) {
		c.File("../../../static/dashboard.html")
	})
	router.GET("/search", func(c *gin.Context) {
		c.File("../../../static/search.html")
	})
	router.GET("/reserved", func(c *gin.Context) {
		c.File("../../../static/reserved.html")
	})
	router.GET("/api-search", func(c *gin.Context) {
		c.File("../../../static/api-search.html")
	})
}
