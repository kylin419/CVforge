package kernel

import (
	"fmt"
	"math"
)

func GaussianKernel(size int, sigma float64) Kernel {
	if size < 3 || size%2 == 0 {
		panic("kernel size must be odd and >= 3")
	}
	if sigma <= 0 {
		panic("sigma must be >0")
	}
	center := size / 2
	data := make([][]float64, size)
	for i := range data {
		data[i] = make([]float64, size)
	}
	var sum float64

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			dx := float64(x - center)
			dy := float64(y - center)
			value := math.Exp(-(dx*dx+dy*dy)/(2*sigma*sigma)) / (2 * math.Pi * sigma * sigma)
			data[y][x] = value
			sum += value
		}
	}
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			data[y][x] /= sum
		}
	}
	return Kernel{
		Name: fmt.Sprintf(
			"Gaussian %dx%d",
			size,
			size,
		),
		Data: data,

		Factor: 1,
		Bias:   0,
	}
}
