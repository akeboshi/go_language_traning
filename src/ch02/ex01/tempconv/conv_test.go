//Copyright (c) 2016 by akeboshi. All Rights Reserved.
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

	data = 100
	var expectedCToK tempconv.Kelvin = 373.15
	actualCToK := tempconv.CToK(tempconv.Celsius(data))
	if expectedCToK != actualCToK {
		t.Errorf("CToK: got %f want %f data is %f", actualCToK, expectedCToK, data)
	}

	data = 41
	var expectedFToC tempconv.Celsius = 5
	actualFToC := tempconv.FToC(tempconv.Fahrenheit(data))
	if expectedFToC != actualFToC {
		t.Errorf("FToC: got %f want %f data is %f", actualFToC, expectedFToC, data)
	}

	data = 41
	var expectedFToK tempconv.Kelvin = 278.15
	actualFToK := tempconv.FToK(tempconv.Fahrenheit(data))
	if expectedFToK != actualFToK {
		t.Errorf("FToK: got %f want %f data is %f", actualFToK, expectedFToK, data)
	}

	data = 100
	var expectedKToC tempconv.Celsius = -173.15
	actualKToC := tempconv.KToC(tempconv.Kelvin(data))
	if expectedKToC-actualKToC > 0.0000001 {
		t.Errorf("KToC: got %f want %f data is %f", actualKToC, expectedKToC, data)
	}

	data = 278.15
	var expectedKToF tempconv.Fahrenheit = 41
	actualKToF := tempconv.KToF(tempconv.Kelvin(data))
	if expectedKToF != actualKToF {
		t.Errorf("KToF: got %f want %f data is %f", actualKToF, expectedKToF, data)
	}
}
