package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// File struct type
type File struct {
	Name string
	Key  string
	Logo string
}

// NewFile constructs a new File struct with the given Name, Key, and Logo
func NewFile(name string, key string, logo string) *File {
	f := File{Name: name, Key: key, Logo: logo}
	return &f
}

func main() {

	info := NewFile("ioio-logo.script", "{LOGO}", "usr/share/plymouth/theme/ioio-logo/ioio-logo.png")

	input, err := ioutil.ReadFile(info.Name)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(info.Key), []byte(info.Logo), -1)

	if err = ioutil.WriteFile(info.Name, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
