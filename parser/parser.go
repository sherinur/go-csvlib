package parser

import (
	"errors"
	"io"
)

type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	GetField(n int) (string, error)
	GetNumberOfFields() int
}

type SimpleCSVParser struct {
	lastLine       string
	fields         []string
	numFields      int
	readLineCalled bool
}

var (
	ErrQuote      = errors.New("excess or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")
)

// Reads one line from open input file.
// Calling ReadLine in a loop allows you to sequentially read each line from the file,
// continuing until the end of the file is reached.
// Returns pointer to line, with terminator removed, or nil if EOF occurred.
func (parser *SimpleCSVParser) ReadLine(r io.Reader) (string, error) {
	buffer := make([]byte, 1)
	var line []byte
	var lastByte byte
	numOfQuotes := 0
	inQuotes := false

	for {
		n, err := r.Read(buffer)
		// reading not empty byte
		if n > 0 {
			b := buffer[0]
			if len(line) == 0 && (b == '\n' || b == '\r') {
				continue
			}

			line = append(line, b)

			if b == '"' {
				inQuotes = !inQuotes
			}

			if (b == '\n' || b == '\r') && !inQuotes {
				lastByte = b
				break
			}
		}

		// error handling
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					// process the last line if it exists
					parser.lastLine = string(line)
					parser.fields = extractFields(parser.lastLine)
					parser.numFields = len(parser.fields)
					parser.readLineCalled = true
					return parser.lastLine, nil
				}
				// check for mismatched quotes
				parser.lastLine = string(line)
				parser.fields = extractFields(parser.lastLine)
				parser.numFields = len(parser.fields)
				parser.readLineCalled = true
				return "", io.EOF
			}
			parser.lastLine = ""
			parser.numFields = 0
			parser.readLineCalled = true
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

	// store lastline
	parser.lastLine = string(line)
	parser.fields = extractFields(parser.lastLine)
	parser.numFields = len(parser.fields)
	parser.readLineCalled = true

	return parser.lastLine, nil
}

// Returns n-th field from last line read by ReadLine;
// Returns ErrFieldCount if n < 0 or beyond last field
// Fields are separated by commas
// Fields may be surrounded by "..."; such quotes are removed
// There can be an arbitrary number of fields of any length
func (parser *SimpleCSVParser) GetField(n int) (string, error) {
	if !parser.readLineCalled {
		return "", errors.New("ReadLine has not been called")
	}

	if n < 0 || n >= len(parser.fields) {
		return "", ErrFieldCount
	}
	return parser.fields[n], nil
}

// Returns number of fields on last line read by ReadLine
// Returns -1 if called before ReadLine is called
func (parser *SimpleCSVParser) GetNumberOfFields() int {
	if !parser.readLineCalled {
		return -1
	}
	return parser.numFields
}
