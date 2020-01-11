package commons

type DTO struct {
	To       []string
	Data     interface{}
	Template string
}

func NewDTO(to []string, data interface{}, template string) *DTO {
	return &DTO{
		To: to,
		Data: data,
		Template: template,
	}
}