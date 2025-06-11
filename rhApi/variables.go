package rhapi


const (
	BaseUrl = "https://meurh.segurosicoob.com.br:8400/restmeurh01"
	LoginUrl = BaseUrl + "/auth/login"
	IsFirstLoginUrl = BaseUrl + "/auth/isFirstLogin"
	GetCurrentTime = BaseUrl + "/timesheet/clockingsGeolocation/currentTime/0/0"
	GetPeriodClockings = BaseUrl + "/timesheet/clockings/currentTime/"
	PostClocking = BaseUrl + "/timesheet/clockingsGeolocation/currentTime"
	GetReasons = BaseUrl + "/timesheet/clockingsReasonTypes/currentTime"
)