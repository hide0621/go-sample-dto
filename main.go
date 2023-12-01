package main

import "fmt"

type UserDTO struct {
	ID       int
	Username string
	Email    string
}

func getUserFromDatabase(userID int) UserDTO {
	// 通常はデータベースからユーザーデータを取得するロジックがここに入る
	// 仮のデータで代用
	return UserDTO{
		ID:       userID,
		Username: "sampleuser",
		Email:    "sampleuser@example.com",
	}
}

// アプリケーションの別の層にユーザーデータを渡す関数
func processUser(user UserDTO) {
	// ユーザーデータを処理するロジックがここに入る
	fmt.Printf("Processing user: ID=%d, Username=%s, Email=%s\n", user.ID, user.Username, user.Email)
}

func main() {

	userID := 1
	userFromDB := getUserFromDatabase(userID)

	processUser(userFromDB)
}
