package generator

// SQLQueryGenerator generates SQL query from workflow specification.
type SQLQueryGenerator struct{}

// Generate generates SQL query from workflow specification.
func (compiler *SQLQueryGenerator) Generate(workflowString string, opts ...interface{}) (string, error) {
	return "", nil
}
