// original: https://github.com/dbravender/go_mapreduce/blob/master/src/file_iter.go
// and: https://gist.github.com/1364737

package File

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func EachLine(filename string) chan string {
	output := make(chan string)
	go func() {
		file, err := os.Open(filename)
		if err != nil {
			return
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			output <- strings.TrimRight(line, "\n\r")
			if err == io.EOF {
				break
			}
		}
		close(output)
	}()
	return output
}

func Exists(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}
