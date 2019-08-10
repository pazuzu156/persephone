package utils

import (
	"fmt"
	"persephone/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/naoina/genmai"
)

var config = Config()

// Migrate migrates database tables (only run on database creation)
func Migrate() {
	var db *genmai.DB
	db, _ = OpenDB()

	db.CreateTableIfNotExists(&models.User{})

	db.Close()
}

// OpenDB opens a database connection
func OpenDB() (*genmai.DB, error) {
	db, err := genmai.New(&genmai.MySQLDialect{}, fmt.Sprintf("%s:%s@(%s:%d)/%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Hostname,
		config.Database.Port,
		config.Database.Name,
	))

	return db, err
	// db, err := genmai.New(&genmai.MySQLDialect{}, "root@/persephone")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// obj := []User{
	// 	{Username: "Apollyon", DiscordID: 12345678, Lastfm: "Pazuzu156"},
	// }
	// n, err := db.Insert(obj)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Inserted rows: %d\n", n)
}
