package models

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}


func (u *User) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Phone, validation.Required, validation.Match(
			regexp.MustCompile(`^(?:\+251[79]\d{8}|0[79]\d{8})$`),
		).Error("invalid phone number"),
		),
		validation.Field(&u.Role, validation.Required, validation.In("user", "admin").Error("role must be 'user' or 'admin'")),
	)
	if err != nil {
		return err
	}
	return nil
}
