package draw2dExtend

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image/draw"
	"math"
)

type GraphicContext struct {
	*draw2dimg.GraphicContext
}

func NewGraphicContext(dest draw.Image) *GraphicContext {
	return &GraphicContext{
		draw2dimg.NewGraphicContext(dest),
	}
}

func (gc *GraphicContext) RotateAndScale(rotateDegrees float64, scale float64, point Point) *GraphicContext {
	// Преобразование градусов в радианы
	rotateRadians := -(rotateDegrees * (math.Pi / 180.0))

	// Translate достаточно дорогая операция, поэтому действия, которые зависят от нее, лучше объединить в один метод
	gc.GraphicContext.Translate(point.X, point.Y)
	gc.GraphicContext.Scale(scale, scale)
	gc.GraphicContext.Rotate(rotateRadians)
	gc.GraphicContext.Translate(-point.X, -point.Y)

	return gc
}

func (gc *GraphicContext) Scale(scale float64, point Point) *GraphicContext {
	// Translate достаточно дорогая операция, поэтому действия, которые зависят от нее, лучше объединить в один метод
	gc.GraphicContext.Translate(point.X, point.Y)
	gc.GraphicContext.Scale(scale, scale)
	gc.GraphicContext.Translate(-point.X, -point.Y)

	return gc
}
