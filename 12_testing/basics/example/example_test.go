package main

import (
	"testing"
	"time"
)

func TestCalcAreaSuccess(t *testing.T) {
	result, err := CalcArea(3, 5)
	if err != nil {
		t.Error("Expected CalcArea(3, 5) to succeed")
	} else if result != 15 {
		t.Errorf("CalcArea(3, 5) returned %d. Expected 15", result)
	}
}

func TestCalcAreaFail(t *testing.T) {
	_, err := CalcArea(-3, 6)
	if err == nil {
		t.Error("Expected CalcArea(-3, 6) to return an error")
	}

	if err.Error() != errorMessage {
		t.Error("Expected error to be: " + errorMessage)
	}
}

func TestCalcAreaViaTable(t *testing.T) {
	var tests = []struct {
		width    int
		height   int
		expected int
	}{
		{1, 1, 1},
		{5, 6, 30},
		{1, 99, 99},
		{7, 6, 42},
		//{-8, -8, 64},
		//{0, 5, 0},
	}

	for _, test := range tests {
		w := test.width
		h := test.height
		r, err := CalcArea(w, h)
		if err != nil {
			t.Errorf("CalcArea(%d, %d) returned an error", w, h)
		} else if r != test.expected {
			t.Errorf("CalcArea(%d, %d) returned %d. Expected %d",
				w, h, r, test.expected)
		}
	}
}

func TestCalcAreaInParallel(t *testing.T) {
	var tests = []struct {
		width    int
		height   int
		expected int
	}{
		{1, 1, 1},
		{5, 6, 30},
		{1, 99, 99},
		{7, 6, 42},
	}

	for _, test := range tests {
		t.Run("", func(tt *testing.T) {
			tt.Parallel()
			time.Sleep(time.Second)
			w := test.width
			h := test.height
			r, err := CalcArea(w, h)
			if err != nil {
				tt.Errorf("CalcArea(%d, %d) returned an error", w, h)
			} else if r != test.expected {
				tt.Errorf("CalcArea(%d, %d) returned %d. Expected %d",
					w, h, r, test.expected)
			}
		})
	}
}
