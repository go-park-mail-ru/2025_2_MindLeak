package auth

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginOutput struct {
}

type UserRegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserRegisterOutput struct {
}

type UserResponse struct {
}
