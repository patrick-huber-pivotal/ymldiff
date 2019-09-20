package diff

import "strconv"

// Operation denotes the type of diff operation
type Operation string

const (
	// Add denotes yml section was added
	Add Operation = "add"
	// Remove denotes a yml section was removed
	Remove Operation = "remove"
	// Modify denotes a yml section was modified
	Modify Operation = "modify"
)

// Change represents a difference between two yml documents
type Change struct {
	Operation Operation
	Path      []string
	From      interface{}
	To        interface{}
}

// IsArrayIndex returns true if the patch points to an array index
func (c *Change) IsArrayIndex() bool {
	if len(c.Path) == 0 {
		return false
	}
	element := c.Path[len(c.Path)-1]
	if _, err := strconv.Atoi(element); err == nil {
		return true
	}
	return false
}

// ChangeLog represents a difference report between two sources
type ChangeLog struct {
	Differences []Change
}
