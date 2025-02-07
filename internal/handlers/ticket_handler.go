// internal/handler/ticket_handler.go
package handler

import (
	dtoin "cs-ticketing/internal/dto/in"
	service "cs-ticketing/internal/services"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	ticketService *service.TicketService
}

func NewTicketHandler(ticketService *service.TicketService) *TicketHandler {
	return &TicketHandler{ticketService: ticketService}
}

func (h *TicketHandler) CreateTicketHandler(c *gin.Context) {
	var req dtoin.Ticket
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dtoin.Validate(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.ticketService.CreateTicket(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Ticket created successfully", "ticket": result})
}

func (h *TicketHandler) GetTickets(c *gin.Context) {
	var req dtoin.TicketList

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dtoin.Validate(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := dtoin.ValidateTicketFilter(req.Filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = dtoin.ValidateTicketSort(req.Sort)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validPageSizes := []int{10, 20, 30, 40, 50}
	for _, size := range validPageSizes {
		if req.PageSize <= size+5 {
			req.PageSize = size
			break
		}
	}

	if req.PageSize > validPageSizes[len(validPageSizes)-1] {
		req.PageSize = validPageSizes[len(validPageSizes)-1]
	}

	var startDate, endDate int64
	if req.Filter != nil {
		now := time.Now()
		// Assuming "between" filter value is in the format "start_date,end_date"
		dates := strings.Split(req.Filter.FilterValue, ",")
		if len(dates) > 0 && len(dates) < 2 && req.Filter.FilterType != "between" {
			tempValue, err := time.Parse("2006-01-02", dates[0])
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid date format (example: 2025-01-20)").Error()})
				return
			}

			startDate = tempValue.Unix()
		}

		if len(dates) == 2 && req.Filter.FilterType == "between" {
			tempValue, err := time.Parse("2006-01-02", dates[1])
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("date must be in the format 'start_date,end_date' (example: 2025-01-20,2025-01-25)").Error()})
				return
			}

			tempValue = time.Date(tempValue.Year(), tempValue.Month(), tempValue.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())

			endDate = tempValue.Unix()

			if startDate > endDate {
				c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("start date must be less than end date").Error()})
				return
			} else if startDate == endDate {
				c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("start date must be different from end date").Error()})
				return
			}
		} else if condition := len(dates) == 1 && req.Filter.FilterType == "between"; condition {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("date must be in the format 'start_date,end_date' (example: 2025-01-20,2025-01-25)").Error()})
			return
		}

		if startDate > now.Unix() || endDate > now.Unix() {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("date must not be greater than today").Error()})
			return
		}
	}

	result, err := h.ticketService.GetTickets(req, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
