package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"theapp/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLoanHandler_CreateLoan(t *testing.T) {
	e := echo.New()
	reqBody := []byte(`{"amount": 10000, "term": 3}`) // Example request body

	req := httptest.NewRequest(http.MethodPost, "/loans", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := &LoanHandlerHTTP{}
	err := h.CreateLoan(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response models.APIResponseMessage
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "SUCCESS", response.Status)
	// You can add more assertions based on your response structure.
}
