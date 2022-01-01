package draw2dExtend

import (
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"image/color"
)

type DrawBuilder struct {
	fontData               draw2d.FontData
	fillColor, strokeColor color.RGBA
	graphicContext         *draw2dimg.GraphicContext
	strokeWidth            float64
}

func NewDrawBuilder(fontData draw2d.FontData) *DrawBuilder {
	return &DrawBuilder{
		fillColor:   ParseHexColor("#ffffff"),
		strokeColor: ParseHexColor("#000000"),
		fontData:    fontData,
	}
}

func (d *DrawBuilder) SetFillColor(fillColor color.RGBA) *DrawBuilder {
	d.fillColor = fillColor

	return d
}

func (d *DrawBuilder) SetStrokeColor(strokeColor color.RGBA) *DrawBuilder {
	d.strokeColor = strokeColor

	return d
}

func (d *DrawBuilder) SetStrokeWidth(strokeWidth float64) *DrawBuilder {
	d.strokeWidth = strokeWidth

	return d
}

func (d *DrawBuilder) SetGraphicContext(graphicContext *draw2dimg.GraphicContext) *DrawBuilder {
	d.graphicContext = graphicContext

	return d
}

func (d *DrawBuilder) DrawText(text string, fontSize float64, backgroundOptions *BackgroundOptions) (width, height float64) {
	var backgroundPadding float64 = 0

	if nil != backgroundOptions {
		backgroundPadding = backgroundOptions.Padding
	}

	d.graphicContext.SetFontData(d.fontData)
	d.graphicContext.SetFontSize(fontSize)

	left, top, right, _ := d.graphicContext.GetStringBounds(text)

	if nil != backgroundOptions {
		d.graphicContext.SetFillColor(backgroundOptions.FillColor)
		d.graphicContext.SetLineWidth(0)
		draw2dkit.RoundedRectangle(
			d.graphicContext,
			0,
			0,
			right+(backgroundPadding*2),
			(-top)+(backgroundPadding*2),
			backgroundOptions.Radius,
			backgroundOptions.Radius,
		)
		d.graphicContext.FillStroke()
	}

	d.graphicContext.SetFillColor(d.fillColor)
	d.graphicContext.FillStringAt(text, backgroundPadding, (-top)+backgroundPadding)

	return left + right + (backgroundPadding * 2), -top + (backgroundPadding * 2)
}

func (d *DrawBuilder) DrawRoundedRectangle(width, height float64, roundedRadius float64) {
	d.graphicContext.SetFillColor(d.fillColor)
	d.graphicContext.SetLineWidth(d.strokeWidth)
	d.graphicContext.SetStrokeColor(d.strokeColor)

	draw2dkit.RoundedRectangle(
		d.graphicContext,
		0,
		0,
		width,
		height,
		roundedRadius,
		roundedRadius,
	)
	d.graphicContext.FillStroke()
}

func (d *DrawBuilder) DrawCircle(radius float64) {
	d.graphicContext.SetFillColor(d.fillColor)
	d.graphicContext.SetLineWidth(d.strokeWidth)
	d.graphicContext.SetStrokeColor(d.strokeColor)

	draw2dkit.Circle(
		d.graphicContext,
		radius,
		radius,
		radius,
	)
	d.graphicContext.FillStroke()
}

func (d *DrawBuilder) DrawTriangle(width, height float64) {
	d.graphicContext.SetFillColor(d.fillColor)
	d.graphicContext.SetLineWidth(d.strokeWidth)
	d.graphicContext.SetStrokeColor(d.strokeColor)

	Triangle(
		d.graphicContext,
		0,
		0,
		width,
		height,
	)
	d.graphicContext.FillStroke()
}
