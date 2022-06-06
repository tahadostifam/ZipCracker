package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	zipcracker "github.com/tahadostifam/ZipCracker/zip_cracker"
)

const PASSWD_LIST_PATH = "files/passwd_list.txt"

func TestReadPasswdList(t *testing.T) {
	cwd, _ := os.Getwd()

	list, err := zipcracker.ReadPasswdList(
		fmt.Sprintf(
			"%s/../%s",
			cwd,
			PASSWD_LIST_PATH,
		),
	)

	if err != nil {
		t.Fatal(err)
	} else {
		log.Printf("Count of passwords: %d", len(list))
	}
}
