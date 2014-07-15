package parser_test

import (
	. "pgsqlog/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create", func() {
	var (
		query  string
		action string
		table  string
	)

	Describe("when create a table", func() {
		Context("without schema name", func() {
			Context("with schema name", func() {
				BeforeEach(func() {
					query = "create table account_5943.import_5 (id serial primary key, email text, text, text, text, text)"
					action, table = TableOfCreate(query)
				})

				It("extracts action name of query", func() {
					Expect(action).To(Equal("create table"))
				})

				It("extracts table name of query", func() {
					Expect(table).To(Equal("account_5943.import_5"))
				})
			})

			Context("without schema name", func() {
				BeforeEach(func() {
					query = `create table "templates" ("id" serial primary key, "name" character varying(255) not null, "account_id" integer, "created_at" timestamp not null, "updated_at" timestamp not null)`
					action, table = TableOfCreate(query)
				})

				It("extracts action name of query", func() {
					Expect(action).To(Equal("create table"))
				})

				It("extracts table name of query", func() {
					Expect(table).To(Equal("templates"))
				})
			})
		})

		Context("when create a index", func() {
			BeforeEach(func() {
				query = `create  index "index_openings_on_contact_id" on "openings" ("contact_id")`
				action, table = TableOfCreate(query)
			})

			It("extracts action name of query", func() {
				Expect(action).To(Equal("create index"))
			})

			It("extracts table name of query", func() {
				Expect(table).To(Equal("index_openings_on_contact_id"))
			})
		})

		Context("when query is not a create command", func() {
			BeforeEach(func() {
				query = "select * from logs where duration > 1"
				action, table = TableOfCreate(query)
			})

			It("returns empty action", func() {
				Expect(action).To(BeEmpty())
			})

			It("returns empty table", func() {
				Expect(table).To(BeEmpty())
			})
		})
	})
})
