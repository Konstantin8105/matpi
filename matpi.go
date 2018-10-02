package matpi

import (
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"

	"gonum.org/v1/gonum/mat"
)

var (
	positiveColor = color.RGBA{0, 100, 0, 0xff}
	negativeColor = color.RGBA{80, 80, 200, 0xff}
	zeroColor     = color.RGBA{255, 250, 205, 0xff}
)

// Convert matrix 'gonum.mat.Matrix' to JPEG picture file with filename.
// Non-zero matrix element is black.
func Convert(m mat.Matrix, filename string) (err error) {

	r, c := m.Dims()

	img := image.NewRGBA(image.Rect(0, 0, r, c))

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			p := m.At(i, j)
			if math.Abs(p) == 0.0 {
				img.Set(i, j, zeroColor)
				continue
			}
			if p > 0 {
				img.Set(i, j, positiveColor)
			} else {
				img.Set(i, j, negativeColor)
			}
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer func() { _ = file.Close() }()

	return jpeg.Encode(file, img, &jpeg.Options{Quality: 80})
}
