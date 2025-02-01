package main

import (
	"cs-ticketing/internal/config"
	pkgdbpostgresql "cs-ticketing/pkg/db/migrations"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
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
		log.Fatalln("Error connecting to the database", err)
	}

	err = pkgdbpostgresql.DBMigration(dbMigration)
	if err != nil {
		log.Fatalln("Error migrating the database", err)
	}

	err = dbMigration.Close()
	if err != nil {
		log.Fatalln("Error closing the database connection", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Start the server
	port := fmt.Sprintf(":%d", c.Server.Port)
	if err := r.Run(port); err != nil {
		log.Fatalln("Failed to start server", err)
	}
}
