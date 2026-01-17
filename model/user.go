package model

type User struct {
	Conditions map[string]any
}

func NewUser() *User {
	return &User{
		Conditions: make(map[string]any),
	}
}

func (u *User) With(attribute string, value any) *User {
	u.Conditions[attribute] = value
	return u
}
