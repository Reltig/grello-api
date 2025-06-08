package request

type CreateUserRequest struct {
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	FirstName  *string `json:"first_name"`
	SecondName *string `json:"second_name"`
}
