package WonderAnimator

import "github.com/mo3golom/wonder-animator/internal/dto"

const (
	defaultFPS = 24
)

type InputObject struct {
	Width           int          `json:"width"`
	Height          int          `json:"height"`
	Duration        float32      `json:"duration"`
	Blocks          *[]dto.Block `json:"blocks"`
	FPS             *int         `json:"fps"`
	BackgroundImage *string      `json:"backgroundImage"`
	BackgroundColor *string      `json:"backgroundColor"`
}

func (i *InputObject) GetFPS() int {
	if nil == i.FPS {
		return defaultFPS
	}

	return *i.FPS
}
