package commands

import (
	"github.com/pazuzu156/aurora"
)

// Crowns command.
type Crowns struct {
	Command Command
}

// InitCrowns initializes the crowns command.
func InitCrowns() Crowns {
	return Crowns{Init(
		"crowns",
		"List your crowns",
		[]UsageItem{},
	)}
}

// Register registers and runs the crowns command.
func (c Crowns) Register() *aurora.Command {
	c.Command.CommandInterface.Run = func(ctx aurora.Context) {
		// crowns := database.GetUser(ctx.Message.Author).Crowns()

		// if len(crowns) > 0 {
		// 	count := len(crowns)
		// 	maxPerPage := 10
		// 	pages := math.Ceil(float64(count) / float64(maxPerPage))
		// 	page := 1

		// 	db, _ := database.OpenDB()
		// 	db.Select(&crowns, db.From(database.Crown{}), db.Limit(maxPerPage), db.Offset(0))
		// 	for _, crown := range crowns {

		// 	}
		// }

		// dbu := database.GetUser(ctx.Message.Author)

		// sql := db.DB()
		// stmt, _ := sql.Prepare("SELECT * FROM crown LIMIT 1 OFFSET 0")
		// res, _ := stmt.Exec()
		// var desc = ""
		// dbu.DB().Select(&out, dbu.DB().From(database.User{}))

		// ctx.Message.Reply(ctx.Aurora, dbu.Crowns())
		// fmt.Println(out)

	}

	return c.Command.CommandInterface
}
