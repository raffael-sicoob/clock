package rhapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/fatih/color"
	"github.com/raffael-sicoob/clock/database"
)

func Login(user string, password string) string {

	loginRequest := LoginRequest{
		User:     user,
		Password: password,
	}

	json, err := json.Marshal(loginRequest)
	if err != nil {
		fmt.Println("Error marshalling login request", err)
		return ""
	}

	res, err := http.Post(LoginUrl, "application/json", bytes.NewBuffer(json))
	if err != nil {
		fmt.Println("Error sending request", err)
		return ""
	}

	defer res.Body.Close()

	token := res.Header.Get("Set-Authorization")
	return token
}

func IsLogged(username string, token string) bool {

	params := url.Values{}
	params.Add("employeeId", username)

	url := IsFirstLoginUrl + "?" + params.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request", err)
		return false
	}

	req.Header.Add("Authorization", token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return false
	}

	defer res.Body.Close()

	var isLoggedResponse IsLoggedResponse
	err = json.NewDecoder(res.Body).Decode(&isLoggedResponse)
	if err != nil {
		fmt.Println("Error decoding response", err)
		return false
	}

	return isLoggedResponse.IsLogged
}

func GetValidToken(user database.User, token string) string {

	if IsLogged(user.Username, token) {
		return token
	}

	return Login(user.Username, user.Password)
}

func GetTime(user database.User, token string) ResponseGetTime {
	newToken := GetValidToken(user, token)

	params := url.Values{}
	params.Add("employeeId", user.Username)

	url := GetCurrentTime + "?" + params.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request", err)
		return ResponseGetTime{}
	}

	req.Header.Add("Authorization", newToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return ResponseGetTime{}
	}

	defer res.Body.Close()

	var response ResponseGetTime
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response", err)
		return ResponseGetTime{}
	}

	return response
}

func RequestClocking(user database.User, token string) ResponseGetTime {
	newToken := GetValidToken(user, token)

	currentDateTime := GetTime(user, newToken)

	body := RequestClockingBody{
		Date:      currentDateTime.ActualDate,
		Hour:      currentDateTime.ActualTime,
		Latitude:  "0",
		Longitude: "0",
		Timezone:  0,
		Address:   "l-location-unavailable",
	}

	json, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshalling body", err)
		return ResponseGetTime{}
	}

	req, err := http.NewRequest("POST", PostClocking, bytes.NewBuffer(json))
	if err != nil {
		fmt.Println("Error creating request", err)
		return ResponseGetTime{}
	}

	req.Header.Add("Authorization", newToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return ResponseGetTime{}
	}

	defer res.Body.Close()

	return currentDateTime

}

func GetClockings(user database.User, token string, initPeriod string, endPeriod string) *ResponseGetClockings {
	// validate format of initPeriod and endPeriod

	initPeriodDate, errParseInitPeriod := time.Parse("2006-01-02", initPeriod)

	endPeriodDate, errParseEndPeriod := time.Parse("2006-01-02", endPeriod)

	if errParseInitPeriod != nil || errParseEndPeriod != nil {
		fmt.Println("❌ ", color.RedString("Error parsing period, please use format YYYY-MM-DD"))
		return nil
	}

	if initPeriodDate.After(endPeriodDate) {
		fmt.Println("❌ ", color.RedString("Init period is after end period"))
		return nil
	}

	validToken := GetValidToken(user, token)

	params := url.Values{}
	params.Add("initPeriod", initPeriodDate.Format("2006-01-02"))
	params.Add("endPeriod", endPeriodDate.Format("2006-01-02"))

	url := GetPeriodClockings + "?" + params.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request", err)
		return nil
	}

	req.Header.Add("Authorization", validToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return nil
	}

	defer res.Body.Close()

	var response ResponseGetClockings
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response", err)
		return nil
	}

	return &response
}

func GetBalanceSummary(user database.User, token string, initPeriod string, endPeriod string) *ResponseGetBalanceSummary {

	initPeriodDate, errParseInitPeriod := time.Parse("2006-01-02", initPeriod)

	endPeriodDate, errParseEndPeriod := time.Parse("2006-01-02", endPeriod)

	if errParseInitPeriod != nil || errParseEndPeriod != nil {
		fmt.Println("❌ ", color.RedString("Error parsing period, please use format YYYY-MM-DD"))
		return nil
	}

	if initPeriodDate.After(endPeriodDate) {
		fmt.Println("❌ ", color.RedString("Init period is after end period"))
		return nil
	}

	validToken := GetValidToken(user, token)

	params := url.Values{}
	params.Add("initPeriod", initPeriodDate.Format("2006-01-02"))
	params.Add("endPeriod", endPeriodDate.Format("2006-01-02"))

	url := GetBalanceSummaryUrl + "?" + params.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request", err)
		return nil
	}

	req.Header.Add("Authorization", validToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return nil
	}

	defer res.Body.Close()

	var response ResponseGetBalanceSummary
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response", err)
		return nil
	}

	return &response
}

func EditClockings(user database.User, token string, date string, hour string) error {
	newToken := GetValidToken(user, token)

	dateTime, err := time.Parse("2006-01-02", date)

	if err != nil {
		return fmt.Errorf("❌ %s", color.RedString("Error parsing date, please use format YYYY-MM-DD"))

	}

	hourTime, err := time.ParseDuration(hour)
	if err != nil {
		return fmt.Errorf("❌ %s", color.RedString("Error parsing hour, please use format ex: 1h30m"))

	}

	clocking := EditClockingsBody{
		Date:          dateTime.Format("2006-01-02"),
		Hour:          hourTime.Milliseconds(),
		Origin:        "manual",
		ReferenceDate: dateTime.Format("2006-01-02"),
		Timezone:      180,
	}

	request := RequestEditClockingBody{
		Date:      time.Now().Format("2006-01-02"),
		Justify:   "Alteração manual",
		Reason:    "01",
		Clockings: []EditClockingsBody{clocking},
	}

	json, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("❌ %s", color.RedString("Error marshalling request body"))

	}
	req, err := http.NewRequest("POST", GetPeriodClockings, bytes.NewBuffer(json))
	if err != nil {
		return fmt.Errorf("❌ %s", color.RedString("Error creating request"))

	}

	req.Header.Add("Authorization", newToken)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("❌ %s : %+v", color.RedString("Error sending request"), err)

	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return fmt.Errorf("❌ %s; code: %d", color.RedString("Error sending request"), res.StatusCode)
	}
	return nil
}
