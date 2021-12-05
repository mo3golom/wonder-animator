package draw2dExtend

type BackgroundOptions struct {
	padding, radius float64
	fillColor       ExtendColor
}

func NewBackgroundOptions() *BackgroundOptions {
	return &BackgroundOptions{
		padding:   0,
		radius:    0,
		fillColor: Black,
	}
}

func (b *BackgroundOptions) Padding() float64 {
	return b.padding
}

func (b *BackgroundOptions) SetPadding(padding float64) *BackgroundOptions {
	b.padding = padding

	return b
}

func (b *BackgroundOptions) FillColor() ExtendColor {
	return b.fillColor
}

func (b *BackgroundOptions) SetFillColor(fillColor ExtendColor) *BackgroundOptions {
	b.fillColor = fillColor

	return b
}

func (b *BackgroundOptions) Radius() float64 {
	return b.radius
}

func (b *BackgroundOptions) SetRadius(radius float64) *BackgroundOptions {
	b.radius = radius

	return b
}
