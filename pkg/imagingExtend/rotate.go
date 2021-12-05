package imagingExtend

import (
	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
	"image"
	"math"
)

func RotateAround(src image.Image, angleDegrees float64, x, y float64) image.Image {
	radians := -(angleDegrees * (math.Pi / 180.0))
	cos := math.Cos(radians)
	sin := math.Sin(radians)

	aff := f64.Aff3{
		cos, -sin, x - x*cos + y*sin,
		sin, cos, y - x*sin - y*cos,
	}

	dst := image.NewRGBA(calcBounds(aff, src.Bounds()))
	draw.BiLinear.Transform(dst, aff, src, src.Bounds(), draw.Src, nil)

	return dst
}

func calcBounds(aff f64.Aff3, b image.Rectangle) image.Rectangle {
	var min, max image.Point

	points := [...]image.Point{
		{b.Min.X, b.Min.Y},
		{b.Max.X - 1, b.Min.Y},
		{b.Min.X, b.Max.Y - 1},
		{b.Max.X - 1, b.Max.Y - 1},
	}

	for i, p := range points {
		x0 := float64(p.X) + 0.5
		y0 := float64(p.Y) + 0.5
		x := aff[0]*x0 + aff[1]*y0 + aff[2]
		y := aff[3]*x0 + aff[4]*y0 + aff[5]
		pmin := image.Point{X: int(math.Floor(x)), Y: int(math.Floor(y))}
		pmax := image.Point{X: int(math.Ceil(x)), Y: int(math.Ceil(y))}

		if i == 0 {
			min = pmin
			max = pmax

			continue
		}

		if min.X > pmin.X {
			min.X = pmin.X
		}

		if min.Y > pmin.Y {
			min.Y = pmin.Y
		}

		if max.X < pmax.X {
			max.X = pmax.X
		}

		if max.Y < pmax.Y {
			max.Y = pmax.Y
		}
	}

	return image.Rectangle{Min: min, Max: max}
}
