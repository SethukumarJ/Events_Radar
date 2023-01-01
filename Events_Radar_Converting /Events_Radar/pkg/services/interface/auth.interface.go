package interfaces

// AuthService is the interface for authentication service
type AuthService interface {
	VerifyAdmin(email string, password string) error
	VerifyUser(email string, password string) error
}