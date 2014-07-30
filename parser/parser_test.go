package parser_test

import (
	"testing"

	. "github.com/marciotrindade/pgsqlog/parser"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	logLine := LogLine{Line: `2014-06-14 15:30:39 brt [9405]: [40764-1] user=emailmarketing,db=emailmarketing log:  duration: 139.527 ms  statement: select  "messages".* from "messages"  where "messages"."id" = 112 limit 1`}
	logLine.Parse()

	assert.Equal(logLine.Action, "select")
	assert.Equal(logLine.CreatedAt, "2014-06-14 15:30:39")
	assert.Equal(logLine.Database, "emailmarketing")
	assert.Equal(logLine.Duration, "139.527")
	assert.Equal(logLine.Query, `select  "messages".* from "messages"  where "messages"."id" = 112 limit 1`)
	assert.Equal(logLine.Table, "messages")
	assert.Equal(logLine.Username, "emailmarketing")
}

func TestSqlSelect(t *testing.T) {
	assert := assert.New(t)

	logLine := LogLine{Query: `select  "messages".* from "messages"  where "messages"."id" = 112 limit 1`}
	logLine.Sql()

	assert.Equal(logLine.Action, "select")
	assert.Equal(logLine.Table, "messages")
}

func TestSqlUpdate(t *testing.T) {
	assert := assert.New(t)

	logLine := LogLine{Query: `update "messages" set "clicks_count" = coalesce("clicks_count", 0) + 1 where "messages"."id" = 3`}
	logLine.Sql()

	assert.Equal(logLine.Action, "update")
	assert.Equal(logLine.Table, "messages")
}

func TestSqlInsert(t *testing.T) {
	assert := assert.New(t)

	logLine := LogLine{Query: `insert into "openings" ("contact_id", "created_at", "ip", "message_id", "updated_at") values ($1, $2, $3, $4, $5) returning "id"`}
	logLine.Sql()

	assert.Equal(logLine.Action, "insert")
	assert.Equal(logLine.Table, "openings")
}

func TestSqlCopy(t *testing.T) {
	assert := assert.New(t)

	logLine := LogLine{Query: `copy account_29656.import_6 (email) from stdin (format 'csv')`}
	logLine.Sql()

	assert.Equal(logLine.Action, "copy")
	assert.Equal(logLine.Table, "account_29656.import_6")
}

func TestSqlDelete(t *testing.T) {
	assert := assert.New(t)

	logLine := LogLine{Query: `delete from "contact_imports" where "contact_imports"."id" = $1`}
	logLine.Sql()

	assert.Equal(logLine.Action, "delete")
	assert.Equal(logLine.Table, "contact_imports")
}

func TestSqlCreate(t *testing.T) {
	assert := assert.New(t)

	logLine := LogLine{Query: `create table account_29656.import_6 (id serial primary key, email text)`}
	logLine.Sql()

	assert.Equal(logLine.Action, "create table")
	assert.Equal(logLine.Table, "account_29656.import_6")
}

func TestSqlDrop(t *testing.T) {
	assert := assert.New(t)

	logLine := LogLine{Query: `drop table account_29656.import_6`}
	logLine.Sql()

	assert.Equal(logLine.Action, "drop table")
	assert.Equal(logLine.Table, "account_29656.import_6")
}

func TestSqlAlter(t *testing.T) {
	assert := assert.New(t)

	logLine := LogLine{Query: `alter table logs rename column sql to query`}
	logLine.Sql()

	assert.Equal(logLine.Action, "alter")
	assert.Equal(logLine.Table, "logs")
}
