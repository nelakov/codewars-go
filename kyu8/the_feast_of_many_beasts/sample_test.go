package the_feast_of_many_beasts

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "The Feast of Many Beasts Suite")
}

var _ = Describe("Sample Test Cases", func() {
	It("should return the correct value", func() {
		Expect(Feast("great blue heron", "garlic naan")).To(BeTrue(), "Testing 'great blue heron' vs 'garlic naan'")
		Expect(Feast("chickadee", "chocolate cake")).To(BeTrue(), "Testing 'chickadee' vs 'chocolate cake'")
		Expect(Feast("brown bear", "bear claw")).To(BeFalse(), "Testing 'brown bear' vs 'bear claw'")
	})
})
