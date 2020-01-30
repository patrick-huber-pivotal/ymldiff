package formatters

import (
	"encoding/json"
	"io"

	"github.com/patrick-huber-pivotal/ymldiff/diff"
)

type jsonFormatter struct {
	changeLog *diff.ChangeLog
}

// NewJSON returns a json change log formatter
func NewJSON(changeLog *diff.ChangeLog) Formatter {
	return &jsonFormatter{
		changeLog: changeLog,
	}
}

func (f *jsonFormatter) Write(out io.Writer) error {
	data, err := json.Marshal(f.changeLog)
	if err != nil {
		return err
	}
	_, err = out.Write(data)
	return err
}
