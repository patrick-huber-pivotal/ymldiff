package formatters_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/patrick-huber-pivotal/ymldiff/diff"
	"github.com/patrick-huber-pivotal/ymldiff/formatters"
)

var _ = Describe("Yaml", func() {
	Describe("Write", func() {
		It("shows items added to array", func() {
			changeLog := &diff.ChangeLog{
				Differences: []diff.Change{
					diff.Change{
						Operation: diff.Add,
						Path:      []string{"0"},
						From:      nil,
						To:        "three",
					},
				},
			}
			formatter := formatters.NewYAML(changeLog)
			writer := &bytes.Buffer{}
			err := formatter.Write(writer)
			Expect(err).To(BeNil())
		})
	})
})
