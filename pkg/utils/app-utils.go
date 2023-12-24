package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func Encoding(data interface{}) ([]byte, error) {
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error in marshling to json", err)
	}
	return json, err
}

func Decoding(data []byte, i interface{}) error {
	err := json.Unmarshal(data, &i)
	if err != nil {
		fmt.Println("Error in unmarshing to Go data", err)
	}
	return err
}

func ParseBody(r *http.Request, i interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = Decoding(body, i)
	return err
}

func GenerateUserId() string {
	const charaset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 30)
	for i := range b {
		b[i] = charaset[rand.Intn(len(charaset))]
	}
	return string(b)
}
