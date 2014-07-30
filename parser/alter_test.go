package parser_test

import (
	"testing"

	. "github.com/marciotrindade/pgsqlog/parser"
	"github.com/stretchr/testify/assert"
)

func TestTableOfAlterWithAlterCommand(t *testing.T) {
	assert := assert.New(t)
	query := "alter table logs rename column sql to query;"
	assert.Equal(TableOfAlter(query), "logs")
}

func TestTableOfAlterWithoutAlterCommand(t *testing.T) {
	assert := assert.New(t)
	query := "select * from logs where duration > 1"
	assert.Equal(TableOfAlter(query), "")
}
