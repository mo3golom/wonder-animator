package dto

type Position struct {
	X float64 `json:"x" mapstructure:"x"`
	Y float64 `json:"y" mapstructure:"y"`
}
