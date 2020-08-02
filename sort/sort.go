package sort

// Sort is a sorter
type Sort struct {
	Field     string
	Direction Direction
}

// Sorts is a collection of sorters
type Sorts struct {
	list []*Sort
}

// Add adds sorters to the collection
func (srt *Sorts) Add(ss ...*Sort) {
	srt.list = append(srt.list, ss...)
}

// All returns all sorters
func (srt *Sorts) All() []*Sort {
	return srt.list
}
