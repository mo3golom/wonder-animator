package processorType

import (
	"errors"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/mitchellh/mapstructure"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/processorOptions"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	"image"
	"math"
)

type CircleProcessor struct {
	options *processorOptions.CircleOptions
	*ProcessorStruct
}

func (cp *CircleProcessor) Processing(_ *dto.Block, _ *dto.FrameData) (output *image.RGBA, err error) {
	if nil == cp.options {
		return nil, errors.New("предварительно необходимо преобразовать настройки")
	}

	diameter := int(math.Ceil(cp.options.Radius)) * 2
	output = image.NewRGBA(
		image.Rect(
			0,
			0,
			diameter,
			diameter,
		),
	)
	graphicContext := draw2dimg.NewGraphicContext(output)

	cp.DrawBuilder.
		SetGraphicContext(graphicContext).
		SetFillColor(draw2dExtend.ParseHexColor(cp.options.FillColor)).
		SetStrokeColor(draw2dExtend.ParseHexColor(cp.options.StrokeColor)).
		SetStrokeWidth(cp.options.StrokeWidth).
		DrawCircle(cp.options.Radius)

	return output, nil
}

func (cp *CircleProcessor) TransformOptions(options *map[string]interface{}) ProcessorInterface {
	circleOptions := processorOptions.NewCircleOptions()
	_ = mapstructure.Decode(*options, circleOptions)

	cp.options = circleOptions

	return cp
}
