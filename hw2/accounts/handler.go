package accounts

import (
	dto2 "awesomeProject/hw2/accounts/dto"
	"awesomeProject/hw2/accounts/models"
	"net/http"
	"sync"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto2.ChangeAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	response := dto2.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteAccount(c echo.Context) error {
	var request dto2.DeleteAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	if _, ok := h.accounts[request.Name]; !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	delete(h.accounts, request.Name)
	return c.NoContent(http.StatusOK)
}

func (h *Handler) ChangeAmount(c echo.Context) error {
	var request dto2.ChangeAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	account, ok := h.accounts[request.Name]
	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	if request.Amount+account.Amount < 0 {
		return c.String(http.StatusBadRequest, "insufficient funds for withdrawal")
	}

	account.Amount += request.Amount
	return c.NoContent(http.StatusOK)
}

func (h *Handler) ChangeAccount(c echo.Context) error {
	var request dto2.ChangeAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 || len(request.NewName) == 0 {
		return c.String(http.StatusBadRequest, "empty name or new name")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	if _, ok := h.accounts[request.Name]; !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	if _, ok := h.accounts[request.NewName]; ok {
		return c.String(http.StatusForbidden, "new account name already exists")
	}

	h.accounts[request.NewName] = h.accounts[request.Name]
	h.accounts[request.NewName].Name = request.NewName
	delete(h.accounts, request.Name)

	return c.NoContent(http.StatusOK)
}

func (h *Handler) MoneyTransaction(c echo.Context) error {
	var request dto2.ChangeAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	if request.Amount < 0 {
		return c.String(http.StatusBadRequest, "incorrect transfer amount")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	accountFrom, ok := h.accounts[request.Name]
	if !ok {
		return c.String(http.StatusNotFound, "sender account not found")
	}
	accountTo, ok := h.accounts[request.To]
	if !ok {
		return c.String(http.StatusNotFound, "recipient account not found")
	}

	if accountFrom.Amount-request.Amount < 0 {
		return c.String(http.StatusBadRequest, "insufficient funds for withdrawal")
	}

	accountFrom.Amount -= request.Amount
	accountTo.Amount += request.Amount

	return c.NoContent(http.StatusOK)
}
