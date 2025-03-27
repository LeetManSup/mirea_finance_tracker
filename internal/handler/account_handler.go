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

type updateAccountRequest struct {
	Name           *string  `json:"name"`
	CurrencyCode   *string  `json:"currency_code"`
	InitialBalance *float64 `json:"initial_balance"`
}

// CreateAccount создаёт новый финансовый счёт
// @Summary Создание счёта
// @Tags accounts
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body CreateAccountInput true "Данные счёта"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /accounts [post]
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

// GetAccounts возвращает список счетов пользователя
// @Summary Получить все счета
// @Tags accounts
// @Security BearerAuth
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /accounts [get]
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

// GetAccount возвращает счёт по ID
// @Summary Получить счёт по ID
// @Tags accounts
// @Security BearerAuth
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} map[string]interface{}
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /accounts/{id} [get]
func (h *AccountHandler) GetAccount(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	accountID := c.Param("id")
	account, err := h.accountService.GetAccountByID(userID.(string), accountID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":              account.ID,
		"name":            account.Name,
		"currency_code":   account.CurrencyCode,
		"initial_balance": account.InitialBalance,
		"created_at":      account.CreatedAt,
	})
}

// DeleteAccount удаляет счёт
// @Summary Удалить счёт
// @Tags accounts
// @Security BearerAuth
// @Param id path string true "Account ID"
// @Success 204
// @Failure 403 {object} map[string]string
// @Router /accounts/{id} [delete]
func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	accountID := c.Param("id")
	err := h.accountService.DeleteAccount(userID.(string), accountID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateAccount обновляет данные счёта
// @Summary Обновить счёт
// @Tags accounts
// @Security BearerAuth
// @Accept json
// @Param id path string true "Account ID"
// @Param input body updateAccountRequest true "Поля для обновления"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /accounts/{id} [patch]
func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input updateAccountRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountID := c.Param("id")

	dto := service.UpdateAccountInput{
		Name:           input.Name,
		CurrencyCode:   input.CurrencyCode,
		InitialBalance: input.InitialBalance,
	}

	err := h.accountService.UpdateAccount(userID.(string), accountID, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
