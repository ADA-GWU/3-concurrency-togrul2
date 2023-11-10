package processor

import (
	"image"
	"image/color"
)

// Finds average color in given bounds of given matrix of colors.
func calculateAverageColor(pixels [][]color.Color, rstart, cstart, rend, cend int) color.Color {
	var totalR, totalG, totalB, totalA, pixelCount int64 = 0, 0, 0, 0, 0
	for i := rstart; i < rend; i++ {
		for j := cstart; j < cend; j++ {
			if i < len(pixels) && j < len(pixels[0]) {
				pixelColor := color.RGBAModel.Convert(pixels[i][j]).(color.RGBA)
				totalR += int64(pixelColor.R)
				totalG += int64(pixelColor.G)
				totalB += int64(pixelColor.B)
				totalA += int64(pixelColor.A)
				pixelCount++
			}
		}
	}

	averageR := uint8(totalR / pixelCount)
	averageG := uint8(totalG / pixelCount)
	averageB := uint8(totalB / pixelCount)
	averageA := uint8(totalA / pixelCount)
	return color.RGBA{R: averageR, G: averageG, B: averageB, A: averageA}
}

func processSquare(pixels [][]color.Color, rstart, cstart, rend, cend int) {
	// Find average color of square.
	averageColor := calculateAverageColor(pixels, rstart, cstart, rend, cend)
	// Set all square pixels to the average color.
	for i := rstart; i < rend; i++ {
		for j := cstart; j < cend; j++ {
			if i < len(pixels) && j < len(pixels[0]) {
				pixels[i][j] = averageColor
			}
		}
	}
}

// Creates RGBA image from matrix of colors.
func createRGBAImage(pixels [][]color.Color) *image.RGBA {
	// Create image.
	img := image.NewRGBA(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{len(pixels[0]), len(pixels)},
		},
	)
	// Set pixels.
	for i := 0; i < len(pixels); i++ {
		for j := 0; j < len(pixels[0]); j++ {
			img.Set(j, i, pixels[i][j])
		}
	}
	return img
}
