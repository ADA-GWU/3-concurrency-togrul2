package main

import (
	"assignment3/processor"
	"log"
	"os"
	"strconv"
)

func main() {
	// Parse command arguments
	if len(os.Args) != 4 {
		log.Fatal("Invalid number of arguments! Hint: main.exe [image path] [size of square (int)] [mode('S' | 'M')]")
	}

	fileName := os.Args[1]
	squareSize, sizeParseErr := strconv.Atoi(os.Args[2])
	mode := os.Args[3]

	if sizeParseErr != nil {
		log.Fatal("Error! Provide a valid integer for square size parameter.")
	}

	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	defer file.Close()

	if err := processor.ProcessImage(file, squareSize, mode); err != nil {
		log.Fatal(err)
	}
}
