package lib

import (
    "errors"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/lastfm-go"
)

// GetNowPlayingTrack returns the currently playing track
func GetNowPlayingTrack(author *disgord.User, lfm *lastfm.API) (Track, error) {
	tracks, err := GetRecentTracks(author, lfm, "3")

	if err == nil {
		if len(tracks) > 0 {
			track := tracks[0]

			if track.NowPlaying == "true" {
				return track, nil
			}

			return Track{}, errors.New("You're not currently listening to anything")
		}
	}

	return Track{}, err
}

// GetRecentTracks retrieves a users recently scrobbled tracks.
func GetRecentTracks(author *disgord.User, lfm *lastfm.API, limit string) ([]Track, error) {
	if user := GetUser(author); user.Username != "" {
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
