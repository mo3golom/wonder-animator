package draw2dExtend

import "github.com/mo3golom/wonder-effects/wonderEffectDTO"

func GetRotatePointByType(pointType string, leftTopX, leftTopY float64, width, height float64) Point {
	// Точка по умолчанию - центр
	point := Point{
		X: leftTopX + (width / 2),
		Y: leftTopY + (height / 2),
	}

	switch pointType {
	case wonderEffectDTO.RotatePointerLeft:
		point.X = leftTopX
		point.Y = leftTopY + (height / 2)

	case wonderEffectDTO.RotatePointTop:
		point.X = leftTopX + (width / 2)
		point.Y = leftTopY

	case wonderEffectDTO.RotatePointRight:
		point.X = leftTopX + width
		point.Y = leftTopY + (height / 2)

	case wonderEffectDTO.RotatePointBottom:
		point.X = leftTopX + (width / 2)
		point.Y = leftTopY + height

	case wonderEffectDTO.RotatePointLeftTop:
		point.X = leftTopX
		point.Y = leftTopY

	case wonderEffectDTO.RotatePointRightTop:
		point.X = leftTopX + width
		point.Y = leftTopY

	case wonderEffectDTO.RotatePointLeftBottom:
		point.X = leftTopX
		point.Y = leftTopY + height

	case wonderEffectDTO.RotatePointRightBottom:
		point.X = leftTopX + width
		point.Y = leftTopY + height
	}

	return point
}
