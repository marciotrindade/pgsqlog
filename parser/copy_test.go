package parser_test

import (
	"testing"

	. "github.com/marciotrindade/pgsqlog/parser"
	"github.com/stretchr/testify/assert"
)

func TestTableOfCopyWithCopyCommandAndSchema(t *testing.T) {
	assert := assert.New(t)

	query := "copy account_29512.import_19 (email,nome) from stdin (format 'csv')"
	assert.Equal(TableOfCopy(query), "account_29512.import_19")
}

func TestTableOfCopyWithCopyCommand(t *testing.T) {
	assert := assert.New(t)

	query := "copy import_19 (email,nome) from stdin (format 'csv')"
	assert.Equal(TableOfCopy(query), "import_19")
}

func TestTableOfCopyWithoutCopyCommand(t *testing.T) {
	assert := assert.New(t)

	query := "select * from logs where duration > 1"
	assert.Equal(TableOfCopy(query), "")
}
