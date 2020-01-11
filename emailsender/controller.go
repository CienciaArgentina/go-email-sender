package emailsender

import (
	"net/http"

	"github.com/CienciaArgentina/email-sender/commons"
	"github.com/CienciaArgentina/email-sender/defines"
	"github.com/gin-gonic/gin"
)

type IEmailController interface {
	SendEmail(c *gin.Context)
}

type EmailController struct {
	Service *EmailSenderService
}

func (emctl *EmailController) SendEmail(c *gin.Context) {
	var dto commons.DTO

	err := c.BindJSON(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, commons.NewBaseResponse(http.StatusBadRequest, nil, err, defines.StringEmpty))
	}
}

func NewController() *EmailController {
	controller := EmailController{Service: NewService()}
	return &controller
}


