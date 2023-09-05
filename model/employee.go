package model

type Employee struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}
