package database

import (
	"persephone/lib"
	"strconv"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/lastfm-go"
)

// Users object.
type Users struct {
	ID                  int64 `db:"pk"`
	Username            string
	Email               string
	DiscordID           uint64 `db:"unique"`
	DiscordToken        string
	DiscordRefreshToken string
	Lastfm              string
	LastfmToken         string
	RememberToken       *string
	Time
}

func init() {
	var err error
	db, err = OpenDB()
	lib.Check(err)
}

// GetUser gets the database user via a Discord user.
func GetUser(user *disgord.User) Users {
	var dbu []Users
	db.Select(&dbu, db.Where("discord_id", "=", GetUInt64ID(user)))

	if len(dbu) > 0 {
		return dbu[0]
	}

	return Users{}
}

// GetUsers returns all the users in the database.
func GetUsers() []Users {
	var dbu []Users
	db.Select(&dbu)

	return dbu
}

// GetLastfmUserInfo gets user info from last.fm.
func GetLastfmUserInfo(user *disgord.User, lfm *lastfm.API) (lastfm.UserGetInfo, error) {
	dbu := GetUser(user)

	return lfm.User.GetInfo(lastfm.P{"user": dbu.Lastfm})
}

// GetUInt64ID returns a uint64 version of the Discord user ID.
func GetUInt64ID(user *disgord.User) uint64 {
	did, _ := strconv.Atoi(user.ID.String())

	return uint64(did)
}

// Crown is a relational function to get crowns model
func (c Users) Crown(id int64) Crowns {
	crowns := GetCrownsList()

	for _, crown := range crowns {
		if crown.ID == id {
			return crown
		}
	}

	return Crowns{}
}

// Crowns is a relational function to get crowns model
func (c Users) Crowns() (crowns []Crowns) {
	dbc := GetCrownsList()

	for _, crown := range dbc {
		if c.DiscordID == crown.DiscordID {
			crowns = append(crowns, crown)
		}
	}

	return
}
