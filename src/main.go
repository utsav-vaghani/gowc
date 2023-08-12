package main

import (
	"fmt"
	"os"

	"github.com/utsav-vaghani/gowc/src/driver"
)

func main() {
	driver := driver.New(map[string]string{
		"-l": "counts the number of lines",
		"-w": "counts the number of words",
		"-c": "counts the number of bytes",
	})

	err := driver.Run(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
}
