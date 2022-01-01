package draw2dExtend

import "image/color"

type BackgroundOptions struct {
	Padding, Radius float64
	FillColor       color.RGBA
}

func NewBackgroundOptions() *BackgroundOptions {
	return &BackgroundOptions{
		Padding:   0,
		Radius:    0,
		FillColor: ParseHexColor("#ffffff"),
	}
}
