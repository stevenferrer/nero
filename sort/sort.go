package sort

// Sort is a sorter
type Sort struct {
	Col       string
	Direction Direction
}

// Sorts is a collection of sorters
type Sorts struct {
	list []*Sort
}

// Add adds sorters to the collection
func (s *Sorts) Add(ss ...*Sort) {
	s.list = append(s.list, ss...)
}

// All returns all sorters
func (s *Sorts) All() []*Sort {
	return s.list
}
