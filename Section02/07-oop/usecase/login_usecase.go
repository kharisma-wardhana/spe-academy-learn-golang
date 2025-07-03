package usecase

import (
	"errors"

	"github.com/kharisma-wardhana/spe-academy-learn-golang/Section02/07-oop/repository"
)

type LoginUseCase struct {
	// Dependencies can be added here, such as repositories or services
	userRepository repository.IUserRepository
}

// NewLoginUseCase creates a new instance of LoginUseCase
func NewLoginUseCase(userRepo repository.IUserRepository) *LoginUseCase {
	return &LoginUseCase{
		// Initialize dependencies if needed
		userRepository: userRepo,
	}
}

// LoginUser logs in a user with the provided email and password
func (uc *LoginUseCase) LoginUser(email, password string) (string, error) {
	// Here you would typically validate the input, check if the user exists,
	// and then authenticate the user against a database or another storage system.
	// For simplicity, we will just return a dummy token to indicate success.

	// Example validation (not implemented):
	if email == "" || password == "" {
		return "", errors.New("email and password cannot be empty")
	}

	// Example authentication (not implemented):
	user, err := uc.userRepository.FindByEmail(email)
	if err != nil || user.Password != password {
		return "", errors.New("invalid email or password")
	}

	return "dummy_token", nil // Indicating success with a dummy token
}
