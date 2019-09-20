package diff_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/patrick-huber-pivotal/ymldiff/diff"
)

var _ = Describe("Diff", func() {
	Describe("NewChangeLogFromStructures", func() {
		It("shows strings added to list", func() {
			from := []string{"dog", "cat", "bear", "lion", "shrimp"}
			to := []string{"dog", "cat", "squirrel", "bear", "lion", "shrimp", "alligator"}
			report, err := diff.NewChangeLogFromStructures(from, to)
			Expect(err).To(BeNil())
			Expect(report).ToNot(BeNil())
			Expect(len(report.Differences)).To(Equal(2))
		})
	})
	Describe("NewChangeLogFromFiles", func() {
		It("shows diff between files", func() {
			from := "../examples/list/add/from.yml"
			to := "../examples/list/add/to.yml"
			report, err := diff.NewChangeLogFromFiles(from, to)
			Expect(err).To(BeNil())
			Expect(report).ToNot(BeNil())
			Expect(len(report.Differences)).To(Equal(4))
		})
	})
})
