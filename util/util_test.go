package util

import "testing"

func TestSum(t *testing.T) {
	res := Sum(34, 90)
	if res != 124 {
		t.Fatal("34 + 90 != 124")
	}
}

func TestSub(t *testing.T) {
	res := Sub(35, 9)
	if res != 26 {
		t.Fatal("35 - 9 != 26")
	}
}

func TestMul(t *testing.T) {
	res := Mul(23, 2)
	if res != 46 {
		t.Fatal("23 * 2 != 46")
	}
}

func TestDiv(t *testing.T) {
	res := Div(24, 2)
	if res != 12 {
		t.Fatal("24 / 2 != 12")
	}
}
