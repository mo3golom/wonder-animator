package dto

type Processor struct {
	Id      string            `json:"id"`
	Options map[string]string `json:"options"`
}
