package main

import (
	"cs-ticketing/internal/config"
	dbpostgresql "cs-ticketing/internal/db"
	handler "cs-ticketing/internal/handlers"
	"cs-ticketing/internal/queries"
	service "cs-ticketing/internal/services"
	utilsconverter "cs-ticketing/internal/utils/converter"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig("./config", "dev.json")
	if err != nil {
		log.Fatalln("Error loading config", err)
	}

	db, err := dbpostgresql.ConnectDB(
		c.Postgresql.Host,
		c.Postgresql.Port,
		c.Postgresql.User,
		c.Postgresql.Password,
		c.Postgresql.DDName,
	)
	if err != nil {
		log.Fatalln("Error connecting to the database", err)
	}

	ticketQueries := queries.NewTicketQueries(db)
	userQueries := queries.NewUserQueries(db)
	ticketStatusQueries := queries.NewTicketStatusQueries(db)

	utilsConverterTicket := utilsconverter.NewUtilsConverterTicket(ticketStatusQueries)

	ticketService := service.NewTicketService(ticketQueries, userQueries, ticketStatusQueries, utilsConverterTicket)
	ticketHandler := handler.NewTicketHandler(ticketService)

	// Initialize Gin router
	r := gin.Default()

	r.POST("/tickets", ticketHandler.CreateTicketHandler)
	r.GET("/tickets", ticketHandler.GetTickets)

	// Start the server
	port := fmt.Sprintf(":%d", c.Server.Port)
	if err := r.Run(port); err != nil {
		log.Fatalln("Failed to start server", err)
	}
}
