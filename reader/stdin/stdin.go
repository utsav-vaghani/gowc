package stdin

import (
	"bufio"
	"fmt"
	"github.com/utsav-vaghani/gowc/reader"
	"os"
	"strings"
)

type stdinReader struct {
	bufferSize int
	data       []string
}

func NewReader() (reader.IReader, error) {
	return &stdinReader{}, nil
}

func (r *stdinReader) Read() ([]string, error) {
	if r.data != nil {
		return r.data, nil
	}

	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 0)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	r.data = data

	return r.data, nil
}

func (r *stdinReader) CountLines() error {
	lines, err := r.Read()
	if err != nil {
		return err
	}

	fmt.Printf("total number of lines: %v\n", len(lines))
	return nil
}

func (r *stdinReader) CountWords() error {
	lines, err := r.Read()
	if err != nil {
		return err
	}

	var words int

	for _, line := range lines {
		words += len(strings.Fields(line))
	}

	fmt.Printf("total number of words: %v\n", words)
	return nil
}

func (r *stdinReader) CountBytes() error {
	lines, err := r.Read()
	if err != nil {
		return err
	}

	var bytes int

	for _, line := range lines {
		bytes += len(line) + 1
	}

	fmt.Printf("total number of bytes: %v\n", bytes)
	return nil
}
