package auth

type LoginDTO struct {
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"passwordHash" bson:"passwordHash"`
}

type User struct {
	UserID   string `json:"user_id"`
	Role     string `json:"role"`
	UserName string `json:"username"`
	Exp      int64  `json:"exp"`
	// other fields...
}