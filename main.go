package main

import (
	"fmt"
	"io"
	"os"

	"a-library-for-others/parser"
)

func main() {
	file, _ := os.Open("tests/testdata/example.csv")
	defer file.Close()

	var csvparser parser.CSVParser = &parser.SimpleCSVParser{}
	for {
		line, err := csvparser.ReadLine(file)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println("Past line:", line)
		for i := 0; i < csvparser.GetNumberOfFields(); i++ {
			Field, err := csvparser.GetField(i)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Field", i, ":", Field)
		}
	}
}
