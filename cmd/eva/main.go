package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/lthibault/eva"
)

const program = `
	(+ "Hello, " "Eva!")
`

func main() {
	var vm eva.VM
	v, err := vm.Exec(strings.NewReader(program))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v.String())
}
