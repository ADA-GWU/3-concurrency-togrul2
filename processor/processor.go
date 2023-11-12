// Processor package with processImage function to process image for assignment3
package processor

import (
	"assignment3/arg"
	"assignment3/pixel"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"sync"
	"time"
)

// Finds the average color of each square of image with given
// size and draws new image with the calculated squares.
func ProcessImage(
	img image.Image,
	size int,
	mode arg.RunMode,
	imagesChannel chan<- image.Image,
	errorsChannel chan<- error,
) {
	pixels := pixel.GetImagePixels(img)
	if mode == arg.SingleThreaded {
		processImageWithSingleThread(pixels, size, imagesChannel, errorsChannel)
	} else {
		processImageWithMultithreads(pixels, size, imagesChannel, errorsChannel)
	}
}

// Saves the image to the file, returns error if any.
func saveToFile(img image.Image) error {
	resultFile, fileErr := os.Create("result.jpg")
	if fileErr != nil {
		return fileErr
	}

	if encodeErr := png.Encode(resultFile, img); encodeErr != nil {
		return encodeErr
	}

	fmt.Println("Saved result to result.jpg file.")
	return nil
}

// Creates rgba image from matrix of colors in singlethread way.
func processImageWithSingleThread(
	pixels [][]color.Color,
	size int,
	imagesChannel chan<- image.Image,
	errorsChannel chan<- error,
) {
	var resultImg *image.RGBA
	// Iteratively process image squares.
	for i := 0; i < len(pixels); i += size {
		for j := 0; j < len(pixels[0]); j += size {
			processSquare(pixels, i, j, i+size, j+size)
			// Send new image to img update channel to update it in gui.
			resultImg = createRGBAImage(pixels)
			imagesChannel <- resultImg
			time.Sleep(time.Second / 25)
		}
	}

	// Save the results to the file.
	if saveErr := saveToFile(resultImg); saveErr != nil {
		errorsChannel <- saveErr
	}

	// Close channels to indicate end of work.
	close(imagesChannel)
	close(errorsChannel)
}

// Creates rgba image from matrix of colors in multithreaded way.
func processImageWithMultithreads(
	pixels [][]color.Color,
	size int,
	imagesChannel chan<- image.Image,
	errorsChannel chan<- error,
) {
	var resultImg *image.RGBA
	var wg sync.WaitGroup
	var matrixLock sync.Mutex

	maxGorutines := runtime.NumCPU()
	// This channel helps us to limit number of running goroutines to number of available threads in machine.
	gorutineGuard := make(chan struct{}, maxGorutines)

	for i := 0; i < len(pixels); i += size {
		wg.Add(1)
		gorutineGuard <- struct{}{}
		go func(row int) {
			defer wg.Done()
			for j := 0; j < len(pixels[0]); j += size {
				// Acquire lock for pixels var.
				matrixLock.Lock()
				processSquare(pixels, row, j, row+size, j+size)
				resultImg = createRGBAImage(pixels)
				// After the job is done, release the lock.
				matrixLock.Unlock()

				// Send new image to img update channel to update it in gui.
				imagesChannel <- resultImg
				time.Sleep(time.Second / 25)
			}
			<-gorutineGuard
		}(i)
	}

	// Wait until all fragments are loaded asyncronously
	wg.Wait()

	// Save the results to the file.
	if saveErr := saveToFile(resultImg); saveErr != nil {
		errorsChannel <- saveErr
	}

	close(imagesChannel)
	close(errorsChannel)
}
