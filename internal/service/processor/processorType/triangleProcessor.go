package processorType

import (
	"errors"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/processorOptions"
	"github.com/mo3golom/wonder-animator/internal/transformer"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	"image/draw"
)

type TriangleProcessor struct {
	options *processorOptions.RectangleOptions
	*ProcessorStruct
}

func (tp *TriangleProcessor) Processing(dest draw.Image, block *dto.Block, frameData *dto.FrameData) (err error) {
	if nil == tp.options {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	graphicContext := draw2dExtend.NewGraphicContext(dest)
	effectValues := tp.applyEffects(block, frameData)

	tp.drawBuilder.
		SetGraphicContext(graphicContext).
		SetX(effectValues.X()).
		SetY(effectValues.Y()).
		SetRotate(effectValues.Rotate()).
		SetOpacity(effectValues.Opacity()).
		SetScale(effectValues.Scale()).
		SetRotatePointType(effectValues.RotatePoint).
		SetFillColor(tp.options.FillColor()).
		SetStrokeColor(tp.options.StrokeColor()).
		SetStrokeWidth(tp.options.StrokeWidth()).
		DrawTriangle(tp.options.Width(), tp.options.Height())

	return nil
}

func (tp *TriangleProcessor) TransformOptions(options *map[string]string) ProcessorInterface {
	tp.options = transformer.TransformRectangleOptions(*options)

	return tp
}
