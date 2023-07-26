package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	provider   string
	logged_in  bool
	User       string
	Email      string
	password   string
	expires_in float64
}

func InitUserForProvider(provider string, creds ...string) User {
	switch provider {
	case "basic":
		return User{
			provider: "basic",
			User:     creds[0],
			Email:    creds[1],
			password: creds[2],
		}
	// Will flesh out api case later
	case "api":
		return User{}
	default:
		return User{}
	}

}

func (u User) IsValidPassword(test_pass string) bool {
	return u.password == test_pass
}

type LoginPayload struct {
	User     string `json:"user" xml:"user" form:"user"`
	Password string `json:"password" xml:"password" form:"password"`
}

type SignupPayload struct {
	User     string `json:"user" xml:"user" form:"user"`
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"password" xml:"password" form:"password"`
}
