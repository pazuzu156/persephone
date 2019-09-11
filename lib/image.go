package lib

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/url"
	"os"
	"strings"

	"github.com/cavaliercoder/grab"
	"github.com/fogleman/gg"
	"github.com/gocolly/colly"
	"github.com/pazuzu156/lastfm-go"
)

// NoArtistURL is the URL to a blank image for no artist images found
// on metal archives.
const NoArtistURL = "https://cdn.kalebklein.com/images/bm.png"

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
	dc.LoadFontFace(LocGet("static/fonts/NotoSans-Regular.ttf"), 14)
	dc.SetRGB(0.9, 0.9, 0.9)
	dc.DrawStringAnchored(
		"Persephone: discord.gg/BtqjBDu",
		float64(dc.Width()), float64(dc.Height()),
		1.04, -1.2,
	)
}

// GetArtistImageURL returns the URL for an artist scraped from metal-archives.
func GetArtistImageURL(artist lastfm.ArtistGetInfo) string {
	// what shall we scrape for?
	col := colly.NewCollector()
	var imgsrc string
	col.OnHTML(".band_img a#photo", func(e *colly.HTMLElement) {
		imgsrc = e.ChildAttr("img", "src")
	})

	maartist := GetMaArtist(artist.Name) // look up a band from artists.json

	if maartist.ID != 0 {
		col.Visit(fmt.Sprintf("https://metal-archives.com/bands/%s/%d", url.QueryEscape(maartist.Name), maartist.ID))
	} else {
		col.Visit(fmt.Sprintf("https://metal-archives.com/bands/%s", url.QueryEscape(artist.Name)))
	}

	// Empty image
	if imgsrc == "" {
		imgsrc = NoArtistURL
	}

	return imgsrc
}

// GetArtistImage scrapes metal-archives for an artist image.
func GetArtistImage(artist lastfm.ArtistGetInfo) image.Image {
	imgsrc := GetArtistImageURL(artist)

	if imgsrc != NoArtistURL {
		res, _ := grab.Get(LocGet("temp/"), imgsrc)
		img := OpenImage(res.Filename)

		os.Remove(res.Filename)

		return img
	}
	aimg := OpenImage(LocGet("static/images/bm.png"))

	return aimg
}