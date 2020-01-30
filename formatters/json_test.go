package formatters_test

import (
	"bufio"
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/patrick-huber-pivotal/ymldiff/diff"
	"github.com/patrick-huber-pivotal/ymldiff/formatters"
)

var _ = Describe("Json", func() {
	Describe("Write", func() {
		It("can deserialze map interface", func() {

			var b bytes.Buffer
			buffer := bufio.NewWriter(&b)
			changeLog := &diff.ChangeLog{
				Differences: []diff.Change{
					diff.Change{
						Operation: diff.Add,
						Path:      []string{"test"},
						From: map[interface{}]interface{}{
							"test": "hi",
						},
						To: map[interface{}]interface{}{
							"hello": "world",
						},
					},
				},
			}
			f := formatters.NewJSON(changeLog)
			Expect(f.Write(buffer)).To(BeNil())
		})
	})
})
