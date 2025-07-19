package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Game struct {
	Name    string
	Acronym string
	Url     string
	ActId   string
}

type GameResponse struct {
	Data    any    `json:"data"`
	Name    string `json:"name"`
	Message string `json:"message"`
	RetCode int    `json:"retcode"`
}

var games []Game = []Game{
	{
		Name:    "Genshin Impact",
		Acronym: "gi",
		Url:     "https://sg-hk4e-api.hoyolab.com/event/sol/sign",
		ActId:   "e202102251931481",
	},
	{
		Name:    "Zenless Zone Zero",
		Acronym: "zzz",
		Url:     "https://sg-act-nap-api.hoyolab.com/event/luna/zzz/os/sign",
		ActId:   "e202406031448091",
	},
	{
		Name:    "Honkai Star Rail",
		Acronym: "hsr",
		Url:     "https://sg-public-api.hoyolab.com/event/luna/os/sign",
		ActId:   "e202303301540311",
	},
}

func Fetch(token string, uid string) ([]GameResponse, error) {
	client := &http.Client{}
	var responses []GameResponse
	for _, game := range games {
		res, err := makeRequest(client, game, token, uid)
		if err != nil {
			return nil, err
		}
		var response GameResponse
		err = json.Unmarshal(res, &response)
		if err != nil {
			return nil, err
		}
		responses = append(responses, GameResponse{
			Name:    game.Name,
			Data:    response.Data,
			Message: response.Message,
			RetCode: response.RetCode,
		})
	}
	return responses, nil
}

func makeRequest(client *http.Client, game Game, token string, uid string) ([]byte, error) {
	body := map[string]string{
		"lang":   "en-us",
		"act_id": game.ActId,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, game.Url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	cookie := fmt.Sprintf("ltoken_v2=%s; ltuid_v2=%s", token, uid)
	req.Header = *makeHeader(cookie, game.Acronym)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}

func makeHeader(cookie string, acronym string) *http.Header {
	header := &http.Header{}
	header.Set("accept", "application/json, text/plain, */*")
	header.Set("accept-encoding", "gzip, deflate, br, zstd")
	header.Set("accept-language", "en-US,en;q=0.9,id;q=0.8,zh-CN;q=0.7,zh;q=0.6")
	header.Set("content-type", "application/json;charset=UTF-8")
	header.Set("origin", "https://act.hoyolab.com")
	header.Set("referer", "https://act.hoyolab.com")
	header.Set("x-rpc-signgame", acronym)
	header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	header.Set("cookie", cookie)
	return header
}
