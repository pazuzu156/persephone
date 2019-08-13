package lib

type TopAlbum struct {
	Rank      string `xml:"rank,attr"`
	Name      string `xml:"name"`
	PlayCount string `xml:"playcount"`
	Mbid      string `xml:"mbid"`
	Url       string `xml:"url"`
	Artist    struct {
		Name string `xml:"name"`
		Mbid string `xml:"mbid"`
		Url  string `xml:"url"`
	} `xml:"artist"`
	Images []struct {
		Size string `xml:"size,attr"`
		Url  string `xml:",chardata"`
	} `xml:"image"`
}

type Track struct {
	NowPlaying string "xml:\"nowplaying,attr,omitempty\""
	Artist     struct {
		Name string "xml:\",chardata\""
		Mbid string "xml:\"mbid,attr\""
	} "xml:\"artist\""
	Name       string "xml:\"name\""
	Streamable string "xml:\"streamable\""
	Mbid       string "xml:\"mbid\""
	Album      struct {
		Name string "xml:\",chardata\""
		Mbid string "xml:\"mbid,attr\""
	} "xml:\"album\""
	Url    string "xml:\"url\""
	Images []struct {
		Size string "xml:\"size,attr\""
		Url  string "xml:\",chardata\""
	} "xml:\"image\""
	Date struct {
		Uts  string "xml:\"uts,attr\""
		Date string "xml:\",chardata\""
	} "xml:\"date\""
}

type TopTrack struct {
	Rank       string `xml:"rank,attr"`
	Name       string `xml:"name"`
	Duration   string `xml:"duration"`
	PlayCount  string `xml:"playcount"`
	Mbid       string `xml:"mbid"`
	Url        string `xml:"url"`
	Streamable struct {
		FullTrack  string `xml:"fulltrack,attr"`
		Streamable string `xml:",chardata"`
	} `xml:"streamable"`
	Artist struct {
		Name string `xml:"name"`
		Mbid string `xml:"mbid"`
		Url  string `xml:"url"`
	} `xml:"artist"`
	Images []struct {
		Size string `xml:"size,attr"`
		Url  string `xml:",chardata"`
	} `xml:"image"`
}

type AlbumPosition struct {
	X      int
	Y      int
	Shadow Shadow
	Info   InfoText
}

type TrackPosition struct {
	X     float64
	Y     float64
	Plays PlaysText
}

type Shadow struct {
	X float64
	Y float64
	R float64
}

type InfoText struct {
	X     float64
	Y     float64
	Plays PlaysText
}

type PlaysText struct {
	X float64
	Y float64
}
