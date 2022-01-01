package dto

type Processor struct {
	Id      string                 `json:"id" mapstructure:"id"`
	Options map[string]interface{} `json:"options" mapstructure:"options"`
}
