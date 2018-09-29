package matpi

import (
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"

	"gonum.org/v1/gonum/mat"
)

// Convert matrix 'gonum.mat.Matrix' to JPEG picture file with filename.
func Convert(m mat.Matrix, filename string) (err error) {
	r, c := m.Dims()

	img := image.NewRGBA(image.Rect(0, 0, r, c))

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if math.Abs(m.At(i, j)) == 0.0 {
				img.Set(i, j, color.RGBA{255, 250, 205, 0xff})
				continue
			}
			img.Set(i, j, color.RGBA{0, 100, 0, 0xff})
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer func() { _ = file.Close() }()

	return jpeg.Encode(file, img, &jpeg.Options{80})
}
