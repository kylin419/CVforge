package kernel

func Clamp(v float64) uint8 {
	switch {
	case v < 0:
		return 0
	case v > 255:
		return 255
	default:
		return uint8(v)
	}
}

func (k Kernel) Valid() bool {
	if len(k.Data) == 0 {
		return false
	}
	w := len(k.Data[0])
	if w == 0 {
		return false
	}
	for _, row := range k.Data {
		if len(row) != w {
			return false
		}
	}
	return true
}

func Normalize(value, factor, bias float64) float64 {
	return value*factor + bias
}
