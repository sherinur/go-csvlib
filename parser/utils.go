package parser

// Extracts fields from the line and returns array of strings
func extractFields(line string) []string {
	if len(line) == 0 {
		return nil
	}

	var fields []string
	field := ""
	inQuotes := false
	for _, ch := range line {
		switch ch {
		case ',':
			if inQuotes {
				field += string(ch)
			} else {
				fields = append(fields, field)
				field = ""
			}
		case '"':
			inQuotes = !inQuotes
			field += string(ch)
		default:
			field += string(ch)
		}
	}

	if len(field) != 0 || line[len(line)-1] == ',' {
		fields = append(fields, field)
	}

	return fields
}

// Removes leading and trailing quotes and returns string
func trimQuotes(field string) string {
	if len(field) > 1 && field[0] == '"' && field[len(field)-1] == '"' {
		return field[1 : len(field)-1]
	}

	return field
}

// Removes newline from the end
func trimNewLine(line []byte) []byte {
	if len(line) > 0 {
		lastByte := line[len(line)-1]

		switch lastByte {
		case '\n':
			line = line[:len(line)-1]
			if len(line) > 0 && line[len(line)-1] == '\r' {
				line = line[:len(line)-1]
			}
		case '\r':
			line = line[:len(line)-1]
		}
	}
	return line
}
