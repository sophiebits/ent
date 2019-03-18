package data

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" //driver not used
)

// all of this needs to be taken from a configuration file instead of hardcoded
const (
	host = "localhost"
	port = 5432
	user = "ola"
	//  password = ""
	dbname = "jarvis"
)

var db *sqlx.DB

// init() initializes the database connection pool for use later
// init function called as package is initalized. Maybe make this explicit with InitDB()?
func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	var err error
	db, err = sqlx.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error opening db", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("DB unreachable", err)
	}
	//fmt.Println("InitDB", db)
}

// DBConn returns a database connection pool to the DB for use
func DBConn() *sqlx.DB {
	return db
}

// CloseDB closes the database connection pool
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
