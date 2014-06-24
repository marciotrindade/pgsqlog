package main

import (
	"database/sql"
	"github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

const (
	postgresConnection = "user=emailmarketing dbname=psqlog sslmode=disable"
	runWithRoutines    = false
	gophers_count      = 20
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("You need to set the filename\nexample: psqlog main.log")
	}

	lines := redFile(os.Args[1])

	var waitGroup sync.WaitGroup

	if runWithRoutines {
		waitGroup.Add(gophers_count)

		lines_count := len(lines)
		partial_count := lines_count / gophers_count

		for i := 0; i < gophers_count; i++ {
			var finish int
			start := i * partial_count
			if i == (gophers_count - 1) {
				finish = lines_count
			} else {
				finish = ((i + 1) * partial_count)
			}

			go gopher(i, lines[start:finish], &waitGroup)
		}
	} else {
		waitGroup.Add(1)

		gopher(1, lines, &waitGroup)
	}
	waitGroup.Wait()
}

func redFile(fileName string) []string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error:", err)
	}

	return strings.Split(string(content), "\n")
}

func gopher(i int, lines []string, waitGroup *sync.WaitGroup) {
	log.Println("Starting:", i)
	db := getConnection()
	defer db.Close()

	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := txn.Prepare(pq.CopyIn("logs", "username", "database", "duration", "action", "table_name", "sql", "created_at"))
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		parseLine(strings.ToLower(line), stmt)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
	log.Println("Finishing:", i)
	waitGroup.Done()
}

func parseLine(line string, stmt *sql.Stmt) {
	re := regexp.MustCompile(`([0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}) brt \[[0-9]*\]: \[(.*?)\] user=(.*?),db=(.*?) log:  duration: (.*?) ms  (.*?): (.*)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) > 0 {
		username := matches[3]
		database := matches[4]
		duration := matches[5]
		sql := strings.Trim(matches[7], " ")
		action, table := parseSql(sql)
		created_at := matches[1]

		if action != "" && table != "" {
			_, err := stmt.Exec(username, database, duration, action, table, sql, created_at)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func parseSql(sql string) (string, string) {
	var action string
	var table string

	re := regexp.MustCompile(`([a-z]*)(.*)`)
	matches := re.FindStringSubmatch(sql)
	switch matches[1] {
	case "select":
		action, table = parseSelect(sql)
	case "update":
		action, table = parseUpdate(sql)
	case "insert":
		action, table = parseInsert(sql)
	case "copy":
		action, table = parseCopy(sql)
	case "delete":
		action, table = parseDelete(sql)
	case "create":
		action, table = parseCreate(sql)
	case "drop":
		action, table = parseDrop(sql)
	case "alter":
		action, table = parseAlter(sql)
	}
	return action, table
}

func parseSelect(sql string) (string, string) {
	var table string
	action := "select"

	re := regexp.MustCompile(`select(.*?)from "?([a-z0-9_\.]*)"?(.*)`)
	matches := re.FindStringSubmatch(sql)

	if len(matches) > 0 {
		table = matches[2]
	}
	return action, table
}

func parseUpdate(sql string) (string, string) {
	var table string
	action := "update"

	re := regexp.MustCompile(`update "?([a-z0-9_\.]*)"?(.*)`)
	matches := re.FindStringSubmatch(sql)

	if len(matches) > 0 {
		table = matches[1]
	}
	return action, table
}

func parseInsert(sql string) (string, string) {
	var table string
	action := "insert"

	re := regexp.MustCompile(`insert into "?([a-z0-9_\.]*)"?(.*)`)
	matches := re.FindStringSubmatch(sql)

	if len(matches) > 0 {
		table = matches[1]
	}
	return action, table
}

func parseCopy(sql string) (string, string) {
	var table string
	action := "copy"

	re := regexp.MustCompile(`copy "?([a-z0-9_\.]*)"?(.*)`)
	matches := re.FindStringSubmatch(sql)

	if len(matches) > 0 {
		table = matches[1]
	}
	return action, table
}

func parseDelete(sql string) (string, string) {
	var table string
	action := "delete"

	re := regexp.MustCompile(`delete from "?([[a-z0-9_\.]*)"?(.*)`)
	matches := re.FindStringSubmatch(sql)

	if len(matches) > 0 {
		table = matches[1]
	}
	return action, table
}

func parseCreate(sql string) (string, string) {
	var table string
	var action string

	re := regexp.MustCompile(`create +(table|index|unique index) "?([[a-z0-9_\.]*)"?(.*)`)
	matches := re.FindStringSubmatch(sql)

	if len(matches) > 0 {
		table = matches[2]

		switch matches[1] {
		case "table":
			action = "create table"
		case "index":
			action = "create index"
		case "unique index":
			action = "create index"
		}
	}
	return action, table
}

func parseDrop(sql string) (string, string) {
	var table string
	var action string

	re := regexp.MustCompile(`drop (table|schema)( if exists)? "?([[a-z0-9_\.]*)"?(.*)`)
	matches := re.FindStringSubmatch(sql)

	if len(matches) > 0 {
		table = matches[3]
		switch matches[1] {
		case "table":
			action = "drop table"
		case "schema":
			action = "drop schema"
		}
	}
	return action, table
}

func parseAlter(sql string) (string, string) {
	var table string
	action := "alter table"

	re := regexp.MustCompile(`alter table "?([[a-z0-9_\.]*)"?(.*)`)
	matches := re.FindStringSubmatch(sql)

	if len(matches) > 0 {
		table = matches[1]
	}
	return action, table
}

func getConnection() *sql.DB {
	connection, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return connection
}
