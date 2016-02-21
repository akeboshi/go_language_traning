package weightconv

import "fmt"

// Gram グラム
type Gram float64

// Pound ポンド
type Pound float64

func (g Gram) String() string  { return fmt.Sprintf("%gg", g) }
func (p Pound) String() string { return fmt.Sprintf("%glb", p) }
