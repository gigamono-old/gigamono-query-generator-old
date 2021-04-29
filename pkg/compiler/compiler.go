package compiler

// SQLCompiler compiles document to JavaScript code.
type SQLCompiler struct{}

// Compile compiles document to JavaScript code.
func (compiler *SQLCompiler) Compile(documentString string, opts ...interface{}) (string, error) {
	return "", nil
}
