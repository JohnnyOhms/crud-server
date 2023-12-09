package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func ParseData(r *http.Request, i interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	Err := Decoding(body, i)
	return Err
}
