package parser_test

import (
	. "pgsqlog/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create", func() {
	Describe("when create a table", func() {
		Context("without schema name", func() {
			Context("with schema name", func() {
				query := "create table account_5943.import_5 (id serial primary key, email text, text, text, text, text)"

				It("extracts table name of query", func() {
					action, table := TableOfCreate(query)
					Expect(action).To(Equal("create table"))
					Expect(table).To(Equal("account_5943.import_5"))
				})
			})

			Context("without schema name", func() {
				query := "create table \"templates\" (\"id\" serial primary key, \"name\" character varying(255) not null, \"account_id\" integer, \"created_at\" timestamp not null, \"updated_at\" timestamp not null)"

				It("extracts table name of query", func() {
					action, table := TableOfCreate(query)
					Expect(action).To(Equal("create table"))
					Expect(table).To(Equal("templates"))
				})
			})
		})

		Context("when create a index", func() {
			query := "create  index \"index_openings_on_contact_id\" on \"openings\" (\"contact_id\")"

			It("extracts table name of query", func() {
				action, table := TableOfCreate(query)
				Expect(action).To(Equal("create index"))
				Expect(table).To(Equal("index_openings_on_contact_id"))
			})
		})

		Context("when query is not a create command", func() {
			query := "select * from logs where duration > 1"

			It("returns empty string", func() {
				action, table := TableOfCreate(query)
				Expect(action).To(BeEmpty())
				Expect(table).To(BeEmpty())
			})
		})
	})
})
