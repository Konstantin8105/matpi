package matpi

import (
	"bytes"
	"io/ioutil"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestConvert(t *testing.T) {
	m := mat.NewDense(100, 80, nil)
	for i := 0; i < 80; i++ {
		m.Set(i, i, 1.0)
	}

	act := "testdata/diagonal.jpg"
	exp := "testdata/diagonal_expect.jpg"

	err := Convert(m, act)
	if err != nil {
		t.Fatal(err)
	}

	a, err := ioutil.ReadFile(act)
	if err != nil {
		t.Fatal(err)
	}

	e, err := ioutil.ReadFile(exp)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(a, e) {
		t.Fatal("Pictures is not same")
	}
}
