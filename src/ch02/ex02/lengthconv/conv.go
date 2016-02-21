package lengthconv

// MToF はメートルをフィートへ変換します
func MToF(m Meter) Feet { return Feet(m / 0.3048) }

// FToM はフィートからメートルへ変換します
func FToM(f Feet) Meter { return Meter(f * 0.3048) }
