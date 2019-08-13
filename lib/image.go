package lib

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// GetExt returns the extension of a given file name
func GetExt(filename string) string {
	s := strings.Split(filename, ".")

	return s[len(s)-1]
}

// OpenImage returns an image.Image instance of a given file
func OpenImage(filename string) image.Image {
	in, _ := os.Open(filename)
	defer in.Close()
	var img image.Image

	switch GetExt(filename) {
	case "png":
		img, _ = png.Decode(in)
		break
	case "jpeg":
		img, _ = jpeg.Decode(in)
		break
	case "jpg":
		img, _ = jpeg.Decode(in)
		break
	}

	return img
}
