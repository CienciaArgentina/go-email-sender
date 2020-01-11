package emailsender

import (
	"github.com/CienciaArgentina/email-sender/defines"
	"github.com/gin-gonic/gin"
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
	Router.Run(defines.Port)
}

func ConfigureRoutes(router *gin.Engine, controller IEmailController) {
	router.GET(defines.Ping, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST(defines.PostEmail, controller.SendEmail)
}
