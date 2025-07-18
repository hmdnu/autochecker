package account

import (
	"encoding/json"
	"myapp/helper"
	"os"
)

type Account struct {
	Id    string
	Name  string
	Token string
	Uid   string
}

const ACCOUNT_DIR = "data/account.json"

func ReadAccount() ([]Account, error) {
	byte, err := os.ReadFile(ACCOUNT_DIR)
	if err != nil {
		return nil, err
	}
	var accounts []Account
	if err := json.Unmarshal(byte, &accounts); err != nil {
		return nil, err
	}
	return accounts, nil
}

func AddAccount(account Account) error {
	if err := helper.EnsureFileExist(ACCOUNT_DIR); err != nil {
		return err
	}
	accounts, err := ReadAccount()
	if err != nil {
		return err
	}
	accounts = append(accounts, account)
	if err := helper.WriteIntoJson(ACCOUNT_DIR, accounts); err != nil {
		return err
	}
	return nil
}

func DeleteAccount(id string) error {
	accounts, err := ReadAccount()
	if err != nil {
		return err
	}
	indexToDelete := -1
	for i, v := range accounts {
		if v.Id == id {
			indexToDelete = i
			break
		}
	}
	accounts = append(accounts[:indexToDelete], accounts[indexToDelete+1:]...)
	err = helper.WriteIntoJson(string(ACCOUNT_DIR), accounts)
	return nil
}
