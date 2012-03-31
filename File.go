// original: https://github.com/dbravender/go_mapreduce/blob/master/src/file_iter.go
// and: https://gist.github.com/1364737

package File

import (
	"bufio"
	"os"
)

func EachLine(filename string) chan string {
	output := make(chan string)
	go func() {
		file, err := os.Open(filename, os.O_RDONLY, 0)
		if err != nil {
			return
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			output <- line
			if err == os.EOF {
				break
			}
		}
		close(output)
	}()
	return output
}
