package main

import (
	"awesomeProject/accounts/dto"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	NewName string
	To      string
	Amount  int
}

func main() {
	portVal := flag.Int("port", 1323, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")
	newNameVal := flag.String("newName", "", "new name of account")
	toVal := flag.String("to", "", "to account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		Amount:  *amountVal,
		NewName: *newNameVal,
		To:      *toVal,
	}

	if err := do(cmd); err != nil {
		panic(err)
	}
}

func do(cmd Command) error {
	switch cmd.Cmd {
	case "create":
		if err := create(cmd); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}
		return nil
	case "get":
		if err := get(cmd); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}
		return nil
	case "delete":
		if err := delete(cmd); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}
		return nil
	case "changeName":
		if err := changeName(cmd); err != nil {
			return fmt.Errorf("change name failed: %w", err)
		}
		return nil
	case "replenishment":
		if err := replenishment(cmd); err != nil {
			return fmt.Errorf("replenishment of balance failed: %w", err)
		}
		return nil
	case "withdrawal":
		if err := withdrawal(cmd); err != nil {
			return fmt.Errorf("withdrawal from balance failed: %w", err)
		}
		return nil
	case "transaction":
		if err := transaction(cmd); err != nil {
			return fmt.Errorf("money transaction failed: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func get(cmd Command) error {
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name))
	if err != nil {
		return fmt.Errorf("http get failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}
		return fmt.Errorf("resp error %s", string(body))
	}

	var response dto.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("response account name: %s and amount: %d\n", response.Name, response.Amount)
	return nil
}

func create(cmd Command) error {
	request := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func delete(cmd Command) error {
	request := dto.DeleteAccountRequest{
		Name: cmd.Name,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/delete", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func changeName(cmd Command) error {
	request := dto.ChangeAccountRequest{
		Name:    cmd.Name,
		NewName: cmd.NewName,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/name", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func replenishment(cmd Command) error {
	request := dto.ChangeAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/amount", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func withdrawal(cmd Command) error {
	request := dto.ChangeAccountRequest{
		Name:   cmd.Name,
		Amount: -cmd.Amount,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/amount", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func transaction(cmd Command) error {
	request := dto.ChangeAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
		To:     cmd.To,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/transaction", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}
