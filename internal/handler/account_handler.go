package handler

import (
	"net/http"

	"mirea_finance_tracker/internal/service"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService}
}

type CreateAccountInput struct {
	Name           string  `json:"name" binding:"required"`
	CurrencyCode   string  `json:"currency_code" binding:"required,len=3"`
	InitialBalance float64 `json:"initial_balance"`
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var input CreateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	accountID, err := h.accountService.CreateAccount(userID.(string), input.Name, input.CurrencyCode, input.InitialBalance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"account_id": accountID})
}

func (h *AccountHandler) GetAccounts(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	accounts, err := h.accountService.GetAccountsByUser(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load accounts"})
		return
	}

	response := make([]gin.H, 0, len(accounts))
	for _, acc := range accounts {
		response = append(response, gin.H{
			"id":              acc.ID,
			"name":            acc.Name,
			"currency_code":   acc.CurrencyCode,
			"initial_balance": acc.InitialBalance,
			"created_at":      acc.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}
