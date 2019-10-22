package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
		fmt.Print("Enter new logo file-> ")
		text, _ := reader.ReadString('\n')

		// Convert input CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		// If input char count not 0 then return it as the new logo path
		if FileExists(text) {
			fmt.Println("Updating...")
			return text
		}

		// Return error message should FileExists = false
		fmt.Println(`File "` + text + `" either does not exist or is a directory.`)
	}
}

// FileExists checks to see that input string file exists
func FileExists(file string) bool {
	exists, err := os.Stat(file)

	// Returns false should file not exist or is directory
	if os.IsNotExist(err) {
		return false
	}
	return !exists.IsDir()
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
		log.Fatalf("Could not read ioio-logo.script file. System returned the following: %s\n", err)
	}

	// Construct regex with logoName
	reg := regexp.MustCompile(info.Key)

	output := reg.ReplaceAll(input, []byte(info.Logo))
	// Run output
	if err = ioutil.WriteFile(info.Name, output, 0666); err != nil {
		log.Fatalf("Could not update ioio-logo.script file. System returned the following: %s\n", err)
	}

	cmd := exec.Command("update-initramfs", "-u")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Could not update system init config. Ensure ioio-boot is being run with elevated permissions. System returned the following: %s\n", err)
	}
	fmt.Println("Boot logo successfully updated!")
}
