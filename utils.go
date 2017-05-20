package main

import (
	"fmt"
)

func printErr(err error) error {
	fmt.Println(fmt.Sprintf("Error: %s", err))
	return err
}
