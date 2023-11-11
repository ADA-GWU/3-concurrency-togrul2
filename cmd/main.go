package main

import (
	"assignment3/arg"
	"assignment3/gui"
	"assignment3/processor"
	"image"
	"log"
	"os"

	"gioui.org/app"
)

func main() {
	// Parse command arguments.
	fileName, squareSize, mode, argParseErr := arg.ParseArguments()
	if argParseErr != nil {
		log.Fatal(argParseErr)
	}

	// Open file and close after program exists.
	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	defer file.Close()

	// Decode the image from source file.
	img, _, imageDecodeErr := image.Decode(file)
	if imageDecodeErr != nil {
		log.Fatal(imageDecodeErr)
	}

	// Entrypoint for our app where GUI and image processing work together with the help of channels.
	// Usage of channels is required for updating processed image in GUI in a real time.
	go func() {
		w := app.NewWindow()

		// Image processing job.
		job := func(imgChannel chan<- image.Image, errorsChannel chan<- error) {
			processor.ProcessImage(img, squareSize, mode, imgChannel, errorsChannel)
		}

		// Run GUI loop in the main thread.
		if err := gui.RunGUIEventLoop(w, &img, job); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
