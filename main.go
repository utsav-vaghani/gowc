package main

import (
	"fmt"
	"github.com/utsav-vaghani/gowc/driver"
	"os"
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
