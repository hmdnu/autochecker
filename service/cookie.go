package service

import (
	"encoding/json"
	"fmt"
	"os"
)

type Cookie struct {
	Token string
	Uid   string
}

func GetCookie() (Cookie, error) {
	var cookie Cookie

	file, err := os.Open("data.json")
	if err != nil {
		return cookie, fmt.Errorf("error opening data.json : %w", err)
	}

	err = json.NewDecoder(file).Decode(&cookie)
	if err != nil {
		return cookie, err
	}

	return cookie, nil
}

// func GetCookie() (Cookie, error) {
// 	file, err := os.Open("cookie.txt")

// 	if err != nil {
// 		return Cookie{}, fmt.Errorf("failed to open file : %w", err)
// 	}

// 	scanner := bufio.NewScanner(file)
// 	var cookie Cookie

// 	for scanner.Scan() {
// 		line := strings.Split(scanner.Text(), "=")
// 		key := line[0]
// 		value := line[1]

// 		if key == "token" {
// 			cookie.Token = value
// 		}
// 		if key == "uid" {
// 			cookie.Uid = value
// 		}
// 	}

// 	return cookie, nil
// }
