package processorType

import (
	"image"

	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
)

type ProcessorInterface interface {
	Processing(block *dto.Block, frameData *dto.FrameData) (output *image.RGBA, err error)
	TransformOptions(options *map[string]interface{}) ProcessorInterface
}

type ProcessorStruct struct {
	DrawBuilder *draw2dExtend.DrawBuilder
}
