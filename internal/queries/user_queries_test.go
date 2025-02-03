package queries_test

import (
	"cs-ticketing/internal/config"
	dbpostgresql "cs-ticketing/internal/db"
	"cs-ticketing/internal/queries"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserQueries", func() {
	var (
		userQueries *queries.UserQueries
	)

	BeforeEach(func() {
		// Set up an in-memory SQLite database for testing
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

		userQueries = queries.NewUserQueries(dbConn)
	})

	Context("IsUserExistsByID", func() {
		It("should return true if the user exists", func() {
			// Act
			exists, err := userQueries.IsUserExistsByID(1)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(BeTrue())
		})

		It("should return false if the user does not exist", func() {
			// Act
			exists, err := userQueries.IsUserExistsByID(999)

			// Assert
			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(BeFalse())
		})
	})
})
