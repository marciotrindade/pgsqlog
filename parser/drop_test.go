package parser_test

import (
	. "pgsqlog/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Drop", func() {
	var (
		query  string
		action string
		table  string
	)

	Describe("#TableOfDrop", func() {

		Context("when drop a table", func() {
			BeforeEach(func() {
				query = `drop table import_5;`
				action, table = TableOfDrop(query)
			})

			It("extracts action name of query", func() {
				Expect(action).To(Equal("drop table"))
			})

			It("extracts table name of query", func() {
				Expect(table).To(Equal("import_5"))
			})
		})

		Context("when drop a schema", func() {
			BeforeEach(func() {
				query = `drop schema account_29633 cascade;`
				action, table = TableOfDrop(query)
			})

			It("extracts action name of query", func() {
				Expect(action).To(Equal("drop schema"))
			})

			It("extracts table name of query", func() {
				Expect(table).To(Equal("account_29633"))
			})
		})

		Context("when query is not a create command", func() {
			BeforeEach(func() {
				query = "select * from logs where duration > 1"
				action, table = TableOfDrop(query)
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
