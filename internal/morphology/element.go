package morphology

type Element struct {
	Name string
	Data [][]bool
}

func (e Element) Width() int {
	if len(e.Data) == 0 {
		return 0
	}
	return len(e.Data[0])
}
func (e Element) Height() int {
	return len(e.Data)
}

func (e Element) Valid() bool {
	if len(e.Data) == 0 {
		return false
	}

	w := len(e.Data[0])
	if w == 0 {
		return false
	}

	for _, row := range e.Data {
		if len(row) != w {
			return false
		}
	}

	return true
}
