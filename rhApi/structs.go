package rhapi

type LoginRequest struct {
	User string `json:"user"`
	Password string `json:"password"`
}


type IsLoggedResponse struct {
	IsLogged bool `json:"isLogged"`
}


type RequestClocking struct {
	
}