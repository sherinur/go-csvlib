package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrQuote      = errors.New("excess or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")
)

type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	// GetField(n int) (string, error)
	// GetNumberOfFields() int
}

type SimpleCSVParser struct{}

// Reads one line from open input file.
// Calling ReadLine in a loop allows you to sequentially read each line from the file,
// continuing until the end of the file is reached.
// Returns pointer to line, with terminator removed, or nil if EOF occurred.
func (parser SimpleCSVParser) ReadLine(r io.Reader) (string, error) {
	buffer := make([]byte, 1)
	var line []byte
	var lastByte byte
	numOfQuotes := 0

	for {
		n, err := r.Read(buffer)
		// reading not empty byte
		if n > 0 {
			b := buffer[0]
			line = append(line, b)

			if b == '"' {
				numOfQuotes++
			}

			if b == '\n' || b == '\r' {
				lastByte = b
				break
			}
		}

		// error handling
		if err != nil {
			if err == io.EOF {
				// check for mismatched quotes
				if numOfQuotes%2 != 0 {
					return "", ErrQuote
				}

				if len(line) > 0 {
					break
				}
				return "", io.EOF
			}
			return "", err
		}

	}

	// remove newline from the end.
	if lastByte == '\n' && len(line) > 0 && line[len(line)-1] == '\r' {
		// \r\n
		line = line[:len(line)-2]
	} else if lastByte == 'n' || lastByte == 'r' {
		// \r or \n
		if len(line) > 0 {
			line = line[:len(line)-1]
		}
	}

	// check for mismatched quotes
	if numOfQuotes%2 != 0 {
		return "", ErrQuote
	}

	return string(line), nil
}

func main() {
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var csvparser CSVParser = SimpleCSVParser{}

	for {
		line, err := csvparser.ReadLine(file)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading line:", err)
			return
		}
		fmt.Print(line)
	}
}
