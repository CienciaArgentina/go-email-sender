package commons

import (
	"fmt"
	"github.com/CienciaArgentina/go-email-sender/defines"
)

const (
	BaseTemplatePath = "../templates/"
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
		Filename: fmt.Sprintf("%s%s.html", BaseTemplatePath ,defines.ConfirmEmail),
		Subject:  "Confirma tu email",
		Entity:   ConfirmationMailBody{},
	}
}
