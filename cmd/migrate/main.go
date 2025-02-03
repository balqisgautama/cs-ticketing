package main

import (
	"cs-ticketing/internal/config"
	pkgdbpostgresql "cs-ticketing/pkg/db/migrations"
	"log"
)

func main() {
	c, err := config.LoadConfig("./internal/config", "migrate.json")
	if err != nil {
		log.Fatalln("Error loading config", err)
	}

	dbMigration, err := pkgdbpostgresql.GetDbConnection(
		c.Postgresql.Schema,
		c.Postgresql.Address,
		c.Postgresql.MaxOpenConnection,
		c.Postgresql.MaxIdleConnection,
	)
	if err != nil {
		log.Fatalln("Error connecting to the database for migration", err)
	}

	err = pkgdbpostgresql.DBMigration(dbMigration)
	if err != nil {
		log.Fatalln("Error migrating the database", err)
	}

	err = dbMigration.Close()
	if err != nil {
		log.Fatalln("Error closing the database connection", err)
	}
}
