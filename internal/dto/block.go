package dto

import (
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
	"github.com/mo3golom/wonder-glitch/wonderGlitchDTO"
)

type Block struct {
	Type         *Processor                    `json:"processor" mapstructure:"processor"`
	StartAt      float32                       `json:"startAt" mapstructure:"startAt"`
	Duration     float32                       `json:"duration" mapstructure:"duration"`
	Position     *Position                     `json:"position" mapstructure:"position"`
	Effects      []wonderEffectDTO.Effect      `json:"effects" mapstructure:"effects"`
	Glitches     []wonderGlitchDTO.InputEffect `json:"glitches" mapstructure:"glitches"`
	GlitchFactor float64                       `json:"glitchFactor" mapstructure:"glitchFactor"`
	Opacity      float32                       `json:"opacity" mapstructure:"opacity"`
	Rotate       float64                       `json:"rotate" mapstructure:"rotate"`
	RotatePoint  string                        `json:"rotatePoint" mapstructure:"rotatePoint"`
	Scale        float64                       `json:"scale" mapstructure:"scale"`
}
