package database

import "persephone/lib"

type ArtistImage struct {
	ID     int64 `db:"pk"`
	Artist string
	MaID   int64 `db:"unique"`
	Time
}

func init() {
	var err error
	db, err = OpenDB()
	lib.Check(err)
}
