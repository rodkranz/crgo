package main

import (
	"os"
	"log"
	"github.com/rodkranz/crgo/crgo"
)

func main() {
	if err := crgo.Run(os.Args[1:]); err != nil {
		log.Printf("err: %v ", err)
		os.Exit(1)
	}
	os.Exit(0)
}
