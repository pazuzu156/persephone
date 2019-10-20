package lib

// ArtistImages object.
type ArtistImages struct {
	ID     int64 `db:"pk"`
	Artist string
	MaID   int64 `db:"unique"`
	Time
}

func init() {
	var err error
	db, err = OpenDB()
	Check(err)
}
