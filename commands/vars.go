package commands

import "github.com/pazuzu156/persephone/lib"

var (
	commands = []CommandItem{}
	config   = lib.Config()

	// FontRegular is the name for the regular typed font.
	FontRegular = lib.LocGet("static/fonts/NotoSans-Regular.ttf")

	// FontBold is the name for the bold typed font.
	FontBold = lib.LocGet("static/fonts/NotoSans-Bold.ttf")

	// AlbumPositions is the album grid positions.
	AlbumPositions = []lib.AlbumPosition{
		{
			X: 355,
			Y: 170,
			Shadow: lib.Shadow{
				X: 350,
				Y: 165,
				R: 10,
			},
			Info: lib.InfoText{
				X: 350,
				Y: 340,
				Plays: lib.PlaysText{
					X: 350,
					Y: 360,
				},
			},
		},
		{
			X: 555,
			Y: 170,
			Shadow: lib.Shadow{
				X: 550,
				Y: 165,
				R: 10,
			},
			Info: lib.InfoText{
				X: 550,
				Y: 340,
				Plays: lib.PlaysText{
					X: 550,
					Y: 360,
				},
			},
		},
		{
			X: 355,
			Y: 390,
			Shadow: lib.Shadow{
				X: 350,
				Y: 385,
				R: 10,
			},
			Info: lib.InfoText{
				X: 350,
				Y: 560,
				Plays: lib.PlaysText{
					X: 350,
					Y: 580,
				},
			},
		},
		{
			X: 555,
			Y: 390,
			Shadow: lib.Shadow{
				X: 550,
				Y: 385,
				R: 10,
			},
			Info: lib.InfoText{
				X: 550,
				Y: 560,
				Plays: lib.PlaysText{
					X: 550,
					Y: 580,
				},
			},
		},
	}

	// TrackPositions is the positions for track listings.
	TrackPositions = []lib.TrackPosition{
		{
			X: 720,
			Y: 180,
			Plays: lib.PlaysText{
				X: 735,
				Y: 200,
			},
		},
		{
			X: 720,
			Y: 230,
			Plays: lib.PlaysText{
				X: 735,
				Y: 250,
			},
		},
		{
			X: 720,
			Y: 280,
			Plays: lib.PlaysText{
				X: 735,
				Y: 300,
			},
		},
		{
			X: 720,
			Y: 330,
			Plays: lib.PlaysText{
				X: 735,
				Y: 350,
			},
		},
		{
			X: 720,
			Y: 380,
			Plays: lib.PlaysText{
				X: 735,
				Y: 400,
			},
		},
		{
			X: 720,
			Y: 430,
			Plays: lib.PlaysText{
				X: 735,
				Y: 450,
			},
		},
		{
			X: 720,
			Y: 480,
			Plays: lib.PlaysText{
				X: 735,
				Y: 500,
			},
		},
		{
			X: 720,
			Y: 530,
			Plays: lib.PlaysText{
				X: 735,
				Y: 550,
			},
		},
	}
)
