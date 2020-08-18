package main

import (
	"fmt"
	"os"
	"strings"
	"io"
	"bufio"
	"regexp"
	"flag"
)

const LINE_LENGTH = 80

func print_break(length int) {
	fmt.Printf("%s\n", strings.Repeat("-", length))
}

func print_wrapped(line string, length int, quote string) {
	len_quote := len(quote)
	buffer := quote
	for index, rune := range line[len_quote:] {
		buffer += string(rune)
		if (index + 1) % (length - len_quote) == 0 {
			fmt.Printf("%s\n", buffer)
			buffer = quote
		}
	}
	if buffer != "" {
		fmt.Printf("%s\n", buffer)
	}
}

func textwrap_stream(reader io.Reader, length int) {
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
				print_break(length)
			} else {
				quote := re_quote.FindString(line)
				print_wrapped(line, length, quote)
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

func textwrap_file(filename string) {
	// Check file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("cannot read file '%s'\n", filename)
		os.Exit(1)
	}
	defer file.Close()

	textwrap_stream(file, LINE_LENGTH)
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

	textwrap_stream(os.Stdin, *length)
}

