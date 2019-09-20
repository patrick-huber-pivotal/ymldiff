package formatters

import (
	"fmt"
	"io"
	"strconv"

	"github.com/patrick-huber-pivotal/ymldiff/diff"
	"gopkg.in/yaml.v2"
)

type bosh struct {
	changeLog *diff.ChangeLog
}

// NewBOSH returns a json change log formatter
func NewBOSH(changeLog *diff.ChangeLog) Formatter {
	return &bosh{
		changeLog: changeLog,
	}
}

func (f *bosh) Write(out io.Writer) error {
	shift := 0
	for _, c := range f.changeLog.Differences {
		var err error
		shift, err = f.writeChange(&c, out, shift)
		if err != nil {
			return err
		}
		fmt.Fprintln(out)
	}
	return nil
}

func (f *bosh) writeChange(change *diff.Change, out io.Writer, shift int) (int, error) {
	switch change.Operation {
	case diff.Add:
		return shift, f.writeAdd(change, out)
	case diff.Remove:
		err := f.writeRemove(change, out, shift)
		if change.IsArrayIndex() {
			shift++
		}
		return shift, err
	case diff.Modify:
		return shift, f.writeModify(change, out)
	}
	return 0, fmt.Errorf("unrecognized change operation %s", change.Operation)
}

func (f *bosh) writeAdd(change *diff.Change, out io.Writer) error {
	fmt.Fprintf(out, "- type: replace")
	fmt.Fprintln(out)
	fmt.Fprintf(out, "  path: %s", f.createAddPath(change.Path))
	fmt.Fprintln(out)
	bytes, err := yaml.Marshal(change.To)
	if err != nil {
		return err
	}
	fmt.Fprintln(out, "  value:")
	fmt.Fprintf(out, "    %s", string(bytes))
	return nil
}

func (f *bosh) writeRemove(change *diff.Change, out io.Writer, shift int) error {
	fmt.Fprintf(out, "- type: remove")
	fmt.Fprintln(out)
	fmt.Fprintf(out, "  path: %s", f.createRemovePath(change.Path, shift))
	fmt.Fprintln(out)
	return nil
}

func (f *bosh) writeModify(change *diff.Change, out io.Writer) error {
	fmt.Fprintf(out, "- type: replace")
	fmt.Fprintln(out)
	fmt.Fprintf(out, "  path: %s", f.createAddPath(change.Path))
	fmt.Fprintln(out)
	bytes, err := yaml.Marshal(change.To)
	if err != nil {
		return err
	}
	fmt.Fprintln(out, "  value:")
	fmt.Fprintf(out, "    %s", string(bytes))
	return nil
}

func (f *bosh) createRemovePath(path []string, shift int) string {
	response := ""
	for i, element := range path {
		response += "/"
		if val, err := strconv.Atoi(element); err == nil {
			if len(path)-1 == i {
				val = val - shift
			}
			response += strconv.Itoa(val)
			continue
		}
		response += element
	}
	return response
}

func (f *bosh) createAddPath(path []string) string {
	response := ""
	for i, element := range path {
		response += "/"
		if _, err := strconv.Atoi(element); err == nil {
			response += element + ":before"
			continue
		}
		response += element
		if len(path)-1 == i {
			response += "?"
		}
	}
	return response
}
