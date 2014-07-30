package parser_test

import (
	"testing"

	. "github.com/marciotrindade/pgsqlog/parser"
	"github.com/stretchr/testify/assert"
)

func TestTableOfCreateWithCreateTableCommandAndSchema(t *testing.T) {
	assert := assert.New(t)

	query := "create table account_5943.import_5 (id serial primary key, email text, text, text, text, text)"
	action, table := TableOfCreate(query)

	assert.Equal(action, "create table")
	assert.Equal(table, "account_5943.import_5")
}

func TestTableOfCreateWithCreateTableCommand(t *testing.T) {
	assert := assert.New(t)

	query := `create table "templates" ("id" serial primary key, "name" character varying(255) not null, "account_id" integer, "created_at" timestamp not null, "updated_at" timestamp not null)`
	action, table := TableOfCreate(query)

	assert.Equal(action, "create table")
	assert.Equal(table, "templates")
}

func TestTableOfCreateWithCreateIndexCommand(t *testing.T) {
	assert := assert.New(t)

	query := `create  index "index_openings_on_contact_id" on "openings" ("contact_id")`
	action, table := TableOfCreate(query)

	assert.Equal(action, "create index")
	assert.Equal(table, "index_openings_on_contact_id")
}

func TestTableOfCreateWithCreateUniqueIndexCommand(t *testing.T) {
	assert := assert.New(t)

	query := `create unique index "index_openings_on_contact_id" on "openings" ("contact_id")`
	action, table := TableOfCreate(query)

	assert.Equal(action, "create index")
	assert.Equal(table, "index_openings_on_contact_id")
}

func TestTableOfCreateWithoutCreateCommand(t *testing.T) {
	assert := assert.New(t)

	query := "select * from logs where duration > 1"
	action, table := TableOfCreate(query)

	assert.Equal(action, "")
	assert.Equal(table, "")
}
