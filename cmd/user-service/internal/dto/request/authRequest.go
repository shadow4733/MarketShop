package request

type RegisterRequest struct {
	Username string  `json:"username" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Phone    *string `json:"phone" binding:"required"`

	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	MiddleName string `json:"middle_name" binding:"required"`
	Country    string `json:"country" binding:"required"`
	City       string `json:"city" binding:"required"`
}
