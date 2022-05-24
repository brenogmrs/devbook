package models

type ChangePassword struct {
	NewPassword     string `json:"new_password"`
	CurrentPassword string `json:"current_password"`
}
