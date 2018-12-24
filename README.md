[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/matpi/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/matpi?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/matpi)](https://goreportcard.com/report/github.com/Konstantin8105/matpi)
[![GoDoc](https://godoc.org/github.com/Konstantin8105/matpi?status.svg)](https://godoc.org/github.com/Konstantin8105/matpi)
![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)

# matpi

Convert matrix from `gonum.mat.Matrix` to PNG picture.
Non-zero matrix element is black.

Example:

```golang
	m := mat.NewDense(100, 80, nil)
	for i := 0; i < 80; i++ {
		m.Set(i, i, 1.0)
	}

	err := matpi.Convert(m, "result.png", matpi.NewConfig())
	if err != nil {
		return
	}
```

![Diagonal](https://raw.githubusercontent.com/Konstantin8105/matpi/master/testdata/diagonal_expect.png)

![Symmetrical LU matrix](https://raw.githubusercontent.com/Konstantin8105/matpi/master/testdata/lu.png)

![Mass Matrix](https://raw.githubusercontent.com/Konstantin8105/matpi/master/testdata/mass.png)

![Matrix for modal calculation](https://raw.githubusercontent.com/Konstantin8105/matpi/master/testdata/modal.png)
