package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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
			fmt.Println("Updating...")
			return text
		}

	}

}

func main() {

	// Variables more for clarity than anything
	logoFile := "ioio-logo.script"
	logoName := `logo_filename = .*;`
	logoPath := `logo_filename = "` + Reader() + `";`

	// Construct new File
	info := NewFile(logoFile, logoName, logoPath)

	input, err := ioutil.ReadFile(info.Name)
	// Run input
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Construct regex with logoName
	reg := regexp.MustCompile(info.Key)

	output := reg.ReplaceAll(input, []byte(info.Logo))
	// Run output
	if err = ioutil.WriteFile(info.Name, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Boot logo successfully updated!")
	}
}
