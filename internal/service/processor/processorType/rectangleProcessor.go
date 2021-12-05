package processorType

import (
	"errors"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/processorOptions"
	"github.com/mo3golom/wonder-animator/internal/transformer"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	"image/draw"
)

type RectangleProcessor struct {
	options *processorOptions.RectangleOptions
	*ProcessorStruct
}

func (rp *RectangleProcessor) Processing(dest draw.Image, block *dto.Block, frameData *dto.FrameData) (err error) {
	if nil == rp.options {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	graphicContext := draw2dExtend.NewGraphicContext(dest)
	effectValues := rp.applyEffects(block, frameData)

	rp.drawBuilder.
		SetGraphicContext(graphicContext).
		SetX(effectValues.X()).
		SetY(effectValues.Y()).
		SetRotate(effectValues.Rotate()).
		SetOpacity(effectValues.Opacity()).
		SetScale(effectValues.Scale()).
		SetRotatePointType(effectValues.RotatePoint).
		SetFillColor(rp.options.FillColor()).
		SetStrokeColor(rp.options.StrokeColor()).
		SetStrokeWidth(rp.options.StrokeWidth()).
		DrawRoundedRectangle(rp.options.Width(), rp.options.Height(), rp.options.RoundedRadius())

	return nil
}

func (rp *RectangleProcessor) TransformOptions(options *map[string]string) ProcessorInterface {
	rp.options = transformer.TransformRectangleOptions(*options)

	return rp
}
