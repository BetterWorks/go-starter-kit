package trace

// ContextKey wraps scalar string for use in context (to avoid naming collisions)
type ContextKey string

// String implements the Stringer interface for print statements
func (c ContextKey) String() string {
	return string(c)
}
