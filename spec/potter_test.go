package spec_test

import (
	"github.com/adamkalnas/potter/bookstore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Potter", func() {

	It("Should charge $0 for 0 books", func() {
		books := make(map[string]int)
		books["Goblet of Fire"] = 0

		Expect(bookstore.Cost(books)).To(Equal(0))
	})

	It("Should charge $8 for 1 book", func() {
		books := make(map[string]int)
		books["Goblet of Fire"] = 1

		Expect(bookstore.Cost(books)).To(Equal(800))
	})

})
