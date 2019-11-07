package lib

import "github.com/andersfylling/disgord"

// Servers object.
type Servers struct {
	ID           int64 `db:"pk"`
	GuildID      uint64
	LogChannelID uint64
	ElevatedRole uint64
	Time
}

func init() {
	var err error
	db, err = OpenDB()
	Check(err)
}

// GetServer returns a server logged in the database.
func GetServer(guildID disgord.Snowflake) (server Servers) {
	var servers []Servers
	err := db.Select(&servers, db.From(Servers{}), db.Where("guild_id", "=", guildID))
	Check(err)

	if len(servers) == 1 {
		server = servers[0]
	}

	return
}
