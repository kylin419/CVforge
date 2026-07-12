package kernel

import (
	"image"
	"image/color"
)

func ToGrayImage(data [][]float64) *image.Gray {
	h := len(data)
	if h == 0 || len(data[0]) == 0 {
		return image.NewGray(image.Rect(0, 0, 0, 0))
	}
	w := len(data[0])
	img := image.NewGray(image.Rect(0, 0, w, h))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.SetGray(x, y, color.Gray{Y: Clamp(data[y][x])})
		}
	}
	return img
}
