package main

import (
	"kreator/internal/cli"
	"log"
	"os"
)

func main() {
	err := cli.ParseCommands(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
