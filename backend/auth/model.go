package main

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
	Phone:    "49123454322", //this is a random number
}
