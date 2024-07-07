package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ChangeAccountRequest struct {
	Name    string `json:"name"`
	Amount  int    `json:"amount"`
	NewName string `json:"new_name,omitempty"`
	To      string `json:"to,omitempty"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}
