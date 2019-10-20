package lib

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
	Check(err)
}

// GetCrownsList returns a list of all crowns in database
func GetCrownsList() (crowns []Crowns) {
	err := db.Select(&crowns, db.From(Crowns{}))
	Check(err)

	return
}

// User is a relational method to retrieve the user from a given crown.
func (c Crowns) User() (user Users) {
	var users []Users
	err := db.Select(&users, db.From(Users{}), db.Where("discord_id", "=", c.DiscordID))
	Check(err)

	if len(users) == 1 {
		user = users[0]
	}

	return
}

// GetUserCrowns is a relational method that returns a list of crowns for a given user.
func (c Crowns) GetUserCrowns(sql ...interface{}) (crowns []Crowns) {
	user := c.User()
	db.Select(&crowns, db.From(Crowns{}), db.Where("discord_id", "=", user.DiscordID))
	return
}
