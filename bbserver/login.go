package main

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Success bool   `json:"success"`
	Message string `json:message"`
}
