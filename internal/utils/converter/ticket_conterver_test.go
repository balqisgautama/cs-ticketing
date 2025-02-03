package utilsconverter_test

import (
	"cs-ticketing/internal/queries"
	utilsconverter "cs-ticketing/internal/utils/converter"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UtilsConverterTiket", func() {
	var (
		utilsConverter *utilsconverter.UtilsConverterTiket
	)

	BeforeEach(func() {
		ticketStatusQueries := &queries.TicketStatusQueries{}
		utilsConverter = utilsconverter.NewUtilsConverterTicket(ticketStatusQueries)
	})

	Describe("UnixToTime", func() {
		It("should convert Unix timestamp to Jakarta time correctly", func() {
			// Arrange
			unixTimestamp := int64(1633072800) // 2021-10-01 00:00:00 UTC

			// Act
			result := utilsConverter.UnixToTime(unixTimestamp)

			// Assert
			Expect(result).To(Equal("2021-10-01 07:00:00"))
		})

		It("should handle Unix timestamp for a different date", func() {
			// Arrange
			unixTimestamp := int64(1609459200) // 2021-01-01 00:00:00 UTC

			// Act
			result := utilsConverter.UnixToTime(unixTimestamp)

			// Assert
			Expect(result).To(Equal("2021-01-01 07:00:00"))
		})
	})
})
