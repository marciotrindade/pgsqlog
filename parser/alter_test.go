package parser_test

import (
	. "pgsqlog/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Alter", func() {

	Describe("TableOfAlter", func() {
		Context("when query is a alter command", func() {
			query := "alter table logs rename column sql to query;"

			It("extracts table name of query", func() {
				Expect(TableOfAlter(query)).To(Equal("logs"))
			})
		})

		Context("when query isn't a alter command", func() {
			query := "select * from logs where duration > 1"

			It("returns empty string", func() {
				Expect(TableOfCopy(query)).To(BeEmpty())
			})
		})
	})

})
