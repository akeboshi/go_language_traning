package tempconv

import (
	"ch02/ex01/tempconv"
	"testing"
)

func TestTempconv(t *testing.T) {
	var data float64 = 100
	var expectedCToF tempconv.Fahrenheit = 212
	actualCToF := tempconv.CToF(tempconv.Celsius(data))
	if expectedCToF != actualCToF {
		t.Errorf("CToF: got %f want %f data is %f", actualCToF, expectedCToF, data)
	}

}
