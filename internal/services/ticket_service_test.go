package service_test

import (
	"cs-ticketing/internal/config"
	dbpostgresql "cs-ticketing/internal/db"
	dtoin "cs-ticketing/internal/dto/in"
	"cs-ticketing/internal/queries"
	service "cs-ticketing/internal/services"
	utilsconverter "cs-ticketing/internal/utils/converter"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TicketService", func() {
	var (
		ticketService       *service.TicketService
		ticketQueries       *queries.TicketQueries
		userQueries         *queries.UserQueries
		ticketStatusQueries *queries.TicketStatusQueries
		utilsConverter      *utilsconverter.UtilsConverterTiket
	)

	BeforeEach(func() {
		// Set up the necessary components for TicketService
		c, err := config.LoadConfig("../../internal/config", "test.json")
		Expect(err).NotTo(HaveOccurred())

		dbConn, err := dbpostgresql.ConnectDB(
			c.Postgresql.Host,
			c.Postgresql.Port,
			c.Postgresql.User,
			c.Postgresql.Password,
			c.Postgresql.DDName,
		)
		Expect(err).NotTo(HaveOccurred())

		ticketQueries = queries.NewTicketQueries(dbConn)
		userQueries = queries.NewUserQueries(dbConn)
		ticketStatusQueries = queries.NewTicketStatusQueries(dbConn)
		utilsConverter = utilsconverter.NewUtilsConverterTicket(ticketStatusQueries)

		ticketService = service.NewTicketService(ticketQueries, userQueries, ticketStatusQueries, utilsConverter)
	})

	Context("GetTickets", func() {
		It("should return tickets within the specified date range", func() {
			// Arrange
			req := dtoin.TicketList{
				PageSize: 10,
				Page:     1,
			}
			startDate := time.Now().Add(-24 * time.Hour).Unix()
			endDate := time.Now().Unix()

			// Act
			tickets, err := ticketService.GetTickets(req, startDate, endDate)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(tickets).NotTo(BeNil())
			Expect(len(tickets.Tickets)).To(BeNumerically(">", 0))
		})

		It("should return no tickets if none exist within the specified date range", func() {
			// Arrange
			req := dtoin.TicketList{
				PageSize: 10,
				Page:     1,
			}
			startDate := time.Now().Add(-48 * time.Hour).Unix()
			endDate := time.Now().Add(-24 * time.Hour).Unix()

			// Act
			tickets, err := ticketService.GetTickets(req, startDate, endDate)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(tickets).NotTo(BeNil())
			Expect(len(tickets.Tickets)).To(Equal(0))
		})

		It("should return an error if the query fails", func() {
			// Arrange
			req := dtoin.TicketList{
				PageSize: 10,
				Page:     1,
			}
			startDate := time.Now().Add(-24 * time.Hour).Unix()
			endDate := time.Now().Unix()

			// Simulate a query failure
			ticketQueries = queries.NewTicketQueries(nil)
			ticketService = service.NewTicketService(ticketQueries, userQueries, nil, utilsConverter)

			// Act
			tickets, err := ticketService.GetTickets(req, startDate, endDate)

			// Assert
			Expect(err).To(HaveOccurred())
			Expect(tickets).To(BeNil())
		})
	})
})
