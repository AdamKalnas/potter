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

	It("Should charge $16 for 2 copies of the same book", func() {
		books := make(map[string]int)
		books["Goblet of Fire"] = 2

		Expect(bookstore.Cost(books)).To(Equal(1600))
	})

	It("Should charge $15.20 for 2 different books", func() {
		books := make(map[string]int)
		books["Prisoner of Azkaban"] = 1
		books["Goblet of Fire"] = 1

		Expect(bookstore.Cost(books)).To(Equal(1520))
	})
})
