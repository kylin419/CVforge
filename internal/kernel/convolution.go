package kernel

import (
	"image"
	"image/color"
)

func Apply(img image.Image, k Kernel) image.Image {
	if !k.Valid() {
		return img
	}
	bound := img.Bounds()
	dst := image.NewRGBA(bound)
	kw := k.Width()
	kh := k.Height()
	offsetX := kw / 2
	offsetY := kh / 2
	for y := bound.Min.Y; y < bound.Max.Y; y++ {
		for x := bound.Min.X; x < bound.Max.X; x++ {
			var rSum float64
			var gSum float64
			var bSum float64
			for ky := 0; ky < kh; ky++ {
				for kx := 0; kx < kw; kx++ {
					nx := x + kx - offsetX
					ny := y + ky - offsetY
					if nx < bound.Min.X ||
						nx >= bound.Max.X ||
						ny < bound.Min.Y ||
						ny >= bound.Max.Y {
						continue
					}
					r, g, b, _ := img.At(nx, ny).RGBA()
					weight := k.Data[ky][kx]
					rSum += float64(r) / 256.0 * weight
					gSum += float64(g) / 256.0 * weight
					bSum += float64(b) / 256.0 * weight
				}
			}
			dst.Set(x, y, color.RGBA{R: Clamp(Normalize(rSum, k.Factor, k.Bias)), G: Clamp(Normalize(gSum, k.Factor, k.Bias)), B: Clamp(Normalize(bSum, k.Factor, k.Bias)), A: 255})
		}
	}
	return dst
}

func ApplyRaw(img *image.Gray, k Kernel) [][]float64 {
	if !k.Valid() {
		return nil
	}
	bound := img.Bounds()
	kw := k.Width()
	kh := k.Height()

	offsetX := kw / 2
	offsetY := kh / 2

	result := make([][]float64, bound.Dy())
	for i := range result {
		result[i] = make([]float64, bound.Dx())
	}
	for y := bound.Min.Y; y < bound.Max.Y; y++ {

		for x := bound.Min.X; x < bound.Max.X; x++ {
			var sum float64
			for ky := 0; ky < kh; ky++ {

				for kx := 0; kx < kw; kx++ {
					nx := x + kx - offsetX
					ny := y + ky - offsetY
					if nx < bound.Min.X ||
						nx >= bound.Max.X ||
						ny < bound.Min.Y ||
						ny >= bound.Max.Y {
						continue
					}
					pixel := img.GrayAt(nx, ny).Y
					weight := k.Data[ky][kx]
					sum += float64(pixel) * weight

				}

			}
			sum = Normalize(sum, k.Factor, k.Bias)
			result[y-bound.Min.Y][x-bound.Min.X] = sum
		}

	}
	return result
}
