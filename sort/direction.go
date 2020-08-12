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

func (d Direction) Description() string {
	switch d {
	case Asc:
		return "Asc is the ascending sort direction"
	case Desc:
		return "Desc is the descending sort direction"
	}

	return ""
}

const (
	// Asc is an ascending sort direction
	Asc Direction = iota
	// Desc is a descending sort direction
	Desc
)
