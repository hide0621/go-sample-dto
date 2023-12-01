package main

import (
	"encoding/json"
	"fmt"
)

type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {

	user := UserDTO{
		ID:       1,
		Username: "sampleuser",
		Email:    "sampleuser@example.com",
	}

	// マーシャリング（構造体をJSONに変換）
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	fmt.Println("JSON Data:", string(jsonData))

	// アンマーシャリング（JSONを構造体に変換）
	var newUser UserDTO
	err = json.Unmarshal(jsonData, &newUser)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	fmt.Printf("Unmarshalled User: %+v\n", newUser)
}
