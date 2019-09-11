package database

import (
	"persephone/lib"
)

// Crowns object.
type Crowns struct {
	ID        int64 `db:"pk"`
	DiscordID uint64
	Artist    string
	PlayCount int
	Time
}

func init() {
	var err error
	db, err = OpenDB()
	lib.Check(err)
}

// GetCrownsList returns a list of all crowns in database
func GetCrownsList() (crowns []Crowns) {
	err := db.Select(&crowns, db.From(Crowns{}))
	lib.Check(err)

	return
}

// User is a relational method to retrieve the user from a given crown.
func (c Crowns) User() (user Users) {
	err := db.Select(&user, db.From(Users{}), db.Where("discord_id", "=", c.DiscordID))
	lib.Check(err)

	return
}

// GetUserCrowns is a relational method that returns a list of crowns for a given user.
func (c Crowns) GetUserCrowns(sql ...interface{}) (crowns []Crowns) {
	user := c.User()
	db.Select(&crowns, db.From(Crowns{}), db.Where("discord_id", "=", user.DiscordID))
	return
}
