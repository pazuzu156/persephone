package database

import (
	"persephone/lib"
)

// Crown object.
type Crown struct {
	ID        int64 `db:"pk"`
	DiscordID uint64
	Artist    string
	PlayCount int
}

func init() {
	var err error
	db, err = OpenDB()
	lib.Check(err)
}

// GetCrownsList returns a list of all crowns in database
func GetCrownsList() (crowns []Crown) {
	err := db.Select(&crowns, db.From(Crown{}))
	lib.Check(err)

	return
}

// User is a relational method to retrieve the user from a given crown.
func (c Crown) User() (user User) {
	err := db.Select(&user, db.From(User{}), db.Where("discord_id", "=", c.DiscordID))
	lib.Check(err)

	return
}

// GetUserCrowns is a relational method that returns a list of crowns for a given user.
func (c Crown) GetUserCrowns(sql ...interface{}) (crowns []Crown) {
	user := c.User()
	db.Select(&crowns, db.From(Crown{}), db.Where("discord_id", "=", user.DiscordID))
	return
}
