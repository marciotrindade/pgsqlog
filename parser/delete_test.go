package parser_test

import (
	"testing"

	. "github.com/marciotrindade/pgsqlog/parser"
	"github.com/stretchr/testify/assert"
)

func TestTableOfDeleteWithDeleteCommand(t *testing.T) {
	assert := assert.New(t)
	query := `delete from "contact_imports" where "contact_imports"."id" = $1`
	assert.Equal(TableOfDelete(query), "contact_imports")
}

func TestTableOfDeleteWithoutDeleteCommand(t *testing.T) {
	assert := assert.New(t)
	query := "select * from logs where duration > 1"
	assert.Equal(TableOfDelete(query), "")
}
