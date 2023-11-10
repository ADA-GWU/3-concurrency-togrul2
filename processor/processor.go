// Processor package with processImage function to process image for assignment3
package processor

import (
	"assignment3/arg"
	"assignment3/pixel"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"
)

// Finds the average color of each square of image with given
// size and draws new image with the calculated squares.
func ProcessImage(
	img image.Image,
	size int,
	mode arg.RunMode,
	messageChannel chan<- image.Image,
	errorsChannel chan<- error,
) {
	switch mode {
	case arg.SingleThreaded:
		processImageWithSingleThread(img, size, messageChannel, errorsChannel)
	case arg.MultiThreaded:
		// processImageWithMultithreads(img, size, messageChannel, errorsChannel)
	default:
		errorsChannel <- fmt.Errorf("invalid value for mode parameter")
	}
}

// Creates rgba image from matrix of colors.

func processImageWithSingleThread(
	img image.Image,
	size int,
	updateChannel chan<- image.Image,
	errorsChannel chan<- error,
) {
	pixels := pixel.GetImagePixels(img)

	var resultImg *image.RGBA
	// Iteratively process image squares.
	for i := 0; i < len(pixels); i += size {
		for j := 0; j < len(pixels[0]); j += size {
			processSquare(pixels, i, j, i+size, j+size)
			// Send new image to img update channel to update it in gui.
			resultImg = createRGBAImage(pixels)
			time.Sleep(time.Second / 5)
			updateChannel <- resultImg

		}
	}

	// Save the results to the file.
	resultFile, fileErr := os.Create("result.jpg")
	if fileErr != nil {
		errorsChannel <- fileErr
	}

	if encodeErr := png.Encode(resultFile, resultImg); encodeErr != nil {
		errorsChannel <- encodeErr
	}
	close(updateChannel)
	close(errorsChannel)
	fmt.Println("Saved result in result.jpg file")
}

func processImageWithMultithreads(
	img image.Image,
	size int,
	imageChannel chan<- image.Image,
	errorsChannel chan<- error,
) {
	//TODO: Implement
}
