package main

import (
	"fmt"
	"log"

	"github.com/mattn/unko/vm"
)

func main() {
	env := vm.NewEnv()
	v, err := env.Execute(`foo=1; foo+3`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}
