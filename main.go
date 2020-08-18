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

func make_breakline(length int) string {
	return strings.Repeat("-", length)
}

func make_wrappedline(line string, length int, re_quote *regexp.Regexp) []string {
	offset := 0
	prefix := ""

	quote := re_quote.FindString(line)
	len_quote := len(quote)
	if len_quote != 0 && len_quote < length {
		offset = len_quote
		prefix = quote
	}

	buffer := []string{prefix}
	line_number := 0
	for index, rune := range line[offset:] {
		buffer[line_number] += string(rune)
		if (index + 1) % (length - offset) == 0 {
			buffer = append(buffer, prefix)
			line_number += 1
		}
	}

	return buffer
}

func wrap_array(lines []string, length int) ([]string, error) {
	// Compile regular expressions
	re_quote, err := regexp.Compile("^([> ]*)")
	if err != nil {
		return nil, err
	}
	re_break, err := regexp.Compile("^(?:-{5,}|={5,})$")
	if err != nil {
		return nil, err
	}

	wrapped := []string{}

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		if len(line) > length {
			if re_break.MatchString(line) {
				wrapped = append(wrapped, make_breakline(length))
			} else {
				wrapped = append(wrapped, make_wrappedline(line, length, re_quote)...)
			}
		} else {
			wrapped = append(wrapped, line)
		}
	}

	return wrapped, nil
}

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
				fmt.Printf("%s\n", make_breakline(length))
			} else {
				for _, wrapped := range make_wrappedline(line, length, re_quote) {
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

