package processorOptions

import "github.com/mo3golom/wonder-animator/pkg/draw2dExtend"

type RectangleOptions struct {
	width, height float64
	roundedRadius float64
	strokeWidth   float64
	fillColor     draw2dExtend.ExtendColor
	strokeColor   draw2dExtend.ExtendColor
}

func NewRectangleOptions() *RectangleOptions {
	return &RectangleOptions{
		width:         0,
		height:        0,
		roundedRadius: 0,
		strokeWidth:   0,
		fillColor:     draw2dExtend.ParseHexColor("#ffffff"),
		strokeColor:   draw2dExtend.ParseHexColor("#000000"),
	}
}

func (r *RectangleOptions) Width() float64 {
	return r.width
}

func (r *RectangleOptions) SetWidth(width float64) {
	r.width = width
}

func (r *RectangleOptions) Height() float64 {
	return r.height
}

func (r *RectangleOptions) SetHeight(height float64) {
	r.height = height
}

func (r *RectangleOptions) RoundedRadius() float64 {
	return r.roundedRadius
}

func (r *RectangleOptions) SetRoundedRadius(roundedRadius float64) {
	r.roundedRadius = roundedRadius
}

func (r *RectangleOptions) StrokeWidth() float64 {
	return r.strokeWidth
}

func (r *RectangleOptions) SetStrokeWidth(strokeWidth float64) {
	r.strokeWidth = strokeWidth
}

func (r *RectangleOptions) FillColor() draw2dExtend.ExtendColor {
	return r.fillColor
}

func (r *RectangleOptions) SetFillColor(fillColor string) {
	r.fillColor = draw2dExtend.ParseHexColor(fillColor)
}

func (r *RectangleOptions) StrokeColor() draw2dExtend.ExtendColor {
	return r.strokeColor
}

func (r *RectangleOptions) SetStrokeColor(strokeColor string) {
	r.strokeColor = draw2dExtend.ParseHexColor(strokeColor)
}
