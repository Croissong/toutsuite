package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/chzyer/readline"
	"github.com/urfave/cli"
	"strings"
)

func insertSuite(c *cli.Context) error {
	url, err := clipboard.ReadAll()
	if err != nil {
		return err
	}
	if !checkUrl(url) {
		fmt.Println("Invalid url")
		return nil
	}
	success, id := writeSuiteDB(url)
	if success {
		fmt.Println("inserted suite")
		insertMeta(id)
		printRow(id)
	}
	return nil
}

func insertMeta(id int64) bool {
	title := readMeta()
	insert_meta, _ := ps.Get("insert_meta.sql")
	row, err := db.Exec(insert_meta, title, id)
	if err != nil {
		if checkIsError(err, "UNIQUE constraint failed") {
			fmt.Println("name already exists")
		} else {
			printErr(err)
		}
		return false
	}
	id, err = row.LastInsertId()
	return true
}

func printRow(id int64) {
	var title string
	var url string
	var tags string
	row := db.QueryRow("select title, url, tags from toutsuite where id = $1", id)
	err := row.Scan(&title, &url, &tags)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(title, url, tags)
	}
}

func writeSuiteDB(url string) (bool, int64) {
	insert_suite, _ := ps.Get("insert_suite.sql")
	row, err := db.Exec(insert_suite, url)
	if err != nil {
		if checkIsError(err, "UNIQUE constraint failed") {
			fmt.Println("url already exists")
		} else {
			printErr(err)
		}
		return false, 0
	}
	var id int64
	id, err = row.LastInsertId()
	if err != nil {
		return false, 0
	} else {
		return true, id
	}
}

func readMeta() string {
	name := readName()
	return name
}

func readName() string {
	rl, err := readline.New("Name: ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	line, err := rl.Readline()
	if err != nil { // io.EOF
		return ""
	}
	return line
}

func checkUrl(url string) bool {
	return strings.HasPrefix(url, "http")
}
