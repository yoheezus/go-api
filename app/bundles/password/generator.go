package password

type PasswordRequest struct {
	Length int
}

// GeneratePassword uses request to make a password
func GeneratePassword(req PasswordRequest) string {
	return "password"
}
