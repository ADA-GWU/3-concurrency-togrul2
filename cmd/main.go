package main

import (
	"assignment3/processor"
	"log"
	"os"
	"strconv"
)

func main() {
	// Parse command arguments
	fileName := os.Args[0]
	squareSize, sizeParseErr := strconv.Atoi(os.Args[1])
	mode := os.Args[2]

	if sizeParseErr != nil {
		log.Fatal("Error! Provide a valid integer for square size parameter.")
		os.Exit(1)
	}

	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		log.Fatal(fileErr)
		os.Exit(1)
	}
	defer file.Close()

	if err := processor.ProcessImage(file, squareSize, mode); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
