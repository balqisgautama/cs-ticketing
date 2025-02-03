package queries_test

import (
	"cs-ticketing/internal/config"
	dbpostgresql "cs-ticketing/internal/db"
	"cs-ticketing/internal/queries"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TicketStatusQueries", func() {
	var (
		ticketStatusQueries *queries.TicketStatusQueries
	)

	BeforeEach(func() {
		// Set up the database connection
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

		ticketStatusQueries = queries.NewTicketStatusQueries(dbConn)
	})

	Context("FindTicketStatusByID", func() {
		It("should return the ticket status if it exists", func() {
			// Act
			result, err := ticketStatusQueries.FindTicketStatusByID(1)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(result).NotTo(BeNil())
			Expect(result.ID).To(Equal(int64(1)))
			Expect(result.Name).To(Equal("opn"))
		})
	})
})
