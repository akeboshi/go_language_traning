package weightconv

// GToP はグラムをポンドへ変換します
func GToP(g Gram) Pound { return Pound(g / 0.0022046) }

// PToG はポンドをグラムへ変換します
func PToG(p Pound) Gram { return Gram(p * 0.0022046) }
