package processor

import (
	"image/color"
)

func calculateAverageColor(pixels [][]color.Color, rstart, cstart, rend, cend int) color.Color {
	var totalR, totalG, totalB, totalA, pixelCount int64
	for i := rstart; i < rend; i++ {
		for j := cstart; j < cend; j++ {
			if i >= len(pixels) || j >= len(pixels[0]) {
				continue
			}
			pixelColor := color.RGBAModel.Convert(pixels[i][j]).(color.RGBA)
			totalR += int64(pixelColor.R)
			totalG += int64(pixelColor.G)
			totalB += int64(pixelColor.B)
			totalA += int64(pixelColor.A)
			pixelCount++
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
			if i >= len(pixels) || j >= len(pixels[0]) {
				continue
			}

			pixels[i][j] = averageColor
		}
	}
}
