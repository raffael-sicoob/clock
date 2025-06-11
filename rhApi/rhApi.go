package rhapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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


// func RequestClocking(token string) (string, error) {
	
// }




