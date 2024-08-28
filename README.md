# go-csvlib
A simple library in GO, to manually process and parse CSV files.

## Context
Comma-separated values, or CSV, is a simple and widely used format for representing tabular data. Each row in a CSV file corresponds to a line of text, with individual fields separated by commas. Here’s an example:
  - Name,Age,Occupation
  - John Doe,28,Engineer
  - Jane Doe,32,Designer
  - Sam Smith,24,Developer

## Resources
[Go File I/O](https://golang.org/pkg/os/)

[CSV Format](https://tools.ietf.org/html/rfc4180)

[Go Interfaces](https://golang.org/doc/effective_go.html#interfaces)


## Mandatory Part
I have built a CSV library in Go. Implemented following interface methods:
```GO
    type CSVParser interface  {
        ReadLine(r io.Reader) (string, error)
        GetField(n int) (string, error)
        GetNumberOfFields() int
    }
    
    var (
        ErrQuote      = errors.New("excess or missing \" in quoted-field")
        ErrFieldCount = errors.New("wrong number of fields")
    )
```

### ReadLine
This function reads a new line from a CSV file.
  ```GO
    type CSVParser interface  {
      // ...
      ReadLine(r io.Reader) (string, error)
      // ...
  }
  ```

  - Reads one line from open input file
  - Returns pointer to line, with terminator removed, or nil if EOF occurred
  - Calling ReadLine in a loop allows you to sequentially read each line from the file, continuing until the end of the file is reached.
  - Assumes that input lines are terminated by \r, \n, \r\n, or EOF
  - If the line has a missing or extra quote, it should return an empty string and an ErrQuote error.
