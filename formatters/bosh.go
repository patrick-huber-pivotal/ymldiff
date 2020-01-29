package formatters

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

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

	// sort the slice first to line up like paths
	sort.Slice(f.changeLog.Differences, func(i, j int) bool {
		first := f.changeLog.Differences[i]
		second := f.changeLog.Differences[j]
		firstPath := strings.Join(first.Path, "/")
		secondPath := strings.Join(second.Path, "/")
		return firstPath < secondPath
	})

	shift := 0
	previousPathPrefix := ""
	for _, c := range f.changeLog.Differences {
		currentPathPrefix := strings.Join(c.Path[:len(c.Path)-1], "/")
		if previousPathPrefix != currentPathPrefix {
			shift = 0
			previousPathPrefix = currentPathPrefix
		}
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
	fmt.Fprint(out, "  value:")
	bytes, err := yaml.Marshal(change.To)
	if err != nil {
		return err
	}
	lines, err := f.toLines(bytes)
	if err != nil {
		return err
	}
	for _, line := range lines {
		fmt.Fprintln(out)
		fmt.Fprintf(out, "    %s", line)
	}
	fmt.Fprintln(out)
	return nil
}

func (f *bosh) toLines(buffer []byte) (lines []string, err error) {
	r := bytes.NewReader(buffer)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
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
	fmt.Fprint(out, "  value:")
	bytes, err := yaml.Marshal(change.To)
	if err != nil {
		return err
	}
	lines, err := f.toLines(bytes)
	if err != nil {
		return err
	}
	for _, line := range lines {
		fmt.Fprintln(out)
		fmt.Fprintf(out, "    %s", line)
	}
	fmt.Fprintln(out)
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
