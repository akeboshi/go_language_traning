//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package tempconv

import "fmt"

// Celsius 摂氏
type Celsius float64

// Fahrenheit 華氏
type Fahrenheit float64

// Kelvin 絶対温度
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
