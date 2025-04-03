package requests

type UpdateUserRequest struct {
	Firstname string `json:"firstname" binding:"required,min=2,max=50"`
	Lastname  string `json:"lastname" binding:"required,min=2,max=50"`
	Login     string `json:"login" binding:"required,min=2,max=50"`
}
