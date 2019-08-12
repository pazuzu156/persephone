package database

import (
	"fmt"
	"persephone/lib"

	_ "github.com/go-sql-driver/mysql" // required for mysql driver
	"github.com/naoina/genmai"
)

var config = lib.Config()

// Migrate migrates database tables (only run on database creation)
func Migrate() {
	var db *genmai.DB
	db, _ = OpenDB()
	db.CreateTableIfNotExists(&User{})
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
}
