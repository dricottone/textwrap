package common

import (
	"strings"
	"regexp"
)

func MakeBreakline(length int) string {
	return strings.Repeat("-", length)
}

func MakeWrappedLine(line string, length int, re_quote *regexp.Regexp) []string {
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

func WrapArray(lines []string, length int) ([]string, error) {
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
				wrapped = append(wrapped, MakeBreakline(length))
			} else {
				wrapped = append(wrapped, MakeWrappedLine(line, length, re_quote)...)
			}
		} else {
			wrapped = append(wrapped, line)
		}
	}

	return wrapped, nil
}

