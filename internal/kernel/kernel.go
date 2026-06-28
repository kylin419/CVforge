package kernel

type Kernel struct {
	Data   [][]float64
	Factor float64
	Bias   float64
}

func (k Kernel) Width() int {
	if len(k.Data[0]) == 0 {
		return 0
	}
	return len(k.Data[0])
}
func (k Kernel) Height() int {
	return len(k.Data)
}
