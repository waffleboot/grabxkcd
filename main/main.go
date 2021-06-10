package main

import (
	"os"

	"github.com/waffleboot/grabxkcd/grabxkcd"
)

func main() {
	os.Exit(grabxkcd.CLI(os.Args[1:]...))
}
