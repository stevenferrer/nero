package nero

// Templater is an interface that wraps the Filename and Template method
type Templater interface {
	// Filename is the filename of the generated file
	Filename() string
	// Template is template for generating the repository implementation
	Template() string
}
