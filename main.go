package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jessevdk/go-flags"
	zipcracker "github.com/tahadostifam/ZipCracker/zip_cracker"
)

var threadsCount int

const threadsCountDefault = 1

var opts struct {
	Passwd       string `long:"passwd" description:"Your password list" required:"true"`
	ZipFile      string `long:"zip_file" description:"Your zip file" required:"true"`
	ThreadsCount string `long:"threads" short:"t" description:"Threads count of cracker"`
}

func checkFilesExistence() {
	errors := []string{}
	_, passwdExistsErr := os.Stat(opts.Passwd)
	_, zipFileExistsErr := os.Stat(opts.ZipFile)

	if os.IsNotExist(passwdExistsErr) {
		errors = append(errors, "Passwd list file not exists")
	}

	if os.IsNotExist(zipFileExistsErr) {
		errors = append(errors, "Zip file not exists")
	}

	if len(errors) > 0 {
		fmt.Println(strings.Join(errors, "\n"))
		os.Exit(1)
	}
}

func checkThreadsCount() {
	if len(strings.TrimSpace(opts.ThreadsCount)) > 0 {
		tc, err := strconv.Atoi(opts.ThreadsCount)
		if err == nil {
			threadsCount = tc
		} else {
			threadsCount = threadsCountDefault
		}
	} else {
		threadsCount = threadsCountDefault
	}
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		os.Exit(1)
	}

	checkFilesExistence()
	checkThreadsCount()
	startCracker()
}

func startCracker() {
	cracker := zipcracker.ZipCracker{
		PasswdListPath: opts.Passwd,
		ZipFilePath:    opts.ZipFile,
		ThreadsCount:   threadsCount,
	}

	cracker.Start()

	fmt.Println(opts.Passwd, opts.ZipFile, threadsCount)
}
