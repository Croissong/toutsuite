package main

import (
	"github.com/urfave/cli"
	"github.com/atotto/clipboard"
	"fmt"
	"os/exec"
	"strings"
)

func insertSuite(c *cli.Context) error {
	url, err := clipboard.ReadAll()
	if(err != nil) { return err }
	if(!checkUrl(url)) {
		fmt.Println("Invalid url")
		return nil
	}
	insert_entry, _ := ps.Get("insert_suite.sql")
	_, err = db.Exec(fmt.Sprintf(insert_entry, url))
	if err != nil {
		if(checkIsError(err, "UNIQUE constraint failed")) {
			fmt.Println(fmt.Sprintf("%s already exists", url))
			return nil
		} else { return printErr(err) }
	}
	fmt.Println("inserted suite")
	return nil
}

func checkUrl(url string) bool {
	return strings.HasPrefix(url, "http")
}

func getRandomeSuite(c *cli.Context) error {
	get_random_suite, _ := ps.Get("get_random_suite.sql")
	var url string
	row, err := db.Query(get_random_suite)
	row.Next()
	row.Scan(&url)
	if err != nil {
		return printErr(err)
	} else {
		var out []byte
		out, err = exec.Command("google-chrome", "--incognito", url).Output()
		if(err != nil) {
			fmt.Printf("%s", out)
			return err
		}
		return nil
		}
}
