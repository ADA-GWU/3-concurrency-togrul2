// Processor package with processImage function to process image for assignment3
package processor

import (
	"assignment3/pixel"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"strings"
)

// Finds the average color of each square of image with given
// size and draws new image with the calculated squares.
func ProcessImage(file io.Reader, size int, mode string) error {
	if strings.ToLower(mode) == "s" {
		if err := processImageWithSingleThread(file, size); err != nil {
			return err
		}
	} else if strings.ToLower(mode) == "m" {
		if err := processImageWithMultithreads(file, size); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("invalid value for mode parameter")
	}
	return nil
}

func processImageWithSingleThread(file io.Reader, size int) error {
	pixels, err := pixel.GetImagePixels(file)
	if err != nil {
		return err
	}
	for i := 0; i < len(pixels); i += size {
		for j := 0; j < len(pixels[0]); j += size {
			processSquare(pixels, i, j, i+size, j+size)
		}
	}

	img := image.NewRGBA(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{len(pixels[0]), len(pixels)},
		},
	)
	for i := 0; i < len(pixels); i++ {
		for j := 0; j < len(pixels[0]); j++ {
			img.Set(j, i, pixels[i][j])
		}
	}
	f, fileErr := os.Create("result.jpg")
	if fileErr != nil {
		return fileErr
	}
	if encodeErr := png.Encode(f, img); encodeErr != nil {
		return encodeErr
	}
	return nil
}

func processImageWithMultithreads(file io.Reader, size int) error {
	// pixels, err := pixel.GetImagePixels(file)
	return nil
}
