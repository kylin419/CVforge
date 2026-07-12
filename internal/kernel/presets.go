package kernel

var (
	Sharpen = Kernel{
		Name: "Sharpen",
		Data: [][]float64{
			{0, -1, 0},
			{-1, 5, -1},
			{0, -1, 0},
		},
		Factor: 1,
		Bias:   0,
	}
	Emboss = Kernel{
		Name: "Emboss",
		Data: [][]float64{
			{-2, -1, 0},
			{-1, 1, 1},
			{0, 1, 2},
		},
		Factor: 1,
		Bias:   128,
	}
	SobelX = Kernel{
		Name: "Sobel X",
		Data: [][]float64{
			{-1, 0, 1},
			{-2, 0, 2},
			{-1, 0, 1},
		},
		Factor: 1,
	}
	SobelY = Kernel{
		Name: "Sobel Y",
		Data: [][]float64{
			{-1, -2, -1},
			{0, 0, 0},
			{1, 2, 1},
		},
		Factor: 1,
	}
	BoxBlur3x3 = Kernel{
		Name: "Box Blur 3x3",
		Data: [][]float64{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
		Factor: 1.0 / 9.0,
	}
	Laplacian = Kernel{
		Name: "Laplacian",
		Data: [][]float64{
			{0, 1, 0},
			{1, -4, 1},
			{0, 1, 0},
		},
		Factor: 1,
	}
)
