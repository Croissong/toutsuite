// +build !prod
//go:generate purse -in="sql_fixtures" -out="sql_fixtures.go" -pack="main" -name="ps"   

package main

import (
	"github.com/smotes/purse"
	"path/filepath"
	"os"
)

var ps purse.Purse

func createPurse() {
	var err error
	ps, err = purse.New(filepath.Join(".", "sql_fixtures"))
	if(err != nil) {
		printErr(err)
		os.Exit(1)
	}
}
