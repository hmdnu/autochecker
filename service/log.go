package service

import (
	"encoding/json"
	"fmt"
	"os"
)

type HistoryLog struct {
	Timestamp string `json:"timestamp"`
	User      string `json:"user"`
	Gi        string `json:"genshin-impact"`
	Hsr       string `json:"honkai-star-rail"`
	Zzz       string `json:"zenless-zone-zero"`
}

func Logger(log HistoryLog) error {
	file, err := os.OpenFile("history.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	// write as json
	encoder := json.NewEncoder(file)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", " ")
	err = encoder.Encode(log)
	if err != nil {
		return fmt.Errorf("error endoing to json: %w", err)
	}
	return nil
}
