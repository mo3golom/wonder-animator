package loader

import (
	"encoding/base64"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"strings"
)

const (
	urlHttp = "http"
)

func LoadImage(data string) (image.Image, error) {
	// Пробуем получить изображение из base64
	base64Reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))

	img, _, err := image.Decode(base64Reader)

	// Если получилось, возвращаем изображение
	if nil == err {
		return img, nil
	}

	// Возможно нам подпихнули ссылку, пробуем получить изображение из ссылки
	if strings.Contains(data, urlHttp) {
		response, err := http.Get(data)

		if err != nil {
			return nil, err
		}

		defer response.Body.Close()

		img, _, err = image.Decode(response.Body)

		if nil == err {
			return img, nil
		}
	}

	// Иначе пробуем получить изображение из файла
	file, err := os.Open(data)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err = image.Decode(file)

	return img, err
}
