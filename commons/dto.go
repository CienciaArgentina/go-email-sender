package commons

type DTO struct {
	To       []string    `json:"to"`
	Data     interface{} `json:"data"`
	Template string      `json:"template"`
}

func NewDTO(to []string, data interface{}, template string) *DTO {
	return &DTO{
		To:       to,
		Data:     data,
		Template: template,
	}
}
