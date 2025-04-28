package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Data    any
	Message string
	Retcode int
	Game    string
}

func CheckIn(token string, uid string) ([]Response, error) {
	parsedCookie := fmt.Sprintf("ltoken_v2=%s; ltuid_v2=%s", token, uid)

	res, err := makeRequest(parsedCookie)
	if err != nil {
		return nil, fmt.Errorf("error checked in: %w", err)
	}

	return res, nil
}

func makeRequest(cookie string) ([]Response, error) {
	var responses []Response
	client := &http.Client{}

	for _, game := range Games {
		// make body
		body := map[string]string{
			"lang":   "en-us",
			"act_id": game.ActId,
		}

		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling json: %w", err)
		}
		// make request
		httpRequest, err := http.NewRequest(http.MethodPost, game.Endpoint, bytes.NewBuffer(jsonBody))
		if err != nil {
			return nil, fmt.Errorf("error making request: %w", err)
		}

		header := makeHeader(cookie, game.Name)
		httpRequest.Header = header

		httpResponse, err := client.Do(httpRequest)
		if err != nil {
			return nil, fmt.Errorf("error making request: %w", err)
		}

		defer httpResponse.Body.Close()

		// read to bytes
		r, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading to bytes: %w", err)
		}

		var response Response
		// read to json
		err = json.Unmarshal(r, &response)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling to json: %w", err)
		}
		responses = append(responses, Response{
			Data:    response.Data,
			Message: response.Message,
			Retcode: response.Retcode,
			Game:    game.Name,
		})
	}
	return responses, nil
}

func makeHeader(cookie string, game string) http.Header {
	header := http.Header{}
	header.Set("accept", "application/json, text/plain, */*")
	header.Set("accept-encoding", "gzip, deflate, br, zstd")
	header.Set("accept-language", "en-US,en;q=0.9,id;q=0.8,zh-CN;q=0.7,zh;q=0.6")
	header.Set("content-type", "application/json;charset=UTF-8")
	header.Set("origin", "https://act.hoyolab.com")
	header.Set("referer", "https://act.hoyolab.com")
	header.Set("x-rpc-signgame", game)
	header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	header.Set("cookie", cookie)

	return header
}
