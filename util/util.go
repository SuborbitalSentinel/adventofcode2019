package util

import (
	"bufio"
	"os"
)

// Readlines returns a channel containing the content of the file
func Readlines(path string) <-chan string {
	fobj, err := os.Open(path)
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(fobj)
	if err := scanner.Err(); err != nil {
		return nil
	}

	chnl := make(chan string)
	go func() {
		for scanner.Scan() {
			chnl <- scanner.Text()
		}
		close(chnl)
	}()

	return chnl
}
