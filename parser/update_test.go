package parser_test

import (
	. "pgsqlog/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Update", func() {
	Describe("TableOfUpdate", func() {
		var query string

		Context("when query is a delete command", func() {
			It("extracts table name of query", func() {
				query = `update "messages" set "clicks_count" = coalesce("clicks_count", 0) + 1 where "messages"."id" = 3`
				Expect(TableOfUpdate(query)).To(Equal("messages"))
			})
		})

		Context("when query is not a copy command", func() {
			It("returns empty string", func() {
				query = `select  "messages".* from "messages"  where "messages"."id" = 112 limit 1`
				Expect(TableOfUpdate(query)).To(BeEmpty())
			})
		})
	})
})
