package processorType

import (
	"errors"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/processorOptions"
	"github.com/mo3golom/wonder-animator/internal/transformer"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	"image/draw"
)

type CircleProcessor struct {
	options *processorOptions.CircleOptions
	*ProcessorStruct
}

func (cp *CircleProcessor) Processing(dest draw.Image, block *dto.Block, frameData *dto.FrameData) (err error) {
	if nil == cp.options {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	graphicContext := draw2dExtend.NewGraphicContext(dest)
	effectValues := cp.applyEffects(block, frameData)

	cp.drawBuilder.
		SetGraphicContext(graphicContext).
		SetX(effectValues.X()).
		SetY(effectValues.Y()).
		SetRotate(effectValues.Rotate()).
		SetOpacity(effectValues.Opacity()).
		SetScale(effectValues.Scale()).
		SetRotatePointType(effectValues.RotatePoint).
		SetFillColor(cp.options.FillColor()).
		SetStrokeColor(cp.options.StrokeColor()).
		SetStrokeWidth(cp.options.StrokeWidth()).
		DrawCircle(cp.options.Radius())

	return nil
}

func (cp *CircleProcessor) TransformOptions(options *map[string]string) ProcessorInterface {
	cp.options = transformer.TransformCircleOptions(*options)

	return cp
}
