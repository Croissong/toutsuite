package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os/exec"
)

func getRandomeSuite(c *cli.Context) error {
	get_random_suite, _ := ps.Get("get_random_suite.sql")
	var url string
	var title string
	row := db.QueryRow(get_random_suite)
	err := row.Scan(&title, &url)
	if err != nil {
		return printErr(err)
	} else {
		var out []byte
		out, err = exec.Command("google-chrome", "--incognito", url).Output()
		if err != nil {
			fmt.Printf("%s", out)
			return err
		}
		return nil
	}
}
