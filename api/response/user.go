package response

import "grello-api/internal/model"

type User struct {
	ID		   uint    `json:"id"`
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	FirstName  *string `json:"first_name"`
	SecondName *string `json:"second_name"`
}

func (resp User) FromModel(user *model.User) User {
	return User{
		ID:			user.ID,
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
	}
}
