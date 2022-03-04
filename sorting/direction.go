package sorting

// Direction is a sort direction
type Direction int

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

const (
	// Asc ascending sort direction
	Asc Direction = iota
	// Desc descending sort direction
	Desc
)
