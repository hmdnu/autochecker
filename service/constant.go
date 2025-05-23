package service

type Game struct {
	Name     string
	Endpoint string
	ActId    string
}

var Games = []Game{
	{
		Name:     "gi",
		Endpoint: "https://sg-hk4e-api.hoyolab.com/event/sol/sign",
		ActId:    "e202102251931481",
	},
	{
		Name:     "zzz",
		Endpoint: "https://sg-act-nap-api.hoyolab.com/event/luna/zzz/os/sign",
		ActId:    "e202406031448091",
	},
	{
		Name:     "hsr",
		Endpoint: "https://sg-public-api.hoyolab.com/event/luna/os/sign",
		ActId:    "e202303301540311",
	},
}

var Status = map[int]string{
	0:      "check in success",
	-5003:  "already checked in for today",
	-100:   "cookie invalid",
	-10002: "you have not played this game yet",
}
