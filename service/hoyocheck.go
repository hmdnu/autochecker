package service

import (
	"fmt"
	"log"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/jasonlvhit/gocron"
)

func CheckIn() (map[string][]Response, error) {
	responses := make(map[string][]Response)

	config, err := GetConfig()
	if err != nil {
		return nil, fmt.Errorf("error getting config: %w", err)
	}

	for _, v := range config.User {
		parsedCookie := fmt.Sprintf("ltoken_v2=%s; ltuid_v2=%s", v.Token, v.Uid)
		res, err := MakeRequest(parsedCookie)
		if err != nil {
			return nil, fmt.Errorf("error checked in: %w", err)
		}

		responses[v.Name] = res
	}

	return responses, nil
}

func AutoCheck() {
	config, err := GetConfig()
	if err != nil {
		log.Fatalln("error getting config time: %w", err)
	}

	// init scheduler
	err = gocron.Every(1).Day().At(config.CheckTime).From(gocron.NextTick()).Do(func() {
		res, err := CheckIn()
		if err != nil {
			log.Fatalln("error checked in: %w", err)
			beepNotification("error checked in")
		}
		beepNotification(displayMessage(res))
		WriteLog(res)
	})

	// start gocron task
	go func() {
		<-gocron.Start()
	}()

	if err != nil {
		log.Fatalln("error checked in: %w", err)
		beepNotification("error checked in")
	}
}

func beepNotification(message string) {
	err := beeep.Notify("Hoyocheck", message, "")
	if err != nil {
		panic(err)
	}
}

func displayMessage(res map[string][]Response) string {
	msg := ""
	keys := make([]string, 0, len(res))
	for k := range res {
		keys = append(keys, k)
	}
	for i, k := range keys {
		msg += k + ":\n"
		for _, vi := range res[k] {
			msg += fmt.Sprintf("%s: %s\n", vi.Game, Status[vi.Retcode])
		}
		if i != len(keys)-1 {
			msg += "\n"
		}
	}
	return msg
}

func WriteLog(res map[string][]Response) {
	var logs []HistoryLog

	for user, responses := range res {
		log := HistoryLog{
			Timestamp: time.Now().Format("02-01-2006 15:04"),
			User:      user,
		}

		for _, r := range responses {
			switch r.Game {
			case "gi":
				log.Gi = Status[r.Retcode]
			case "hsr":
				log.Hsr = Status[r.Retcode]
			case "zzz":
				log.Zzz = Status[r.Retcode]
			}
		}

		logs = append(logs, log)
	}

	for _, l := range logs {
		Logger(l)
	}
}
