package storage

// Storage represents the data persistence boundary.
// In accordance with AGENTS.md, no database schema or business queries
// are implemented in this initial skeleton phase.
type Storage struct{}

// New initializes an empty Storage placeholder.
func New() *Storage {
	return &Storage{}
}
