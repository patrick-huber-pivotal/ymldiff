package formatters

import "io"

// Formatter exposes a write method for writing the change log to a writer stream
type Formatter interface {
	Write(out io.Writer) error
}
