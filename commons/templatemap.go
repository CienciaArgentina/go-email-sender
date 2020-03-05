package commons

import (
	"fmt"
	"github.com/CienciaArgentina/go-email-sender/defines"
)

type TemplateInfo struct {
	Type     string
	Subject  string
	Filename string
	Entity   interface{}
}

var TemplateMap map[string]TemplateInfo

func init() {
	TemplateMap = make(map[string]TemplateInfo)

	TemplateMap[defines.ConfirmEmail] = TemplateInfo{
		Type:     defines.ConfirmEmail,
		Filename: fmt.Sprintf("%s.html", defines.ConfirmEmail),
		Subject:  "Ciencia Argentina - Confirma tu email",
		Entity:   ConfirmationMailBody{},
	}

	TemplateMap[defines.ForgotUsername] = TemplateInfo{
		Type:     defines.ForgotUsername,
		Filename: fmt.Sprintf("%s.html", defines.ForgotUsername),
		Subject:  "Ciencia Argentina - Recuperar nombre de usuario",
		Entity:   ForgotUsernameBody{},
	}

	TemplateMap[defines.SendPasswordReset] = TemplateInfo{
		Type:     defines.SendPasswordReset,
		Filename: fmt.Sprintf("%s.html", defines.SendPasswordReset),
		Subject:  "Ciencia Argentina - Reestablecer contraseña",
		Entity:   SendPasswordResetBody{},
	}

	TemplateMap[defines.PasswordResetNotification] = TemplateInfo{
		Type:     defines.PasswordResetNotification,
		Filename: fmt.Sprintf("%s.html", defines.PasswordResetNotification),
		Subject:  "Ciencia Argentina - Contraseña reestablecida",
		Entity:   nil,
	}
}
