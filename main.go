package main

import (
	"os"

	"github.com/mdayat/sorta/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
