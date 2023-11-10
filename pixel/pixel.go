package pixel

import (
	"image"
	"image/color"
)

// Get matrix of pixels of image.
func GetImagePixels(img image.Image) [][]color.Color {
	bounds := img.Bounds()
	height := bounds.Max.Y
	width := bounds.Max.X

	// Create matrix of pixels and fill its cells with colors from correspoding coordinates.
	pixels := make([][]color.Color, height)
	for i := 0; i < height; i++ {
		pixels[i] = make([]color.Color, width) // Create a row.
		for j := 0; j < width; j++ {
			pixels[i][j] = img.At(j, i)
		}
	}
	return pixels
}
