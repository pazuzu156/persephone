package lib

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/fogleman/gg"
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

// BrandImage tags an image with a discord link and bot's name
func BrandImage(dc *gg.Context) {
	dc.LoadFontFace("static/fonts/NotoSans-Regular.ttf", 14)
	dc.SetRGB(0.9, 0.9, 0.9)
	dc.DrawStringAnchored(
		"Persephone: discord.gg/BtqjBDu",
		float64(dc.Width()), float64(dc.Height()),
		1.04, -1.2,
	)
}
