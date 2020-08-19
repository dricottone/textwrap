package main

import (
	"fmt"
	"os"
	"strings"
	"io"
	"bufio"
	"regexp"
	"flag"

	"git.dominic-ricottone.com/textwrap/common"
)

const LINE_LENGTH = 80

func wrap_stream(reader io.Reader, length int) {
	// Create scanner from reader
	input := bufio.NewScanner(reader)

	// Compile regular expressions
	re_quote, err := regexp.Compile("^([> ]*)")
	if err != nil {
		fmt.Printf("internal error - %v\n", err)
		os.Exit(1)
	}
	re_break, err := regexp.Compile("^(?:-{5,}|={5,})$")
	if err != nil {
		fmt.Printf("internal error - %v\n", err)
		os.Exit(1)
	}

	// Scan line by line
	for input.Scan() {
		line := input.Text()
		line = strings.TrimSpace(line)

		if len(line) > length {
			if re_break.MatchString(line) {
				fmt.Printf("%s\n", common.MakeBreakline(length))
			} else {
				for _, wrapped := range common.MakeWrappedLine(line, length, re_quote) {
					fmt.Printf("%s\n", wrapped)
				}
			}
		} else {
			fmt.Printf("%s\n", line)
		}
	}

	// Check for scanner errors
	if err = input.Err(); err != nil {
		fmt.Printf("internal error - %v\n", err)
		os.Exit(1)
	}
}

func wrap_file(filename string) {
	// Check file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("cannot read file '%s'\n", filename)
		os.Exit(1)
	}
	defer file.Close()

	wrap_stream(file, LINE_LENGTH)
}

func main() {
	// Check STDIN
	_, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("cannot read input")
		os.Exit(1)
	}

	// Look for arguments
	var length = flag.Int("length", LINE_LENGTH, "maximum length of lines")
	flag.Parse()

	wrap_stream(os.Stdin, *length)
}

