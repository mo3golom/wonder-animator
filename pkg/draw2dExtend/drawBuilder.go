package draw2dExtend

import (
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
)

type DrawBuilder struct {
	fontData               draw2d.FontData
	x, y                   float64
	fillColor, strokeColor ExtendColor
	rotateDegrees          float64
	rotatePointType        string
	scale                  float64
	opacity                float32
	graphicContext         *GraphicContext
	strokeWidth            float64
}

func NewDrawBuilder(fontData draw2d.FontData) *DrawBuilder {
	return &DrawBuilder{
		fillColor:       White,
		strokeColor:     Black,
		fontData:        fontData,
		scale:           1,
		rotatePointType: wonderEffectDTO.RotatePointCenter,
	}
}

func (d *DrawBuilder) SetX(x float64) *DrawBuilder {
	d.x = x

	return d
}

func (d *DrawBuilder) SetY(y float64) *DrawBuilder {
	d.y = y

	return d
}

func (d *DrawBuilder) SetRotate(rotateDegrees float64) *DrawBuilder {
	d.rotateDegrees = rotateDegrees

	return d
}

func (d *DrawBuilder) SetRotatePointType(rotatePointType string) *DrawBuilder {
	d.rotatePointType = rotatePointType

	return d
}

func (d *DrawBuilder) SetOpacity(opacity float32) *DrawBuilder {
	d.opacity = opacity

	return d
}

func (d *DrawBuilder) SetScale(scale float64) *DrawBuilder {
	d.scale = scale

	return d
}

func (d *DrawBuilder) SetFillColor(fillColor ExtendColor) *DrawBuilder {
	d.fillColor = fillColor

	return d
}

func (d *DrawBuilder) SetStrokeColor(strokeColor ExtendColor) *DrawBuilder {
	d.strokeColor = strokeColor

	return d
}

func (d *DrawBuilder) SetStrokeWidth(strokeWidth float64) *DrawBuilder {
	d.strokeWidth = strokeWidth

	return d
}

func (d *DrawBuilder) SetGraphicContext(graphicContext *GraphicContext) *DrawBuilder {
	d.graphicContext = graphicContext

	return d
}

func (d *DrawBuilder) DrawText(text string, fontSize float64, backgroundOptions *BackgroundOptions) (width, height float64) {
	var backgroundPadding float64 = 0

	if nil != backgroundOptions {
		backgroundPadding = backgroundOptions.Padding()
	}

	d.graphicContext.SetFontData(d.fontData)
	d.graphicContext.SetFontSize(fontSize)

	left, top, right, _ := d.graphicContext.GetStringBounds(text)

	d.graphicContext.RotateAndScale(
		d.rotateDegrees,
		d.scale,
		GetRotatePointByType(
			d.rotatePointType,
			d.x+backgroundPadding,
			d.y+backgroundPadding,
			left+right,
			-top,
		),
	)

	if nil != backgroundOptions {

		d.graphicContext.SetFillColor(backgroundOptions.fillColor.SetOpacity(d.opacity))
		d.graphicContext.SetLineWidth(0)
		draw2dkit.RoundedRectangle(
			d.graphicContext,
			d.x,
			d.y,
			d.x+right+(backgroundPadding*2),
			d.y+(-top)+(backgroundPadding*2),
			backgroundOptions.radius,
			backgroundOptions.radius,
		)
		d.graphicContext.FillStroke()
	}

	d.graphicContext.SetFillColor(d.fillColor.SetOpacity(d.opacity))
	d.graphicContext.FillStringAt(text, d.x+backgroundPadding, d.y+(-top)+backgroundPadding)

	return left + right + (backgroundPadding * 2), -top + (backgroundPadding * 2)
}

func (d *DrawBuilder) DrawRoundedRectangle(width, height float64, roundedRadius float64) {
	d.prepareGraphicContextForFigures(width, height)

	draw2dkit.RoundedRectangle(
		d.graphicContext,
		d.x,
		d.y,
		d.x+width,
		d.y+height,
		roundedRadius,
		roundedRadius,
	)
	d.graphicContext.FillStroke()
}

func (d *DrawBuilder) DrawCircle(radius float64) {

	diameter := radius * 2

	d.prepareGraphicContextForFigures(diameter, diameter)

	draw2dkit.Circle(
		d.graphicContext,
		d.x+radius,
		d.y+radius,
		radius,
	)
	d.graphicContext.FillStroke()
}

func (d *DrawBuilder) DrawTriangle(width, height float64) {
	d.prepareGraphicContextForFigures(width, height)

	Triangle(
		d.graphicContext,
		d.x,
		d.y,
		d.x+width,
		d.y+height,
	)
	d.graphicContext.FillStroke()
}

// Подготовка графического контекста для отрисовки фигур
func (d *DrawBuilder) prepareGraphicContextForFigures(width, height float64) {
	d.graphicContext.SetFillColor(d.fillColor.SetOpacity(d.opacity))
	d.graphicContext.SetLineWidth(d.strokeWidth)
	d.graphicContext.SetStrokeColor(d.strokeColor.SetOpacity(d.opacity))

	d.graphicContext.RotateAndScale(
		d.rotateDegrees,
		d.scale,
		GetRotatePointByType(
			d.rotatePointType,
			d.x,
			d.y,
			width,
			height,
		),
	)
}
