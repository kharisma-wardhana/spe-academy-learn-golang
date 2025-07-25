package usecase

import (
	"errors"

	"github.com/kharisma-wardhana/spe-academy-learn-golang/Section02/07-oop/repository"
)

type RegisterUseCase struct {
	// Dependencies can be added here, such as repositories or services
	userRepository repository.IUserRepository
}

// NewRegisterUseCase creates a new instance of RegisterUseCase
func NewRegisterUseCase(userRepo repository.IUserRepository) *RegisterUseCase {
	return &RegisterUseCase{
		// Initialize dependencies if needed
		userRepository: userRepo,
	}
}

// RegisterUser registers a new user with the provided name and email
func (uc *RegisterUseCase) RegisterUser(name, email, password string) error {
	// Here you would typically validate the input, check if the user already exists,
	// and then save the user to a database or another storage system.
	// For simplicity, we will just return nil to indicate success.

	// Example validation (not implemented):
	if name == "" || email == "" || password == "" {
		return errors.New("name and email cannot be empty")
	}

	// Example saving to a database (not implemented):
	err := uc.userRepository.Save(name, email, password)
	if err != nil {
		return err
	}

	return nil // Indicating success
}
