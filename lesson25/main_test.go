package main

import (
	"errors"
	"testing"
)

func TestSayHello(t *testing.T) {
	want := "Hello"
	res := SayHello()
	if res != want {
		t.Errorf("incorrect output: %s; want: %s", res, want)
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a            float64
		b            float64
		wantedResult float64
		wantedError  error
	}{
		{
			a:            4,
			b:            2,
			wantedResult: 2,
			wantedError:  nil,
		},
		{
			a:            10,
			b:            2,
			wantedResult: 5,
			wantedError:  nil,
		},
		{
			a:            10,
			b:            0,
			wantedResult: 0,
			wantedError:  ErrDivisionByZero,
		},
	}

	for _, test := range tests {
		res, err := divide(test.a, test.b)

		if res != test.wantedResult {
			t.Errorf("got: %f; want: %f", res, test.wantedResult)
		}

		if !errors.Is(err, test.wantedError) {
			t.Errorf("got: %v; wanted: %v", err, test.wantedError)
		}
	}
}
