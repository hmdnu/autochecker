package main

import (
	"context"
	"fmt"
	"myapp/account"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) HandleAddAccount(user account.Account) error {
	err := account.AddAccount(user)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

func (a *App) HandleReadAccount() ([]account.Account, error) {
	accounts, err := account.ReadAccount()
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	return accounts, nil
}

func (a *App) HandleDeleteAccount(id string) error {
	err := account.DeleteAccount(id)
	if err != nil {
		return fmt.Errorf("%e", err)
	}
	return nil
}
