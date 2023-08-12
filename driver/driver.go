package driver

import (
	"fmt"
	"github.com/utsav-vaghani/gowc/reader"
	"github.com/utsav-vaghani/gowc/reader/file"
	"github.com/utsav-vaghani/gowc/reader/stdin"
	"strings"
)

type driver struct {
	flagMap map[string]string
}

func New(flagMap map[string]string) Driver {
	return &driver{
		flagMap: flagMap,
	}
}

// Run takes actual argument passed while running the program
func (d *driver) Run(args []string) error {
	flags, filePaths, err := d.parser(args)
	if err != nil {
		return err
	}

	var reader reader.IReader

	if len(filePaths) == 0 {
		reader, err = stdin.NewReader()
		if err != nil {
			return err
		}

		processFlags(reader, flags)
		return nil
	}

	for _, filePath := range filePaths {
		reader, err = file.NewReader(filePath)
		if err != nil {
			return err
		}

		processFlags(reader, flags)
	}

	return nil
}

func processFlags(reader reader.IReader, flags map[string]struct{}) {
	for flag := range flags {
		switch flag {
		case "-l":
			err := reader.CountLines()
			if err != nil {
				fmt.Printf("error while counting lines: %v\n", err)
				continue
			}
		case "-w":
			err := reader.CountWords()
			if err != nil {
				fmt.Printf("error while counting words: %v\n", err)
				continue
			}
		case "-c":
			err := reader.CountBytes()
			if err != nil {
				fmt.Printf("error while counting bytes: %v\n", err)
				continue
			}
		}
	}
}

func (d *driver) parser(args []string) (map[string]struct{}, []string, error) {
	flags := make(map[string]struct{})
	filePaths := make([]string, 0)

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if _, ok := d.flagMap[arg]; !ok {
				return nil, nil, fmt.Errorf("invalid flag `%s` which is not supported\nusage: wc [-clw] [file ...]\n", arg)
			}

			flags[arg] = struct{}{}
		} else {
			filePaths = append(filePaths, arg)
		}
	}

	return flags, filePaths, nil
}
