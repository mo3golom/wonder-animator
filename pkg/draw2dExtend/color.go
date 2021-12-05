package draw2dExtend

import "image/color"

const alphaMax = 255

const (
	opacityMax = 1
	opacityMin = 0
)

var (
	White = &ExtendRGBA{R: 255, G: 255, B: 255, A: 255}
	Black = &ExtendRGBA{R: 0, G: 0, B: 0, A: 255}
)

type ExtendColor interface {
	color.Color
	SetOpacity(alpha float32) ExtendColor
}

type ExtendRGBA struct {
	R, G, B, A uint8
}

func (c *ExtendRGBA) SetOpacity(alpha float32) ExtendColor {
	if opacityMax < alpha {
		alpha = opacityMax
	}

	if opacityMin > alpha {
		alpha = opacityMin
	}

	convertAlpha := uint8(alphaMax * alpha)

	if c.R > convertAlpha {
		c.R = convertAlpha
	}

	if c.G > convertAlpha {
		c.G = convertAlpha
	}

	if c.B > convertAlpha {
		c.B = convertAlpha
	}

	c.A = convertAlpha

	return c
}

func (c *ExtendRGBA) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

func ParseHexColor(s string) *ExtendRGBA {
	c := &ExtendRGBA{}
	c.A = 0xff

	if s[0] != '#' {
		return c
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
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])

	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	}

	return c
}
