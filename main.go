package main

import (
	"./parser"
	"database/sql"
	"github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

const (
	postgresConnection = "user=emailmarketing dbname=psqlog sslmode=disable"
	gophers_count      = 10
)

var lineRegexp = regexp.MustCompile(`([0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}) brt \[[0-9]*\]: \[(.*?)\] user=(.*?),db=(.*?) log:  duration: (.*?) ms  (.*?): (.*)`)
var sqlRegexp = regexp.MustCompile(`([a-z]*)(.*)`)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) < 2 {
		log.Fatal("You need to set the filename\nexample: psqlog main.log")
	}

	lines := redFile(os.Args[1])

	var waitGroup sync.WaitGroup

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
	// log.Println("Starting:", i)
	db := getConnection()
	defer db.Close()

	txn := startTransaction(db)
	stmt := copyPrepare(txn)

	for _, line := range lines {
		parseLine(strings.ToLower(line), stmt)
	}

	copyClose(stmt)
	commitTransaction(txn)

	db.Close()
	// log.Println("Finishing:", i)
	waitGroup.Done()
}

func parseLine(line string, stmt *sql.Stmt) {
	logLine := parse.LogLine{Line: line}
	logLine.Parse()
	if logLine.Action != "" && logLine.Table != "" {
		_, err := stmt.Exec(logLine.Username, logLine.Database, logLine.Duration, logLine.Action, logLine.Table, logLine.Query, logLine.Created_at)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func copyPrepare(transaction *sql.Tx) *sql.Stmt {
	stmt, err := transaction.Prepare(pq.CopyIn("logs", "username", "database", "duration", "action", "table_name", "sql", "created_at"))
	if err != nil {
		log.Fatal(err)
	}
	return stmt
}

func copyClose(stmt *sql.Stmt) {
	_, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func startTransaction(db *sql.DB) *sql.Tx {
	transaction, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	return transaction
}

func commitTransaction(txn *sql.Tx) {
	err := txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func getConnection() *sql.DB {
	connection, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return connection
}
