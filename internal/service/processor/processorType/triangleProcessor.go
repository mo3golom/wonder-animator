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

type TriangleProcessor struct {
	options *processorOptions.RectangleOptions
	*ProcessorStruct
}

func (tp *TriangleProcessor) Processing(_ *dto.Block, _ *dto.FrameData) (output *image.RGBA, err error) {
	if nil == tp.options {
		return nil, errors.New("предварительно необходимо преобразовать настройки")
	}

	output = image.NewRGBA(
		image.Rect(
			0,
			0,
			int(math.Ceil(tp.options.Width)),
			int(math.Ceil(tp.options.Height)),
		),
	)
	graphicContext := draw2dimg.NewGraphicContext(output)

	tp.DrawBuilder.
		SetGraphicContext(graphicContext).
		SetFillColor(draw2dExtend.ParseHexColor(tp.options.FillColor)).
		SetStrokeColor(draw2dExtend.ParseHexColor(tp.options.StrokeColor)).
		SetStrokeWidth(tp.options.StrokeWidth).
		DrawTriangle(tp.options.Width, tp.options.Height)

	return output, nil
}

func (tp *TriangleProcessor) TransformOptions(options *map[string]interface{}) ProcessorInterface {
	triangleOptions := processorOptions.NewRectangleOptions()
	_ = mapstructure.Decode(*options, triangleOptions)

	tp.options = triangleOptions

	return tp
}
