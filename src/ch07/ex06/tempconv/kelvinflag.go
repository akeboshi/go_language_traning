//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package tempconv

import (
	"ch02/ex01/tempconv"
	"flag"
	"fmt"
)

type kelvinFlag struct{ tempconv.Kelvin }

func (f *kelvinFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "K":
		f.Kelvin = tempconv.Kelvin(value)
		return nil
	case "C", "℃":
		f.Kelvin = tempconv.CToK(tempconv.Celsius(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func KelvinFlag(name string, value tempconv.Kelvin, usage string) *tempconv.Kelvin {
	f := kelvinFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Kelvin
}
