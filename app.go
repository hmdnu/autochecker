package main

import (
	"context"
	"fmt"
	helper "myapp/handler"
	"myapp/service"
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

	err := helper.EnsureFileExist(service.ACCOUNT_DIR)
	if err != nil {
		println(err)
	}
}

func (a *App) HandleAddAccount(user service.Account) error {
	err := service.AddAccount(user)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}

func (a *App) HandleReadAccount() ([]service.Account, error) {
	accounts, err := service.ReadAccount()
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	return accounts, nil
}

func (a *App) HandleDeleteAccount(id string) error {
	err := service.DeleteAccount(id)
	if err != nil {
		return fmt.Errorf("%e", err)
	}
	return nil
}

func (a *App) HandleCheck() error {
	resp, err := service.HandleCheck()
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}
