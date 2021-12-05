package processorOptions

import "github.com/mo3golom/wonder-animator/pkg/draw2dExtend"

type CircleOptions struct {
	radius      float64
	strokeWidth float64
	fillColor   draw2dExtend.ExtendColor
	strokeColor draw2dExtend.ExtendColor
}

func NewCircleOptions() *CircleOptions {
	return &CircleOptions{
		radius:      0,
		strokeWidth: 0,
		fillColor:   draw2dExtend.ParseHexColor("#ffffff"),
		strokeColor: draw2dExtend.ParseHexColor("#000000"),
	}
}

func (c *CircleOptions) Radius() float64 {
	return c.radius
}

func (c *CircleOptions) Diameter() float64 {
	return c.radius * 2
}

func (c *CircleOptions) SetRadius(radius float64) {
	c.radius = radius
}

func (c *CircleOptions) StrokeWidth() float64 {
	return c.strokeWidth
}

func (c *CircleOptions) SetStrokeWidth(strokeWidth float64) {
	c.strokeWidth = strokeWidth
}

func (c *CircleOptions) FillColor() draw2dExtend.ExtendColor {
	return c.fillColor
}

func (c *CircleOptions) SetFillColor(fillColor string) {
	c.fillColor = draw2dExtend.ParseHexColor(fillColor)
}

func (c *CircleOptions) StrokeColor() draw2dExtend.ExtendColor {
	return c.strokeColor
}

func (c *CircleOptions) SetStrokeColor(strokeColor string) {
	c.strokeColor = draw2dExtend.ParseHexColor(strokeColor)
}
