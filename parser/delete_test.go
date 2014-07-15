package parser_test

import (
	. "pgsqlog/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete", func() {
	Describe("TableOfDelete", func() {
		var query string

		Context("when query is a delete command", func() {
			It("extracts table name of query", func() {
				query = `delete from "contact_imports" where "contact_imports"."id" = $1`
				Expect(TableOfDelete(query)).To(Equal("contact_imports"))
			})
		})

		Context("when query is not a copy command", func() {
			It("returns empty string", func() {
				query = "select * from logs where duration > 1"
				Expect(TableOfDelete(query)).To(BeEmpty())
			})
		})
	})
})
