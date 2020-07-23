package main

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/sky0621/wolf-bff/src/setup"
)

func main() {
	os.Exit(int(setup.ExecMain()))
}
