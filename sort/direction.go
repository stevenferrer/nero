package sort

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
	// Asc is an ascending sort direction
	Asc Direction = iota
	// Desc is a descending sort direction
	Desc
)
