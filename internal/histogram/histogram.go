package histogram

import "image"

func Compute(img *image.Gray) [256]int {
	var hist [256]int
	bound := img.Bounds()
	for y := bound.Min.Y; y < bound.Max.Y; y++ {

		for x := bound.Min.X; x < bound.Max.X; x++ {
			pixel := img.GrayAt(x, y).Y
			hist[pixel]++
		}

	}
	return hist
}
