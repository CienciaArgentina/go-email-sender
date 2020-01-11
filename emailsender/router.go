package emailsender

import (
	"github.com/CienciaArgentina/email-sender/defines"
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
	port := os.Getenv("EMAILSENDER_PORT")
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
