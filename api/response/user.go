package response

import "grello-api/internal/model"

type UserResponse struct {
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name"`
	SecondName *string `json:"second_name"`
}

func (resp UserResponse) FromModel(user *model.User) UserResponse {
	return UserResponse{
		Username: user.Username,
		Email: user.Email,
		FirstName: user.FirstName,
		SecondName: user.SecondName,
	}
}