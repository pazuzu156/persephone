package models

import (
	"persephone/utils"

	"github.com/andersfylling/disgord"
)

type User struct {
	ID        int64 `db:"pk"`
	Username  string
	DiscordID uint64
	Lastfm    string
}

func GetUser(user *disgord.User) {
    db, _ := utils.OpenDB()
    defer db.Close()

    
}
