package request

type Create struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`

	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	DateOfBirth string `json:"date_of_birth"`

	Country    string `json:"country"`
	City       string `json:"city"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
}
