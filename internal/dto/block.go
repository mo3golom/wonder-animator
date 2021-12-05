package dto

import (
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
)

type Block struct {
	Type        *Processor                `json:"type"`
	StartAt     float32                   `json:"startAt"`
	Duration    float32                   `json:"duration"`
	Position    *Position                 `json:"position"`
	Effects     *[]wonderEffectDTO.Effect `json:"effects"`
	Opacity     float32                   `json:"opacity"`
	Rotate      float64                   `json:"rotate"`
	RotatePoint string                    `json:"rotatePoint"`
	Scale       float64                   `json:"scale"`
}
