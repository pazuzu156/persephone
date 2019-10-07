package fm

// TopAlbum holds top album info.
type TopAlbum struct {
	Rank      string `xml:"rank,attr"`
	Name      string `xml:"name"`
	PlayCount string `xml:"playcount"`
	Mbid      string `xml:"mbid"`
	URL       string `xml:"url"`
	Artist    struct {
		Name string `xml:"name"`
		Mbid string `xml:"mbid"`
		URL  string `xml:"url"`
	} `xml:"artist"`
	Images []struct {
		Size string `xml:"size,attr"`
		URL  string `xml:",chardata"`
	} `xml:"image"`
}

// Track holds track info.
type Track struct {
	NowPlaying string `xml:"nowplaying,attr,omitempty"`
	Artist     struct {
		Name string `xml:",chardata"`
		Mbid string `xml:"mbid,attr"`
	} `xml:"artist"`
	Name       string `xml:"name"`
	Streamable string `xml:"streamable"`
	Mbid       string `xml:"mbid"`
	Album      struct {
		Name string `xml:",chardata"`
		Mbid string `xml:"mbid,attr"`
	} `xml:"album"`
	URL    string `xml:"url"`
	Images []struct {
		Size string `xml:"size,attr"`
		URL  string `xml:",chardata"`
	} `xml:"image"`
	Date struct {
		Uts  string `xml:"uts,attr"`
		Date string `xml:",chardata"`
	} `xml:"date"`
}

// TopTrack holds top track info.
type TopTrack struct {
	Rank       string `xml:"rank,attr"`
	Name       string `xml:"name"`
	Duration   string `xml:"duration"`
	PlayCount  string `xml:"playcount"`
	Mbid       string `xml:"mbid"`
	URL        string `xml:"url"`
	Streamable struct {
		FullTrack  string `xml:"fulltrack,attr"`
		Streamable string `xml:",chardata"`
	} `xml:"streamable"`
	Artist struct {
		Name string `xml:"name"`
		Mbid string `xml:"mbid"`
		URL  string `xml:"url"`
	} `xml:"artist"`
	Images []struct {
		Size string `xml:"size,attr"`
		URL  string `xml:",chardata"`
	} `xml:"image"`
}

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

// YouTubeResponse holds youtube response info.
type YouTubeResponse struct {
	Kind string `json:"kind"`
	ETag struct {
		NextPageToken string `json:"nextPageToken"`
		RegionCode    string `json:"regionCode"`
	} `json:"etag"`
	PageInfo struct {
		TotalResults   string `json:"totalResults"`
		ResultsPerPage string `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind string `json:"kind"`
		ETag string `json:"etag"`
		ID   struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt string `json:"publishedAt"`
			ChannelID   string `json:"channelId"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
				ChannelTitle         string `json:"channelTitle"`
				LiveBroadcastContent string `json:"liveBroadcastContent"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}
