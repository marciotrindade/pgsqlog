package parser_test

import (
	"testing"

	. "github.com/marciotrindade/pgsqlog/parser"
	"github.com/stretchr/testify/assert"
)

func TestTableOfInsertWithInsertCommand(t *testing.T) {
	assert := assert.New(t)

	query := `insert into "openings" ("contact_id", "created_at", "ip", "message_id", "updated_at") values ($1, $2, $3, $4, $5) returning "id"`

	assert.Equal(TableOfInsert(query), "openings")
}

func TestTableOfInsertWithoutInsertCommand(t *testing.T) {
	assert := assert.New(t)

	query := "select * from logs where duration > 1"

	assert.Equal(TableOfInsert(query), "")
}
