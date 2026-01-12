package models

type User struct {
	Username string
	Password []byte
	OTP      string
}
