package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Konstantin8105/matpi"
	"gonum.org/v1/gonum/mat"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stdout, "please enter filename in triplet format\n")
		return
	}
	for i := 1; i < len(os.Args); i++ {
		name := os.Args[i]
		png, err := fromTriplets(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", name, err)
			continue
		}
		fmt.Fprintf(os.Stdout, "convert %v to %v\n", name, png)
	}
}

func fromTriplets(filename string) (pngFn string, err error) {
	type triplet struct {
		i, j int
		x    float64
	}

	// read file
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	// read entries
	var n int
	var tr []triplet
	for {
		var i, j int
		var x float64

		n, err = fmt.Fscanf(f, "%d %d %f\n", &i, &j, &x)
		if err == io.EOF {
			break
		}
		if n != 3 {
			return "", fmt.Errorf("scan more then 3 variables")
		}
		if err != nil {
			return "", fmt.Errorf("cannot scan: %v", err)
		}
		tr = append(tr, triplet{i, j, x})
	}
	size := 0
	for i := range tr {
		if size < tr[i].i {
			size = tr[i].i
		}
		if size < tr[i].j {
			size = tr[i].j
		}
	}
	size++
	m := mat.NewDense(size, size, nil)
	for i := range tr {
		m.Set(tr[i].i, tr[i].j, tr[i].x)
	}
	config := matpi.NewConfig()
	config.Scale = 3
	pngFn = filename + ".png"
	err = matpi.Convert(m, pngFn, config)
	if err != nil {
		return
	}
	return
}
