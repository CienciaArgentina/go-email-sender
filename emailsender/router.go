package emailsender

import (
	"github.com/CienciaArgentina/go-backend-commons/pkg/rest"
	"github.com/CienciaArgentina/go-email-sender/defines"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	Router *gin.Engine
)

func InitRouter(controller IEmailController) {
	Router = gin.Default()
	ConfigureRoutes(Router, controller)

	gin.ForceConsoleColor()
	Router.RedirectTrailingSlash = true
	Router.RedirectFixedPath = true
	Router.Use(rest.SetContextInformation)
	port := os.Getenv(defines.Port)
	if port == "" {
		port = defines.DefaultPort
	}
	Router.Run(port)
}

func ConfigureRoutes(router *gin.Engine, controller IEmailController) {
	router.GET(defines.Ping, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST(defines.PostEmail, controller.SendEmail)
}
