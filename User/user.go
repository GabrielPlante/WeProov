package weproovuser

//A standard user model
type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordhash"`
	Email        string `json:"email"`
}

//Hash a password
func HashPassword(password string) string {
	return password + "hashed" //This is not a hash, but no time
}
