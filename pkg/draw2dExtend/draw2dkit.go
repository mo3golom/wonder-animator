package draw2dExtend

import "github.com/llgcode/draw2d"

// Triangle draws a triangle using a path between (x1,y1) and (x2,y2)
func Triangle(path draw2d.PathBuilder, x1, y1, x2, y2 float64) {
	width := x2 - x1

	path.MoveTo(x1+(width/2), y1)
	path.LineTo(x2, y2)
	path.LineTo(x1, y2)
	path.LineTo(x1+(width/2), y1)
	path.Close()
}
