package nero

// Templater is an interface for implementing a repository
type Templater interface {
	// Filename is the fileame of the generated file
	Filename() string
	// Content is the actual template to be used for
	// generating the repository implemntation
	Content() string
}
