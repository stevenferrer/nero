package sorting

// Sorting is a sorting
type Sorting struct {
	Field     string
	Direction Direction
}

type Sortings []Sorting

// Func is a sorting func
type Func func(Sortings) Sortings

func Build(sortFuncs ...Func) Sortings {
	sortings := make(Sortings, 0, len(sortFuncs))
	for _, sortFunc := range sortFuncs {
		sortings = sortFunc(sortings)
	}

	return sortings
}

// Direction is a sort direction
type Direction int

const (
	// Asc ascending sort direction
	Asc Direction = iota
	// Desc descending sort direction
	Desc
)

func (d Direction) String() string {
	return [...]string{
		"Asc",
		"Desc",
	}[d]
}

// Desc is a sort description
func (d Direction) Desc() string {
	return [...]string{
		"ascending",
		"descending",
	}[d]
}
