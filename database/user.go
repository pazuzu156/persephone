package database

import (
	"strconv"

	"github.com/andersfylling/disgord"
)

// User object.
type User struct {
	ID        int64 `db:"pk"`
	Username  string
	DiscordID uint64
	Lastfm    string
}

// GetUser gets the database user via a Discord user.
func GetUser(user *disgord.User) User {
	db, err := OpenDB()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	var dbu []User
	db.Select(&dbu, db.Where("discord_id", "=", GetUInt64ID(user)))

	if len(dbu) > 0 {
		return dbu[0]
	}

	return User{}
}

// GetUsers returns all the users in the database.
func GetUsers() []User {
	db, err := OpenDB()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	var dbu []User
	db.Select(&dbu)

	return dbu
}

// GetUInt64ID returns a uint64 version of the Discord user ID.
func GetUInt64ID(user *disgord.User) uint64 {
	did, _ := strconv.Atoi(user.ID.String())

	return uint64(did)
}
