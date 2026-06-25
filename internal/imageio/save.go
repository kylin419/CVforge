package imageio

import (
	"image"
	"image/jpeg"
	"os"
)

func Save(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return jpeg.Encode(file, img, nil)
}
