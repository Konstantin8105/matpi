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
	// read file
	f, err := os.Open("./testdata/1544775921503561857.modal")
	if err != nil {
		t.Fatal(err)
	}

	// read entries
	m := mat.NewDense(204, 204, nil)
	var n int
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
		m.Set(i, j, x)
	}
	config := matpi.NewConfig()
	config.Scale = 4
	err = matpi.Convert(m, "./testdata/big.png", config)
	if err != nil {
		t.Error(err)
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
