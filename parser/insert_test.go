package parser_test

import (
	. "pgsqlog/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Insert", func() {
	Describe("TableOfInsert", func() {
		var query string

		Context("when query is a delete command", func() {
			It("extracts table name of query", func() {
				query = `insert into "openings" ("contact_id", "created_at", "ip", "message_id", "updated_at") values ($1, $2, $3, $4, $5) returning "id"`
				Expect(TableOfInsert(query)).To(Equal("openings"))
			})
		})

		Context("when query is not a copy command", func() {
			It("returns empty string", func() {
				query = "select * from logs where duration > 1"
				Expect(TableOfInsert(query)).To(BeEmpty())
			})
		})
	})
})
