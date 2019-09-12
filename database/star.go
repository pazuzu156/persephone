package database

import "persephone/lib"

// Stars object.
type Stars struct {
	ID                 int64 `db:"pk"`
	MessageID          int64
	StarboardMessageID int64
	Time
}

func init() {
	var err error
	db, err = OpenDB()
	lib.Check(err)
}
