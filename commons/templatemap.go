package commons

import (
	"fmt"
	"github.com/CienciaArgentina/email-sender/defines"
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
		Subject:  "Confirma tu email",
		Entity:   ConfirmationMailBody{},
	}
}
