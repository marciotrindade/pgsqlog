package parser_test

import (
	"testing"

	. "github.com/marciotrindade/pgsqlog/parser"
	"github.com/stretchr/testify/assert"
)

func TestTableOfUpdateWithUpdateCommand(t *testing.T) {
	assert := assert.New(t)

	query := `update "messages" set "clicks_count" = coalesce("clicks_count", 0) + 1 where "messages"."id" = 3`

	assert.Equal(TableOfUpdate(query), "messages")
}

func TestTableOfUpdateWithoutUpdateCommand(t *testing.T) {
	assert := assert.New(t)

	query := "select * from logs where duration > 1"

	assert.Equal(TableOfUpdate(query), "")
}
