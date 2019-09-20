package formatters

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v2"

	"github.com/patrick-huber-pivotal/ymldiff/diff"
)

type ymlFormatter struct {
	changeLog *diff.ChangeLog
}

// NewYAML returns a yaml change log formatter
func NewYAML(changeLog *diff.ChangeLog) Formatter {
	return &ymlFormatter{
		changeLog: changeLog,
	}
}

func (f *ymlFormatter) Write(out io.Writer) error {
	if f.changeLog == nil {
		return fmt.Errorf("changelog is nil")
	}
	bytes, err := yaml.Marshal(f.changeLog)
	if err != nil {
		return err
	}
	_, err = out.Write(bytes)
	return err
}
