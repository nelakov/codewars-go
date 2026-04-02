package contamination_1_string

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contamination #1 -String- Suite")
}

func dotest(text, char, expected string) {
	Expect(Contamination(text, char)).To(Equal(expected), "With text = \"%s\", char = \"%s\"", text, char)
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		dotest("abc", "z", "zzz")
		dotest("", "z", "")
		dotest("abc", "", "")
		dotest("_3ebzgh4", "&", "&&&&&&&&")
		dotest("//case", " ", "      ")
	})
})
