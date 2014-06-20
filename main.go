package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You need to set the filename\nexample: ./clean_csv sample.csv")
		return
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		parseLine(strings.ToLower(line), i)
	}
	// fmt.Println("Total of lines:", len(lines))
}

func parseLine(line string, lineNumber int) {
	re := regexp.MustCompile(`([0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}) brt \[[0-9]*\]: \[(.*?)\] user=(.*?),db=(.*?) log:  duration: (.*?) ms  (.*?): (.*)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) > 0 {
		date := matches[1]
		user := matches[3]
		db := matches[4]
		t, _ := time.ParseDuration(matches[5] + "ms")
		time := strconv.FormatFloat(t.Seconds(), 'f', 3, 64)
		sql := strings.Trim(matches[7], " ")

		action, table := parseSql(sql)
		if action != "" && table != "" {
			fmt.Println(date, user, db, time, action, table, sql)
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
