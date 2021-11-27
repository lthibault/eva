package main

import (
	"log"

	"github.com/lthibault/eva"
)

func main() {
	var vm eva.VM

	err := vm.Exec(`
		42
	`)

	if err != nil {
		log.Fatal(err)
	}
}
