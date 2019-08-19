package utils

import (
	"errors"
	"persephone/database"
	"persephone/lib"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/lastfm-go"
)

// GetNowPlayingTrack returns the currently playing track
func GetNowPlayingTrack(author *disgord.User, lfm *lastfm.API) (lib.Track, error) {
	tracks, err := GetRecentTracks(author, lfm, "3")

	if err == nil {
		if len(tracks) > 0 {
			track := tracks[0]

			if track.NowPlaying == "true" {
				return track, nil
			}

			return lib.Track{}, errors.New("You're not currently listening to anything")
		}
	}

	return lib.Track{}, err
}

// GetRecentTracks retrieves a users recently scrobbled tracks.
func GetRecentTracks(author *disgord.User, lfm *lastfm.API, limit string) ([]lib.Track, error) {
	if user := database.GetUser(author); user.Username != "" {
		np, _ := lfm.User.GetRecentTracks(lastfm.P{
			"user":  user.Lastfm,
			"limit": limit,
		})

		var tracks = []lib.Track{} // for the return

		if len(np.Tracks) > 0 {
			for _, track := range np.Tracks {
				tracks = append(tracks, track)
			}

			return tracks, nil
		}
	}

	return []lib.Track{}, errors.New("You're not currently logged in with Last.fm")
}
