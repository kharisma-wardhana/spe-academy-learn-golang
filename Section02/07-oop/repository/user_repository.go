package repository

type IUserRepository interface {
	// Save saves a new user to the repository
	Save(name, email, password string) error
	// FindByEmail finds a user by their email in the repository
	FindByEmail(email string) (User, error)
}

type User struct {
	Name        string
	Email       string
	Password    string
	AccessToken string // This could be used for authentication purposes
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository() *User {
	return &User{
		// Initialize dependencies if needed
	}
}

// Save saves a new user to the repository
func (r *User) Save(name, email, password string) error {
	// Here you would typically save the user to a database or another storage system.
	// For simplicity, we will just return nil to indicate success.
	return nil // Indicating success
}

// FindByEmail finds a user by their email in the repository
func (r *User) FindByEmail(email string) (User, error) {
	// Here you would typically query the user from a database or another storage system.
	// For simplicity, we will just return a dummy user to indicate success.
	user := User{
		Name:        "Dummy User",
		Email:       email,
		Password:    "password123", // This should be hashed in a real application
		AccessToken: "dummy_access_token",
	}

	return user, nil // Indicating success with a dummy user
}
