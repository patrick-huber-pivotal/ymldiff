package formatters_test

import (
	"bytes"
	"fmt"

	"gopkg.in/yaml.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/patrick-huber-pivotal/ymldiff/diff"
	"github.com/patrick-huber-pivotal/ymldiff/formatters"
)

type boshTest struct {
}

var _ = Describe("Yaml", func() {
	var (
		tester BoshTest
	)
	BeforeEach(func() {
		tester = NewBoshTest()
	})
	Describe("Write", func() {
		It("shows items added to array", func() {

			from := []string{"one", "two", "three", "four", "seven", "ten"}
			to := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

			expected := []interface{}{
				map[interface{}]interface{}{
					"type":  "replace",
					"path":  "/4:before",
					"value": "five",
				},
				map[interface{}]interface{}{
					"type":  "replace",
					"path":  "/5:before",
					"value": "six",
				},
				map[interface{}]interface{}{
					"type":  "replace",
					"path":  "/7:before",
					"value": "eight",
				},
				map[interface{}]interface{}{
					"type":  "replace",
					"path":  "/8:before",
					"value": "nine",
				},
			}
			tester.Equal(from, to, expected)
		})

		It("shows items removed from array", func() {
			from := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
			to := []string{"one", "two", "three", "four", "seven", "ten"}

			expected := []interface{}{
				map[interface{}]interface{}{
					"type": "remove",
					"path": "/4",
				},
				map[interface{}]interface{}{
					"type": "remove",
					"path": "/4",
				},
				map[interface{}]interface{}{
					"type": "remove",
					"path": "/5",
				},
				map[interface{}]interface{}{
					"type": "remove",
					"path": "/5",
				},
			}
			tester.Equal(from, to, expected)
		})

		It("shows proper offsets from two different arrays", func() {
			from := map[string]interface{}{
				"first": []string{
					"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
				},
				"second": []string{
					"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
				},
			}
			to := map[string]interface{}{
				"first": []string{
					"one", "two", "three", "four", "seven", "ten",
				},
				"second": []string{
					"one", "two", "three", "four", "seven", "ten",
				},
			}
			expected := []interface{}{
				map[interface{}]interface{}{
					"type": "remove",
					"path": "/first/4",
				},
				map[interface{}]interface{}{
					"type": "remove",
					"path": "/first/4",
				},
				map[interface{}]interface{}{
					"type": "remove",
					"path": "/first/5",
				},
				map[interface{}]interface{}{
					"type": "remove",
					"path": "/first/5",
				},
				map[interface{}]interface{}{
					"type": "remove",
					"path": "/second/4",
				},
				map[interface{}]interface{}{
					"type": "remove",
					"path": "/second/4",
				},
				map[interface{}]interface{}{
					"type": "remove",
					"path": "/second/5",
				},
				map[interface{}]interface{}{
					"type": "remove",
					"path": "/second/5",
				},
			}
			tester.Equal(from, to, expected)
		})
	})
})

type BoshTest interface {
	Equal(from interface{}, to interface{}, expected interface{})
}

func NewBoshTest() BoshTest {
	return &boshTest{}
}

func (t *boshTest) Equal(from interface{}, to interface{}, expected interface{}) {
	changeLog, err := diff.NewChangeLogFromStructures(from, to)
	Expect(err).To(BeNil())

	formatter := formatters.NewBOSH(changeLog)
	writer := &bytes.Buffer{}
	err = formatter.Write(writer)
	Expect(err).To(BeNil())
	fmt.Println(writer.String())

	var actual interface{}
	err = yaml.Unmarshal(writer.Bytes(), &actual)
	Expect(err).To(BeNil())

	changeLog, err = diff.NewChangeLogFromStructures(actual, expected)
	Expect(err).To(BeNil())
	Expect(len(changeLog.Differences)).To(Equal(0))
}
