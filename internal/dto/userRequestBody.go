package dto

type AddUserResponseBody struct {
	Name  string `json:"name" binding:"required,min=2,max=50"`
	Email string `json:"email" binding:"required,email"`
}

type UpdateUserResponseBody struct {
	Name string `json:"name" binding:"required,min=2,max=50"`
}
