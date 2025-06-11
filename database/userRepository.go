package database

import (
	"encoding/json"
	"fmt"

	bolt "go.etcd.io/bbolt"
)


type User struct {
	Name string
	Username string
	Password string
}


func CreateUser(user User)  {
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(ClockBucket))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}
		userBytes, err := json.Marshal(user)
		if err != nil {
			fmt.Println("Error marshalling user")
		}
		err = bucket.Put([]byte("user"), userBytes)
		if err != nil {
			fmt.Println("Error saving user")
		}
		
		fmt.Println("User created")
		return nil
	})
	
}

func GetUser() (User) {
	var user User
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(ClockBucket))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}
		userBytes := bucket.Get([]byte("user"))
		if userBytes == nil {
			return fmt.Errorf("user not found")
		}
		err := json.Unmarshal(userBytes, &user)
		if err != nil {
			return fmt.Errorf("error unmarshalling user")
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error getting user")
	}
	return user
}


func SaveToken(token string) error {
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(ClockBucket))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}
		err := bucket.Put([]byte("token"), []byte(token))
		return err
	})
	return nil
}

func GetToken() (string, error) {
	var token string
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(ClockBucket))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}
		tokenBytes := bucket.Get([]byte("token"))
		if tokenBytes == nil {
			return fmt.Errorf("token not found")
		}
		token = string(tokenBytes)
		return nil
	})
	return token, err
}