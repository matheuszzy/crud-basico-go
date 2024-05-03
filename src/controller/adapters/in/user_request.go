package in

type UserRequest struct {
	Email    string `json:"email" binding:"required,email" example:"test@test.com"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*" example:"password#@#@!2121"`
	Name     string `json:"name" binding:"required,min=4,max=100" example:"John Doe"`
	Age      int8   `json:"age" binding:"required,gte=0,lte=130"`
}

type UserUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=4,max=100" example:"John Doe"`
	Age  int8   `json:"age" binding:"omitempty,gte=0,lte=130"`
}
