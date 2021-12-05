package processorOptions

import (
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
)

type TextOptions struct {
	text            string
	fontSize        float64
	padding         float64
	radius          float64
	textColor       draw2dExtend.ExtendColor
	backgroundColor draw2dExtend.ExtendColor
}

func NewTextOptions() *TextOptions {
	return &TextOptions{
		text:            "",
		fontSize:        12,
		padding:         0,
		radius:          0,
		textColor:       draw2dExtend.ParseHexColor("#ffffff"),
		backgroundColor: draw2dExtend.ParseHexColor("#000000"),
	}
}

func (t *TextOptions) Text() string {
	return t.text
}

func (t *TextOptions) SetText(text string) {
	t.text = text
}

func (t *TextOptions) FontSize() float64 {
	return t.fontSize
}

func (t *TextOptions) SetFontSize(fontSize float64) {
	t.fontSize = fontSize
}

func (t *TextOptions) Padding() float64 {
	return t.padding
}

func (t *TextOptions) SetPadding(padding float64) {
	t.padding = padding
}

func (t *TextOptions) Radius() float64 {
	return t.radius
}

func (t *TextOptions) SetRadius(radius float64) {
	t.radius = radius
}

func (t *TextOptions) TextColor() draw2dExtend.ExtendColor {
	return t.textColor
}

func (t *TextOptions) SetTextColor(textColor string) {
	t.textColor = draw2dExtend.ParseHexColor(textColor)
}

func (t *TextOptions) BackgroundColor() draw2dExtend.ExtendColor {
	return t.backgroundColor
}

func (t *TextOptions) SetBackgroundColor(backgroundColor string) {
	t.backgroundColor = draw2dExtend.ParseHexColor(backgroundColor)
}
