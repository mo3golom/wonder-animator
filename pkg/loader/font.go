package loader

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"io/ioutil"
)

func Load(path string) (font *truetype.Font) {
	fontBytes, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	f, err := freetype.ParseFont(fontBytes)

	if err != nil {
		panic(err)
	}

	return f
}

func LoadAndRegister(name string, path string) *draw2d.FontData {
	font := Load(path)
	fontData := &draw2d.FontData{
		Name:   name,
		Family: draw2d.FontFamilyMono,
		Style:  draw2d.FontStyleNormal,
	}

	draw2d.RegisterFont(
		*fontData,
		font,
	)

	return fontData
}
