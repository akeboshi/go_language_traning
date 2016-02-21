package lengthconv

import "fmt"

// Feet フィート
type Feet float64

// Meter メートル
type Meter float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
