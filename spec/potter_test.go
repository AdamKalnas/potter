package spec_test

import (
	"github.com/adamkalnas/potter/bookstore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Potter", func() {

	It("Should charge $0 for 0 books", func() {
		books := []string{}
		Expect(bookstore.Cost(books)).To(Equal(0))
	})

})
