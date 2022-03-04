package sorting

// Sorting is a sorting
type Sorting struct {
	Field     string
	Direction Direction
}

// Func is a sorting func
type Func func([]Sorting) []Sorting
