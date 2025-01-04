package data

const (
	RoleClient string = "client"
	RoleSeller string = "seller"
	RoleAdmin  string = "admin"
)

type User struct {
	UserPublic
	Password string `json:"password"`
}

type UserPublic struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
