package kernel

import "math"

func CombineGradient(gx, gy [][]float64) [][]float64 {
	h := len(gx)
	if h == 0 {
		return nil
	}
	w := len(gx[0])
	if len(gy) != h {
		return nil
	}
	if len(gy[0]) != w {
		return nil
	}
	result := make([][]float64, h)
	for i := range result {
		result[i] = make([]float64, w)
	}
	for y := 0; y < h; y++ {

		for x := 0; x < w; x++ {
			mag := math.Sqrt(
				gx[y][x]*gx[y][x] +
					gy[y][x]*gy[y][x],
			)
			result[y][x] = mag
		}
	}
	return result
}
