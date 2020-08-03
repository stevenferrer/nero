package sort

// Direction is a sort direction
type Direction int

func (d Direction) String() string {
	switch d {
	case Asc:
		return "Asc"
	case Desc:
		return "Desc"
	}

	return "Invalid"
}

const (
	// Asc is an ascending sort direction
	Asc Direction = iota
	// Desc is a descending sort direction
	Desc
)
