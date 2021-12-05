package imagingExtend

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
)

func Opacity(src image.Image, opacity float64) image.Image {
	size := src.Bounds()
	background := imaging.New(size.Dx(), size.Dy(), color.Transparent)

	return imaging.OverlayCenter(background, src, opacity)
}
