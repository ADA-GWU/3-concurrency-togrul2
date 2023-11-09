package pixel

import (
	"image"
	"image/color"
	_ "image/jpeg" // For processing jpeg format.
	"io"
)

// Get matrix of pixels of image.
func GetImagePixels(imageFile io.Reader) ([][]color.Color, error) {
	img, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	height := bounds.Max.Y
	width := bounds.Max.X
	// Create slice of pixels
	pixels := make([][]color.Color, height)
	for i := 0; i < height; i++ {
		// Create row of pixel matrix
		pixels[i] = make([]color.Color, width)
		for j := 0; j < width; j++ {
			pixels[i][j] = img.At(j, i)
		}
	}
	return pixels, nil
}
