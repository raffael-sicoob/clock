package rhapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/raffael-sicoob/clock/database"
)

func Login(user string, password string) string{
	

	loginRequest := LoginRequest{
		User: user,
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


func IsLogged(username string, token string) bool{


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

func GetValidToken(user database.User, token string) string{

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


func RequestClocking( user database.User, token string) ResponseGetTime  {
	newToken := GetValidToken(user, token)

	currentDateTime := GetTime(user, newToken)

	body := RequestClockingBody{
		Date: currentDateTime.ActualDate,
		Hour: currentDateTime.ActualTime,
		Latitude: "0",
		Longitude: "0",
		Timezone: 0,
		Address: "l-location-unavailable",
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




