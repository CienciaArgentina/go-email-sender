package emailsender

import (
	"bytes"
	"encoding/json"
	"github.com/CienciaArgentina/go-backend-commons/pkg/apierror"
	"github.com/CienciaArgentina/go-backend-commons/pkg/performance"
	"github.com/CienciaArgentina/go-backend-commons/pkg/rest"
	"net/http"
	"time"

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
	ctx := rest.GetContextInformation("SendEmail", c)

	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)

	err := json.Unmarshal(buf.Bytes(), &dto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, commons.NewBaseResponse(http.StatusBadRequest, nil, err, defines.StringEmpty))
		return
	}

	var apierr apierror.ApiError
	performance.TrackTime(time.Now(), "InvokeEmailSender", ctx, func() {
		apierr = emctl.Service.InvokeEmailSender(dto, ctx)
	})

	if apierr != nil {
		c.JSON(apierr.Status(), apierr)
	}
	return

}

func NewController() *EmailController {
	controller := EmailController{Service: NewService()}
	return &controller
}