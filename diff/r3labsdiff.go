package diff

import (
	"io/ioutil"

	structdiff "github.com/r3labs/diff"
	"gopkg.in/yaml.v2"
)

// NewChangeLogFromFiles returns a report from the given paths
func NewChangeLogFromFiles(from, to string) (*ChangeLog, error) {
	fromContent, err := ioutil.ReadFile(from)
	if err != nil {
		return nil, err
	}

	toContent, err := ioutil.ReadFile(to)
	if err != nil {
		return nil, err
	}

	return NewChangeLogFromContent(fromContent, toContent)
}

// NewChangeLogFromContent returns a report from the content
// the conent will be unmarshaled into structures
func NewChangeLogFromContent(from, to []byte) (*ChangeLog, error) {
	var fromInterface interface{}
	var toInterface interface{}

	err := yaml.Unmarshal(from, &fromInterface)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(to, &toInterface)
	if err != nil {
		return nil, err
	}

	return NewChangeLogFromStructures(fromInterface, toInterface)
}

// NewChangeLogFromStructures returns a report from the given structures
func NewChangeLogFromStructures(from, to interface{}) (*ChangeLog, error) {
	if from == nil && to == nil {
		return &ChangeLog{
			Differences: []Change{},
		}, nil
	}
	changeLog, err := structdiff.Diff(from, to)
	if err != nil {
		return nil, err
	}

	return mapStructChangeLogToChangeLog(changeLog), nil
}

func mapStructChangeLogToChangeLog(structChangeLog structdiff.Changelog) *ChangeLog {
	changeLog := &ChangeLog{
		Differences: []Change{},
	}
	for _, c := range structChangeLog {
		change := mapStructDiffToDiff(&c)
		changeLog.Differences = append(changeLog.Differences, *change)
	}
	return changeLog
}

func mapStructDiffToDiff(structChange *structdiff.Change) *Change {
	change := &Change{
		Path: structChange.Path,
		From: structChange.From,
		To:   structChange.To,
	}

	switch structChange.Type {
	case structdiff.CREATE:
		change.Operation = Add
		break
	case structdiff.DELETE:
		change.Operation = Remove
		break
	case structdiff.UPDATE:
		change.Operation = Modify
		break
	}
	return change
}
