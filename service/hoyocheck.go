package service

import (
	"encoding/json"
	"myapp/handler"
	"os"
	"time"
)

type Response struct {
	Id          string
	AccountName string
	DateCheck   string
	Game        []handler.GameResponse
}

const LOG_DIR = "data/log.json"

func HandleCheck() ([]Response, error) {
	accounts, err := ReadAccount()
	if err != nil {
		return nil, err
	}
	existingLogs, err := ReadLogs()
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		gameResponses, err := handler.Fetch(account.Token, account.Uid)
		if err != nil {
			return nil, err
		}
		existingLogs = append(existingLogs, Response{
			Id:          account.Id,
			AccountName: account.Name,
			DateCheck:   time.Now().Format(time.RFC1123),
			Game:        gameResponses,
		})
	}
	if err := writeTofile(existingLogs); err != nil {
		return nil, err
	}
	return existingLogs, nil
}

func writeTofile(responses []Response) error {
	if err := handler.EnsureFileExist(LOG_DIR); err != nil {
		return err
	}
	if err := handler.WriteIntoJson(LOG_DIR, responses); err != nil {
		return err
	}
	return nil
}

func ReadLogs() ([]Response, error) {
	byte, err := os.ReadFile(LOG_DIR)
	if err != nil {
		return nil, err
	}
	var logs []Response
	if err := json.Unmarshal(byte, &logs); err != nil {
		return nil, err
	}
	return logs, nil
}
