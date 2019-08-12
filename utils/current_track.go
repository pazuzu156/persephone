package utils

import (
	"errors"
	"persephone/database"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/lastfm-go"
)

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

// GetNowPlayingTrack returns the currently playing track
func GetNowPlayingTrack(author *disgord.User, lfm *lastfm.Api) (Track, error) {
	tracks, err := GetRecentTracks(author, lfm, "1")

	if len(tracks) > 0 {
		track := tracks[0]

		if track.NowPlaying == "true" {
			return track, nil
		}

		return Track{}, errors.New("You're not currently listening to anything")
	}

	return Track{}, err
}

func GetRecentTracks(author *disgord.User, lfm *lastfm.Api, limit string) ([]Track, error) {
	if user := database.GetUser(author); user.Username != "" {
		np, _ := lfm.User.GetRecentTracks(lastfm.P{
			"user":  user.Lastfm,
			"limit": limit,
		})

		var tracks = []Track{} // for the return

		if len(np.Tracks) > 0 {
			for _, track := range np.Tracks {
				tracks = append(tracks, track)
			}

			return tracks, nil
		}
	}

	return []Track{}, errors.New("You're not currently logged in with Last.fm")
}
