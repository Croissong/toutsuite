// +build !prod
//go:generate purse -in="sql_fixtures" -out="sql_fixtures.go" -pack="main" -name="ps"

package main

import (
	"github.com/smotes/purse"
	"os"
	"path/filepath"
)

var ps = createPurse()

func createPurse() purse.Purse {
	ps, err := purse.New(filepath.Join(".", "sql_fixtures"))
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	return ps
}
