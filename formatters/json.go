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
	changes := []diff.Change{}
	for _, change := range changeLog.Differences {
		newChange := diff.Change{
			From:      clean(change.From),
			To:        clean(change.To),
			Operation: change.Operation,
			Path:      change.Path,
		}
		changes = append(changes, newChange)
	}
	return &jsonFormatter{
		changeLog: &diff.ChangeLog{
			Differences: changes,
		},
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

func clean(data interface{}) interface{} {
	switch t := data.(type) {
	case (map[interface{}]interface{}):
		return cleanInterfaceMap(t)
	}
	return data
}

func cleanInterfaceMap(data map[interface{}]interface{}) map[string]interface{} {
	m := map[string]interface{}{}
	for k, v := range data {
		s := k.(string)
		m[s] = clean(v)
	}
	return m
}
