package users

type CreateUserDTO struct {
	Username     string `json:"username" validate:"required,min=3,max=32"`
	Password     string `json:"password" validate:"required,min=6"`
	Email        string `json:"email" validate:"required,email"`
	Role         string `json:"role" validate:"required,oneof=admin user"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

type UpdateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UpdateUserPasswordDTO struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}