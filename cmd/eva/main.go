package main

import (
	"fmt"
	"log"

	"github.com/lthibault/eva"
)

func main() {
	var vm eva.VM

	v, err := vm.Exec(`
		42
	`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v.Int32())
}
