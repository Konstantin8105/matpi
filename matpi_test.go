package matpi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

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

	err := Convert(m, act)
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

	err := Convert(m, act)
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
	err := Convert(m, "/////")
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

	err := Convert(m, "result.png")
	if err != nil {
		return
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
	for {
		var i, j int
		var x float64

		n, err := fmt.Fscanf(f, "%d %d %f\n", &i, &j, &x)
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
	err = Convert(m, "./testdata/big.png")
	if err != nil {
		t.Error(err)
	}
}
