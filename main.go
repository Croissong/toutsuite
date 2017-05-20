package main

import (
  "os"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
  "github.com/urfave/cli"
	"github.com/smotes/purse"
	"path/filepath"
	"fmt"
	"strings"
)

const db_name = "toutsuite"
var ps purse.Purse
var db *sql.DB

func main() {
	db = openDb()
	defer db.Close()
  app := cli.NewApp()
	createCommands(app)
	app.Run(os.Args)
}

func createCommands(app *cli.App) {
	  app.Commands = []cli.Command{
			{
      Name:    "insert",
      Aliases: []string{"i"},
      Usage:   "insert suite",
      Action:  insertSuite,
			},
			{
      Name:    "random",
      Aliases: []string{"r"},
      Usage:   "get random suite",
      Action:  getRandomeSuite,
    },
		}
}

func openDb() *sql.DB {
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s.db", db_name))
	if err != nil {
		log.Fatal(err)
	}
	ps, err = purse.New(filepath.Join(".", "sql"))
	create_db, _ := ps.Get("create_db.sql")
	_, err = db.Exec(create_db)
	if err != nil && !checkIsError(err, "table toutsuite already exists") {
		fmt.Printf("sql: error %s", err)
		return nil
	}
	return db
}

func checkIsError(err error, msg string) bool {
	return strings.Contains(err.Error(), msg)
}
