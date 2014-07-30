package parser_test

import (
	"testing"

	. "github.com/marciotrindade/pgsqlog/parser"
	"github.com/stretchr/testify/assert"
)

func TestTableOfSelectWithSelectCommand(t *testing.T) {
	assert := assert.New(t)

	query := `select  "messages".* from "messages"  where "messages"."id" = 112 limit 1`

	assert.Equal(TableOfSelect(query), "messages")
}

func TestTableOfSelectWithoutSelectCommand(t *testing.T) {
	assert := assert.New(t)

	query := `insert into "openings" ("contact_id", "created_at", "ip", "message_id", "updated_at") values ($1, $2, $3, $4, $5) returning "id"`

	assert.Equal(TableOfSelect(query), "")
}
