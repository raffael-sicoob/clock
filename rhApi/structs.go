package rhapi

type LoginRequest struct {
	User string `json:"user"`
	Password string `json:"password"`
}


type IsLoggedResponse struct {
	IsLogged bool `json:"isLogged"`
}


type RequestClockingBody struct {
	Date string `json:"date"`
	Hour int16 `json:"hour"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	Timezone int8 `json:"timezone"`
	Address string `json:"address"`
}

type ResponseGetTime struct {
	ActualDate string `json:"actualDate"`
	ActualTime int16 `json:"actualTime"`
}