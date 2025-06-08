package rhapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Login(user string, password string) (string, error){
	

	loginRequest := LoginRequest{
		User: user,
		Password: password,
	}

	json, err := json.Marshal(loginRequest)
	if err != nil {
		fmt.Println("Error marshalling login request", err)
		return "", err
	}


	res, err := http.Post(LoginUrl, "application/json", bytes.NewBuffer(json))
	if err != nil {
		fmt.Println("Error sending request", err)
		return "", err
	}

	defer res.Body.Close()
	
	token := res.Header.Get("Set-Authorization")
	return token, nil
}





