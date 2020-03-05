package emailsender

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/CienciaArgentina/go-email-sender/commons"
	"github.com/CienciaArgentina/go-email-sender/defines"
	"github.com/gin-gonic/gin"
)

type IEmailController interface {
	SendEmail(c *gin.Context)
}

type EmailController struct {
	Service *EmailSenderService
}

func (emctl *EmailController) SendEmail(c *gin.Context) {
	dto := commons.DTO{}

	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)

	err := json.Unmarshal(buf.Bytes(), &dto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, commons.NewBaseResponse(http.StatusBadRequest, nil, err, defines.StringEmpty))
		return
	}

	result := emctl.Service.InvokeEmailSender(dto)
	c.JSON(result.Code, result)
	return

}

func NewController() *EmailController {
	controller := EmailController{Service: NewService()}
	return &controller
}