package imageutil

import (
	"image"
	"image/color"
)

func Grayscale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(originColor)
			grayImg.Set(x, y, grayColor)
		}
	}
	return grayImg
}
