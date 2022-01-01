package draw2dExtend

import "image/color"

func ParseHexColor(s string) (color color.RGBA) {
	color.A = 0xff

	if s[0] != '#' {
		return color
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'

		case b >= 'a' && b <= 'f':
			return b - 'a' + 10

		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}

		return 0
	}

	switch len(s) {
	case 7:
		color.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		color.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		color.B = hexToByte(s[5])<<4 + hexToByte(s[6])

	case 4:
		color.R = hexToByte(s[1]) * 17
		color.G = hexToByte(s[2]) * 17
		color.B = hexToByte(s[3]) * 17
	}

	return color
}
