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
	Hour uint32 `json:"hour"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	Timezone uint8 `json:"timezone"`
	Address string `json:"address"`
}

type ResponseGetTime struct {
	ActualDate string `json:"actualDate"`
	ActualTime uint32 `json:"actualTime"`
}

// Structs para resposta de GetClockings
// ClockingStatus representa o status de aprovação do ponto
type ClockingStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Label  string `json:"label"`
}

// Clocking representa um registro de ponto individual
// Campos principais necessários: Date, Hour, Direction
type Clocking struct {
	ID        string         `json:"id"`
	Date      string         `json:"date"`      // Data do ponto (campo necessário)
	Hour      uint32         `json:"hour"`      // Horário em segundos (campo necessário) 
	Direction string         `json:"direction"` // "entry" ou "exit" (campo necessário)
	Origin    string         `json:"origin"`
	Status    ClockingStatus `json:"status"`
	Sequence  int            `json:"sequence"`
}

// ResponseGetClockings é a resposta completa da API de consulta de pontos
type ResponseGetClockings struct {
	InitPeriod string     `json:"initPeriod"`
	EndPeriod  string     `json:"endPeriod"`
	Clockings  []Clocking `json:"clockings"`
}

type ResponseGetBalanceSummary struct {
	Previous int32 `json:"previous"`
	Current int32 `json:"current"`
	Next int32 `json:"next"`
}