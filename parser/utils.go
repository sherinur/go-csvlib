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
				fields = append(fields, trimQuotes(field))
				field = ""
			}
		case '"':
			inQuotes = !inQuotes
		default:
			field += string(ch)
		}
	}

	if len(field) != 0 {
		fields = append(fields, trimQuotes(field))
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
