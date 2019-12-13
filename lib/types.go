package lib

// Artist holds simple artist info.
type Artist struct {
	Rank       string `xml:"rank,attr"`
	Name       string `xml:"name"`
	PlayCount  string `xml:"playcount"`
	Mbid       string `xml:"mbid"`
	URL        string `xml:"url"`
	Streamable string `xml:"streamable"`
	Images     []struct {
		Size string `xml:"size,attr"`
		URL  string `xml:",chardata"`
	} `xml:"image"`
}

// Artists holds simple artist info in slice format.
type Artists []struct {
	Rank       string `xml:"rank,attr"`
	Name       string `xml:"name"`
	PlayCount  string `xml:"playcount"`
	Mbid       string `xml:"mbid"`
	URL        string `xml:"url"`
	Streamable string `xml:"streamable"`
	Images     []struct {
		Size string `xml:"size,attr"`
		URL  string `xml:",chardata"`
	} `xml:"image"`
}

// AlbumPosition holds album art positions.
type AlbumPosition struct {
	X      int
	Y      int
	Shadow Shadow
	Info   InfoText
}

// TrackPosition holds track label positions.
type TrackPosition struct {
	X     float64
	Y     float64
	Plays PlaysText
}

// Shadow holds image shadow positions.
type Shadow struct {
	X float64
	Y float64
	R float64
}

// InfoText holds info text positions.
type InfoText struct {
	X     float64
	Y     float64
	Plays PlaysText
}

// PlaysText holds play count text positions.
type PlaysText struct {
	X float64
	Y float64
}
