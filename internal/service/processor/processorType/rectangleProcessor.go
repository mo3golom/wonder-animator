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

type RectangleProcessor struct {
	options *processorOptions.RectangleOptions
	*ProcessorStruct
}

func (rp *RectangleProcessor) Processing(_ *dto.Block, _ *dto.FrameData) (output *image.RGBA, err error) {
	if nil == rp.options {
		return nil, errors.New("предварительно необходимо преобразовать настройки")
	}

	output = image.NewRGBA(
		image.Rect(
			0,
			0,
			int(math.Ceil(rp.options.Width)),
			int(math.Ceil(rp.options.Height)),
		),
	)
	graphicContext := draw2dimg.NewGraphicContext(output)

	rp.DrawBuilder.
		SetGraphicContext(graphicContext).
		SetFillColor(draw2dExtend.ParseHexColor(rp.options.FillColor)).
		SetStrokeColor(draw2dExtend.ParseHexColor(rp.options.StrokeColor)).
		SetStrokeWidth(rp.options.StrokeWidth).
		DrawRoundedRectangle(rp.options.Width, rp.options.Height, rp.options.RoundedRadius)

	return output, nil
}

func (rp *RectangleProcessor) TransformOptions(options *map[string]interface{}) ProcessorInterface {
	rectOptions := processorOptions.NewRectangleOptions()
	_ = mapstructure.Decode(*options, rectOptions)

	rp.options = rectOptions

	return rp
}
