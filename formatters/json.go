package formatters

import (
	"io"

	"github.com/patrick-huber-pivotal/ymldiff/diff"
)

type json struct {
}

// NewJSON returns a json change log formatter
func NewJSON(changeLog *diff.ChangeLog) Formatter {
	return &json{}
}

func (f *json) Write(out io.Writer) error {
	return nil
}
