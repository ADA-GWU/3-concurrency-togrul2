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
	// Entrypoint for our app where GUI and image processing work together with the help of channels.
	// Usage of channels is required for updating processed image in GUI in a real time.
	go func() {
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

		w := app.NewWindow()
		// Run GUI loop.
		if err := gui.RunGUIEventLoop(w, img, func(imageChannel chan<- image.Image, errorsChannel chan<- error) {
			processor.ProcessImage(img, squareSize, mode, imageChannel, errorsChannel)
		}); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
