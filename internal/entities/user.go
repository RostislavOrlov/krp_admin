package entities

type User struct {
	Id         int    `json:"id"`
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
