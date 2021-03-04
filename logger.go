package nero

// Logger is an interface that wraps the Printf method
type Logger interface {
	Printf(string, ...interface{})
}
