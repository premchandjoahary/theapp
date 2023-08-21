package http

import (
	"encoding/json"
	"net/http"
	"theapp/models"
	"time"

	"github.com/labstack/echo/v4"
)

// LoanHandlerHTTP represent the http handler
type LoanHandlerHTTP struct{}

var loans []models.Loan

// ConfigureLoanHTTP configures the handler functions
func ConfigureLoanHTTP(e *echo.Echo) {
	handler := &LoanHandlerHTTP{}
	handler.AddHandlers(e)
}

func (hndHttp *LoanHandlerHTTP) AddHandlers(e *echo.Echo) {
	e.POST("/loans", hndHttp.CreateLoan)
	e.POST("/loans/approve", hndHttp.ApproveLoan)
	e.POST("/loans/customer/{customer_id}", hndHttp.GetLoansByCustomer)
	e.POST("/repayments", hndHttp.AddRepayment)

}

func (hndHttp *LoanHandlerHTTP) CreateLoan(c echo.Context) error {
	var loan models.Loan
	_ = json.NewDecoder(c.Request().Body).Decode(&loan)

	// Create scheduled repayments
	repayments := hndHttp.generateRepayments(loan)

	loan.Repayments = repayments
	loan.Status = "PENDING"
	loans = append(loans, loan)

	return c.JSON(http.StatusCreated, models.APIResponseMessage{
		Status: "SUCCESS",
		data:   loans,
	})
}

func (hndHttp *LoanHandlerHTTP) generateRepayments(loan models.Loan) []models.Repayment {
	repayments := make([]models.Repayment, loan.Term)
	startDate := time.Now().AddDate(0, 0, 7) // Starting from 7 days from now

	for i := 0; i < loan.Term; i++ {
		repayments[i] = models.Repayment{
			ID:      i + 1,
			LoanID:  loan.ID,
			Amount:  loan.Amount / float64(loan.Term),
			DueDate: startDate.AddDate(0, 0, 7*i),
			Status:  "PENDING",
		}
	}

	return repayments
}

func (hndHttp *LoanHandlerHTTP) ApproveLoan(c echo.Context) error {
	// Here we can write logic to change loan status to APPROVED. As I am not yet configure database connection so just knowing you my idea of implementation
}

func (hndHttp *LoanHandlerHTTP) GetLoansByCustomer(c echo.Context) error {
	// Here we can write logic to get loans by customer ID. Again need DB connection.
}

func (hndHttp *LoanHandlerHTTP) AddRepayment(c echo.Context) error {
	// Here we can write logic to add a repayment and update repayment status. As DB connection will again need some more time to implement.
}
