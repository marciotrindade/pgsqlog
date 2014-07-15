package parser_test

import (
	. "pgsqlog/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Select", func() {
	Describe("TableOfSelect", func() {
		var query string

		Context("when query is a delete command", func() {
			It("extracts table name of query", func() {
				query = `select  "messages".* from "messages"  where "messages"."id" = 112 limit 1`
				Expect(TableOfSelect(query)).To(Equal("messages"))
			})
		})

		Context("when query is not a copy command", func() {
			It("returns empty string", func() {
				query = `insert into "openings" ("contact_id", "created_at", "ip", "message_id", "updated_at") values ($1, $2, $3, $4, $5) returning "id"`
				Expect(TableOfSelect(query)).To(BeEmpty())
			})
		})
	})
})
