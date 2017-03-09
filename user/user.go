package user

type User struct {
	email string
}

func (u *User) Email() string {
	return u.email
}
