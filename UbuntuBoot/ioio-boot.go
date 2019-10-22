package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

// Reader takes console input, converts to LF and returns as string
func Reader() string {
	reader := bufio.NewReader(os.Stdin)

	// Looping to wait for input
	for {
		// Request input
		fmt.Print("Enter new logo path-> ")
		text, _ := reader.ReadString('\n')

		// Convert input CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		// If input char count not 0 then return it as the new logo path
		if strings.Compare("", text) != 0 {
			fmt.Println("Building...")
			return text
		}

	}

}

func main() {

	logoPath := Reader()

	info := NewFile("ioio-logo.script", "{LOGO}", logoPath)

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
