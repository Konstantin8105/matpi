package matpi

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/Konstantin8105/errors"

	"gonum.org/v1/gonum/mat"
)

// Config is configuration of matrix picture
type Config struct {
	// color of positive value
	PositiveColor color.RGBA

	// color of negative value
	NegativeColor color.RGBA

	// color of zero value
	ZeroColor color.RGBA

	// scale of picture
	Scale int
}

// DefaultConfig is default configuration
func DefaultConfig() Config {
	return Config{
		PositiveColor: color.RGBA{255, 0, 0, 0xff},    // RED
		NegativeColor: color.RGBA{25, 181, 254, 0xff}, // BLUE
		ZeroColor:     color.RGBA{255, 255, 0, 0xff},  // YELLOW
		Scale:         1,
	}
}

// Convert matrix 'gonum.mat.Matrix' to PNG picture file with filename.
// Color of picture pixel in according to `config`.
func Convert(m mat.Matrix, filename string, config Config) error {

	// check input data
	et := errors.New("check input data")
	if config.Scale < 0 {
		_ = et.Add(fmt.Errorf("not acceptable scale less zero: %d", config.Scale))
	}
	if config.Scale == 0 {
		_ = et.Add(fmt.Errorf("not acceptable zero scale"))
	}
	if filename == "" {
		_ = et.Add(fmt.Errorf("filename is empty"))
	}
	if m == nil {
		_ = et.Add(fmt.Errorf("matrix is nil"))
	}

	if et.IsError() {
		return et
	}

	// generate picture
	r, c := m.Dims()

	img := image.NewRGBA(image.Rect(0, 0, r*config.Scale, c*config.Scale))

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			p := m.At(i, j)

			// color identification
			var color color.RGBA
			switch {
			case p > math.SmallestNonzeroFloat64:
				color = config.PositiveColor
			case p < -math.SmallestNonzeroFloat64:
				color = config.NegativeColor
			default:
				color = config.ZeroColor
			}

			// iteration by pixels
			for k := 0; k < config.Scale; k++ {
				for g := 0; g < config.Scale; g++ {
					img.Set(i*config.Scale+k, j*config.Scale+g, color)
				}
			}
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	return png.Encode(file, img)
}
