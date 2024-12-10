package data

type User struct {
	UserPublic
	Password string `json:"password"`
}

type UserPublic struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}
