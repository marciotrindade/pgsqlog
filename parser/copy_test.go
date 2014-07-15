package parser_test

import (
	. "pgsqlog/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Copy", func() {
	Describe("TableOfCopy", func() {
		Context("with schema name", func() {
			It("extracts table name of query", func() {
				query := "copy account_29512.import_19 (email,nome) from stdin (format 'csv')"
				Expect(TableOfCopy(query)).To(Equal("account_29512.import_19"))
			})
		})

		Context("without schema name", func() {
			It("extracts table name of query", func() {
				query := "copy import_19 (email,nome) from stdin (format 'csv')"
				Expect(TableOfCopy(query)).To(Equal("import_19"))
			})
		})

		Context("when query is not a copy command", func() {
			It("returns empty string", func() {
				query := "select * from logs where duration > 1"
				Expect(TableOfCopy(query)).To(BeEmpty())
			})
		})
	})
})
