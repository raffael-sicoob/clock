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


func CreateUser(user User) error {
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(ClockBucket))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}
		userBytes, err := json.Marshal(user)
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("user"), userBytes)
		return err
	})
	return nil
}

func GetUser() (User, error) {
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
			return err
		}
		return nil
	})
	return user, err
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