package pkgdbpostgresql

import (
	"database/sql"
	"log"
	"strconv"
	"sync"

	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

type DBInfo struct {
	instance      *sql.DB
	driver        string
	connectionStr string
	setParams     []string
}

var instance *sql.DB
var once sync.Once

func GetDbConnection(defaultSchema string, connectionString string, maxOpenConnection int, maxIdleConnection int) (*sql.DB, error) {
	log.Println("SCHEMA", defaultSchema)
	_dbInfo := DBInfo{nil, "pgx",
		connectionString, []string{"search_path = '" + defaultSchema + "'"}}
	_db, _err := getInstance(_dbInfo)
	if _err != nil {
		return nil, _err
	}
	_db.SetMaxOpenConns(maxOpenConnection)
	_db.SetMaxIdleConns(maxIdleConnection)
	return _db, nil
}

func getInstance(connInfo DBInfo) (*sql.DB, error) {
	var _errOpen error
	once.Do(func() {
		dbConnStr := connInfo.connectionStr
		if len(connInfo.setParams) > 0 {
			for _, _param := range connInfo.setParams {
				dbConnStr = dbConnStr + " " + _param
			}
		}
		instance, _errOpen = sql.Open(connInfo.driver, dbConnStr)

		if _errOpen != nil {
			log.Println("Error in opening connection", _errOpen.Error())
			instance = nil
		}
	})
	return instance, _errOpen
}

func DBMigration(db *sql.DB) error {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	log.Println("Applied " + strconv.Itoa(n) + " migrations!")
	return nil
}
