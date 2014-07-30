package parser_test

import (
	"testing"

	. "github.com/marciotrindade/pgsqlog/parser"
	"github.com/stretchr/testify/assert"
)

func TestTableOfDropWithTable(t *testing.T) {
	assert := assert.New(t)

	query := "drop table import_5;"
	action, table := TableOfDrop(query)

	assert.Equal(action, "drop table")
	assert.Equal(table, "import_5")
}

func TestTableOfDropWithSchema(t *testing.T) {
	assert := assert.New(t)

	query := "drop schema account_29633 cascade;"
	action, table := TableOfDrop(query)

	assert.Equal(action, "drop schema")
	assert.Equal(table, "account_29633")
}

func TestTableOfDropWithoutDropCommand(t *testing.T) {
	assert := assert.New(t)

	query := "select * from logs where duration > 1"
	action, table := TableOfDrop(query)

	assert.Equal(action, "")
	assert.Equal(table, "")
}
