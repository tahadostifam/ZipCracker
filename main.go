package main

import (
	zipcracker "github.com/tahadostifam/ZipCracker/zip_cracker"
	"github.com/urfave/cli/v2"
)

func main() {
	// Init of Cli Tool
	cli := &cli.App{
		Name:  "ZipCracker",
		Usage: "",
	}

	// Starting ZipCracker
	cracker := zipcracker.ZipCracker{
		PasswdListPath: "",
		ThreadsCount:   4,
	}

	cracker.Start()
}
