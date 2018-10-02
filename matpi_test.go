package matpi

import (
	"bytes"
	"io/ioutil"
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

	act := "testdata/diagonal.jpg"
	exp := "testdata/diagonal_expect.jpg"

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

	act := "testdata/diagonal2.jpg"
	exp := "testdata/diagonal2_expect.jpg"

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

	err := Convert(m, "result.jpg")
	if err != nil {
		return
	}
}
