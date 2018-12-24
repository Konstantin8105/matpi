package matpi_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Konstantin8105/matpi"
	"gonum.org/v1/gonum/mat"
)

func sameFiles(act, exp string) (bool, error) {
	a, err := ioutil.ReadFile(act)
	if err != nil {
		return false, err
	}
	e, err := ioutil.ReadFile(exp)
	if err != nil {
		return false, err
	}
	if !bytes.Equal(a, e) {
		return false, nil
	}
	return true, nil
}

func TestConvert(t *testing.T) {
	m := mat.NewDense(100, 80, nil)
	for i := 0; i < 80; i++ {
		m.Set(i, i, 1.0)
	}

	act := "testdata/diagonal.png"
	exp := "testdata/diagonal_expect.png"

	err := matpi.Convert(m, act, matpi.NewConfig())
	if err != nil {
		t.Fatal(err)
	}

	result, err := sameFiles(act, exp)
	if err != nil {
		t.Fatal(err)
	}
	if !result {
		t.Fatal("Pictures are not the same")
	}

}

func TestConvert2(t *testing.T) {
	m := mat.NewDense(100, 80, nil)
	for i := 0; i < 80; i++ {
		m.Set(i, i, 1.0)
		m.Set(i, 79-i, -1.0)
	}

	act := "testdata/diagonal2.png"
	exp := "testdata/diagonal2_expect.png"

	err := matpi.Convert(m, act, matpi.NewConfig())
	if err != nil {
		t.Fatal(err)
	}

	result, err := sameFiles(act, exp)
	if err != nil {
		t.Fatal(err)
	}
	if !result {
		t.Fatal("Pictures are not the same")
	}
}

func TestConvertFail(t *testing.T) {
	m := mat.NewDense(100, 80, nil)
	for i := 0; i < 80; i++ {
		m.Set(i, i, 1.0)
	}
	err := matpi.Convert(m, "/////", matpi.NewConfig())
	if err == nil {
		t.Fatal("Error is empty")
	}

}

func ExampleConvert() {
	m := mat.NewDense(100, 80, nil)
	for i := 0; i < 80; i++ {
		m.Set(i, i, 1.0)
		m.Set(i, 79-i, -1.0)
	}

	err := matpi.Convert(m, "result.png", matpi.NewConfig())
	if err != nil {
		panic(err)
	}
}

func TestTriplets(t *testing.T) {
	tcs := []struct {
		filename    string
		pngFilename string
	}{
		{
			filename:    "./testdata/lu",
			pngFilename: "./testdata/lu.png",
		},
		{
			filename:    "./testdata/mass",
			pngFilename: "./testdata/mass.png",
		},
		{
			filename:    "./testdata/modal",
			pngFilename: "./testdata/modal.png",
		},
	}

	type triplet struct {
		i, j int
		x    float64
	}

	for _, tc := range tcs {
		// read file
		f, err := os.Open(tc.filename)
		if err != nil {
			t.Fatal(err)
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
				t.Fatalf("scan more then 3 variables")
			}
			if err != nil {
				t.Fatalf("cannot scan: %v", err)
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
		err = matpi.Convert(m, tc.pngFilename, config)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestWrong(t *testing.T) {
	errs := []error{
		func() error {
			err := matpi.Convert(nil, "", nil)
			return err
		}(),
		func() error {
			err := matpi.Convert(nil, "", &matpi.Config{})
			return err
		}(),
		func() error {
			c := matpi.Config{}
			c.Scale = -1
			err := matpi.Convert(nil, "", &c)
			return err
		}(),
	}

	for i := range errs {
		if errs[i] == nil {
			t.Errorf("not fail for case %d", i)
		}
		t.Log(errs[i])
	}
}
