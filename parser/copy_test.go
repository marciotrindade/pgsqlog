package parser_test

import (
	. "pgsqlog/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Copy", func() {
	Describe("TableOfCopy", func() {
		Context("with schema name", func() {
			query := "copy account_29512.import_19 (email,nome) from stdin (format 'csv')"

			It("extracts table name of query", func() {
				Expect(TableOfCopy(query)).To(Equal("account_29512.import_19"))
			})
		})

		Context("without schema name", func() {
			query := "copy import_19 (email,nome) from stdin (format 'csv')"

			It("extracts table name of query", func() {
				Expect(TableOfCopy(query)).To(Equal("import_19"))
			})
		})

		Context("when query is not a copy command", func() {
			query := "select * from logs where duration > 1"

			It("returns empty string", func() {
				Expect(TableOfCopy(query)).To(BeEmpty())
			})
		})
	})
})
