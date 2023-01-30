package calculator

import (
	"math"
	"testing"
)

func TestExpressionReturnsAValidValue(t *testing.T) {

	got := Calc("42+8+25*2")
	want := float64(100)

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}

}

func TestExpressionReturnsANonValidValue(t *testing.T) {

	got := Calc("1/0")
	want := math.NaN()

	if !math.IsNaN(got) || !math.IsNaN(want) {
		t.Errorf("got %f, wanted %f", got, want)
	}

}
