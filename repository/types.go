// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type RegistrationRequest struct {
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	Password    string `json:"password"`
}

type Users struct {
	ID           int    `json:"id"`
	FullName     string `json:"full_name"`
	Password     string `json:"password"`
	PhoneNumber  string `json:"phone_number"`
	LoginCounter int    `json:"login_counter"`
}
