package sort

// Sort is a sort parameter
type Sort struct {
	Field     string
	Direction Direction
}

// SortFunc is a sort list decorator
type SortFunc func([]Sort) []Sort
