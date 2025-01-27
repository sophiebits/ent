package data

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //driver not used
	"github.com/lolopinto/ent/ent/config"
)

var db *sqlx.DB
var dbMutex sync.RWMutex

// init() initializes the database connection pool for use later
// init function called as package is initalized. Maybe make this explicit with InitDB()?
func init() {
	connStr := config.GetConnectionStr()
	var err error
	db, err = sqlx.Open("postgres", connStr)
	if err != nil {
		fmt.Println("error opening db", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("DB unreachable", err)
	}
}

// GetSQLAlchemyDatabaseURIgo returns the databause uri needed by sqlalchemy to generate a schema file
func GetSQLAlchemyDatabaseURIgo() string {
	return config.Get().DB.GetSQLAlchemyDatabaseURIgo()
}

// DBConn returns a database connection pool to the DB for use
func DBConn() *sqlx.DB {
	dbMutex.RLock()
	defer dbMutex.RUnlock()
	return db
}

// CloseDB closes the database connection pool
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// TODO this obviously needs to be cleaned up
// used by tests
func ResetDB(db2 *sqlx.DB, rdbi *config.DBConfig) error {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if db != nil {
		if err := db.Close(); err != nil {
			return err
		}
	}
	*db = *db2
	config.ResetConfig(rdbi)
	return nil
}
