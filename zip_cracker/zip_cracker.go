package zipcracker

import (
	"bufio"
	"log"
	"os"
	"strings"

	color "github.com/fatih/color"
	"github.com/gosuri/uilive"
	zip "github.com/yeka/zip"
)

var zipFilePath string

type ZipCracker struct {
	PasswdListPath string
	ZipFilePath    string
	ThreadsCount   int
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

func doUnzip(passwd string) bool {
	r, _ := zip.OpenReader(zipFilePath)
	defer r.Close()

	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword(passwd)
		}

		_, err := f.Open()
		if err != nil {
			return true
		}
	}

	return false
}

type ThreadMessage struct {
	finish bool
	index  int
}

func (c *ZipCracker) Start() {
	passwdList, passwdListErr := ReadPasswdList(c.PasswdListPath)
	if passwdListErr != nil {
		log.Fatal(passwdListErr)
	}

	zipFilePath = c.ZipFilePath
	writer := uilive.New()
	writer.Start()
	passwdListLen := len(passwdList)
	threadMessages := make(chan ThreadMessage)

	for j := 0; j < c.ThreadsCount; j++ {
		go func() {
			for i := 0; i < len(passwdList); i++ {
				badPasswd := doUnzip(passwdList[i])

				if i == passwdListLen-1 {
					writer.Stop()
					color.Red("Cracking failed!\nPassword not found in password list.")
					os.Exit(0)
				}

				if badPasswd {
					// password is not valid
					threadMessages <- ThreadMessage{
						finish: false,
						index:  i,
					}
				} else {
					// password is valid
					color.Green("Congratulations! Password Found!\n")
					color.Green("Password: %s\n", string(passwdList[i]))
					os.Exit(0)
				}
			}

		}()
	}

	for {
		select {
		case m := <-threadMessages:
			if m.finish {
				writer.Stop()
				os.Exit(0)
				return
			} else {
				color.New(color.FgRed).Fprintf(writer, "Remaining %d of %d\n", m.index, passwdListLen)
			}
		}
	}
}
