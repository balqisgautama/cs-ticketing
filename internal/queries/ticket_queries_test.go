package queries_test

import (
	"cs-ticketing/internal/config"
	dbpostgresql "cs-ticketing/internal/db"
	modeldb "cs-ticketing/internal/models/db"
	"cs-ticketing/internal/queries"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("TicketQueries", func() {
	var (
		ticketQueries *queries.TicketQueries
		dbConn        *gorm.DB
	)

	BeforeEach(func() {
		// Set up an in-memory SQLite database for testing
		c, err := config.LoadConfig("../../internal/config", "test.json")
		Expect(err).NotTo(HaveOccurred())

		dbConn, err = dbpostgresql.ConnectDB(
			c.Postgresql.Host,
			c.Postgresql.Port,
			c.Postgresql.User,
			c.Postgresql.Password,
			c.Postgresql.DDName,
		)
		Expect(err).NotTo(HaveOccurred())

		ticketQueries = queries.NewTicketQueries(dbConn)
	})

	AfterEach(func() {
		dbConn.Exec("DELETE FROM tickets")
	})

	Context("CreateTicket", func() {
		It("should create a ticket successfully", func() {
			// Arrange
			ticket := &modeldb.Ticket{
				UserID:    1,
				Title:     "Software is not working",
				Msg:       "<p>This is a valid ticket message with more than 100 characters. It should pass the validation checks.</p>",
				CreatedAt: time.Now().Unix(),
				StatusID:  1,
			}

			// Act
			createdTicket, err := ticketQueries.CreateTicket(ticket)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(createdTicket).NotTo(BeNil())
			Expect(createdTicket.ID).NotTo(BeZero())
			Expect(createdTicket.Title).To(Equal("Software is not working"))
		})

		It("should return an error if ticket creation fails", func() {
			// Arrange
			// Create a ticket with invalid data (e.g., missing required fields)
			ticket := &modeldb.Ticket{
				// Intentionally leaving out required fields
			}

			// Act
			createdTicket, err := ticketQueries.CreateTicket(ticket)

			// Assert
			Expect(err).To(HaveOccurred())
			Expect(createdTicket).To(BeNil())
		})
	})

	Context("GetTickets", func() {
		BeforeEach(func() {
			// Seed the database with some tickets
			for i := 0; i < 5; i++ {
				ticket := &modeldb.Ticket{
					UserID:    1,
					Title:     "Software is not working",
					Msg:       "<p>This is a valid ticket message with more than 100 characters. It should pass the validation checks.</p>",
					CreatedAt: time.Now().Unix(),
					StatusID:  1,
				}
				Expect(dbConn.Create(ticket).Error).NotTo(HaveOccurred())
			}
		})

		It("should return all tickets with pagination", func() {
			// Act
			tickets, total, err := ticketQueries.GetTickets(&modeldb.TicketFilter{}, &modeldb.TicketSort{}, 2, 1)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(total).To(Equal(int64(5)))
			Expect(tickets).To(HaveLen(2))
		})

		It("should filter tickets by created_at before a certain date", func() {
			// Arrange
			tempValue, err := time.Parse("2006-01-02", "2025-02-03")
			Expect(err).NotTo(HaveOccurred())
			filter := &modeldb.TicketFilter{
				FilterType:  "before",
				FilterValue: tempValue.Unix(),
			}

			// Act
			_, _, err = ticketQueries.GetTickets(filter, &modeldb.TicketSort{}, 10, 1)

			// Assert
			Expect(err).NotTo(HaveOccurred())
		})

		It("should sort tickets by created_at in descending order", func() {
			// Arrange
			sort := &modeldb.TicketSort{
				SortName: "created_at",
				SortDir:  "desc",
			}

			// Act
			_, _, err := ticketQueries.GetTickets(&modeldb.TicketFilter{}, sort, 10, 1)

			// Assert
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
