package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/utsav-vaghani/gowc/src/reader"
)

type fileReader struct {
	path string
}

func NewReader(path string) (reader.IReader, error) {
	fr := &fileReader{
		path: path,
	}

	return fr, nil
}

func (r *fileReader) Read() ([]string, error) {
	file, err := os.Open(r.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err != nil {
		return nil, err
	}

	scanner.Buffer(make([]byte, 0, 1024*1024*1024), 1024*1024*1024)

	data := make([]string, 0)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func (r *fileReader) CountLines() error {
	lines, err := r.Read()
	if err != nil {
		return err
	}

	fmt.Printf("file %s:: total number of lines: %v\n", r.path, len(lines))
	return nil
}

func (r *fileReader) CountWords() error {
	lines, err := r.Read()
	if err != nil {
		return err
	}

	var words int

	for _, line := range lines {
		words += len(strings.Fields(line))
	}

	fmt.Printf("file %s:: total number of lines: %v\n", r.path, words)
	return nil
}

func (r *fileReader) CountBytes() error {
	lines, err := r.Read()
	if err != nil {
		return err
	}

	var bytes int

	for _, line := range lines {
		bytes += len(line) + 1
	}

	fmt.Printf("file %s:: total number of bytes: %v\n", r.path, bytes)
	return nil
}
