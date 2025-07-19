package service

import "myapp/handler"

type Response struct {
	Id          string
	AccountName string
	Game        []handler.GameResponse
}

func HandleCheck() ([]Response, error) {
	var responses []Response
	accounts, err := ReadAccount()
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		res, err := handler.Fetch(account.Token, account.Uid)
		if err != nil {
			return nil, err
		}
		responses = append(responses, Response{
			Id:          account.Id,
			AccountName: account.Name,
			Game:        res,
		})
	}
	return responses, nil
}

func writeTofile() {

}
