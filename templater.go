package nero

// Templater is the contract for a template that implements the repository interface
type Templater interface {
	// Filename is the fileame of the generated file
	Filename() string
	// Content is the actual template to be used for
	// generating the repository implemntation
	Content() string
}
