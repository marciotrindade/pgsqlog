package parser_test

import (
	. "pgsqlog/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Alter", func() {

	Describe("TableOfAlter", func() {
		Context("when query is a alter command", func() {
			It("extracts table name of query", func() {
				query := "alter table logs rename column sql to query;"
				Expect(TableOfAlter(query)).To(Equal("logs"))
			})
		})

		Context("when query isn't a alter command", func() {
			It("returns empty string", func() {
				query := "select * from logs where duration > 1"
				Expect(TableOfCopy(query)).To(BeEmpty())
			})
		})
	})

})
