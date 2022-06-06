package zipcracker

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gosuri/uilive"
)

type ZipCracker struct {
	PasswdListPath string
	ThreadsCount   int
}

func (c *ZipCracker) Start() {
	passwdList, passwdListErr := ReadPasswdList(c.PasswdListPath)
	if passwdListErr != nil {
		log.Fatal(passwdListErr)
	}

	fmt.Println(passwdList)

	writer := uilive.New()
	writer.Start()

	for i := 0; i <= 100; i++ {
		fmt.Fprintf(writer, "Downloading.. (%d/%d) GB\n", i, 100)
		time.Sleep(time.Millisecond * 5)
	}

	fmt.Fprintln(writer, "Finished: Downloaded 100GB")
	writer.Stop()
}

func ReadPasswdList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}

	defer file.Close()

	var lines []string = []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		if len(strings.TrimSpace(currentLine)) > 0 {
			lines = append(lines, currentLine)
		}
	}

	return lines, nil
}
