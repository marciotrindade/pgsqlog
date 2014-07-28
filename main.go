package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"github.com/lib/pq"
	"github.com/marciotrindade/pgsqlog/parser"
)

const (
	postgresConnection = "user=marciotrindade dbname=psqlog sslmode=disable"
	gophersCount       = 10
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

	waitGroup.Add(gophersCount)

	linesCount := len(lines)
	partialCount := linesCount / gophersCount

	for i := 0; i < gophersCount; i++ {
		var finish int
		start := i * partialCount
		if i == (gophersCount - 1) {
			finish = linesCount
		} else {
			finish = ((i + 1) * partialCount)
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
	logLine := parser.LogLine{Line: line}
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
