package requests

type CreateUserRequest struct {
	Firstname string `json:"firstname" binding:"required,min=2,max=50"`
	Lastname  string `json:"lastname" binding:"required,min=2,max=50"`
	Login     string `json:"login" binding:"required,min=2,max=50"`
	Password  string `json:"password" binding:"required,min=8,max=72"`
}
