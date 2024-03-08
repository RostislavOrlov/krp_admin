package dto

import "krp_admin/internal/entities"

type AddUserResponse struct {
	EmployeeId int    `json:"employee_id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Passport   string `json:"passport"`
	Inn        string `json:"inn"`
	Snils      string `json:"snils"`
	Birthday   string `json:"birthday"`
	Role       string `json:"role"`
}

type EditUserResponse struct {
	EmployeeId int    `json:"employee_id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	Passport   string `json:"passport"`
	Inn        string `json:"inn"`
	Snils      string `json:"snils"`
}

type DeleteUserResponse struct {
	EmployeeId int `json:"employee_id"`
}

type ListUserResponse struct {
	Users []entities.User `json:"users"`
}
