package WonderAnimator

import "github.com/mo3golom/wonder-animator/internal/dto"

const (
	defaultFPS = 24
)

type InputObject struct {
	Width           int          `json:"width" mapstructure:"width"`
	Height          int          `json:"height" mapstructure:"height"`
	Duration        float32      `json:"duration" mapstructure:"duration"`
	Blocks          *[]dto.Block `json:"blocks" mapstructure:"blocks"`
	FPS             *int         `json:"fps" mapstructure:"fps"`
	BackgroundImage *string      `json:"backgroundImage" mapstructure:"backgroundImage"`
	BackgroundColor *string      `json:"backgroundColor" mapstructure:"backgroundColor"`
}

func (i *InputObject) GetFPS() int {
	if nil == i.FPS {
		return defaultFPS
	}

	return *i.FPS
}
