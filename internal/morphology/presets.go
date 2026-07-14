package morphology

var Square3x3 = Element{
	Name: "Square 3x3",
	Data: [][]bool{
		{true, true, true},
		{true, true, true},
		{true, true, true},
	},
}
var Cross3x3 = Element{
	Name: "Cross 3x3",
	Data: [][]bool{
		{false, true, false},
		{true, true, true},
		{false, true, false},
	},
}

var Square5x5 = Element{
	Name: "Square 5x5",
	Data: [][]bool{
		{true, true, true, true, true},
		{true, true, true, true, true},
		{true, true, true, true, true},
		{true, true, true, true, true},
		{true, true, true, true, true},
	},
}
