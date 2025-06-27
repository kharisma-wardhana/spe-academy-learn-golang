package main

import (
	"fmt"

	"github.com/kharisma-wardhana/spe-academy-learn-golang/Section02/07-oop/repository"

	LoginUseCase "github.com/kharisma-wardhana/spe-academy-learn-golang/Section02/07-oop/usecase"
	RegisterUseCase "github.com/kharisma-wardhana/spe-academy-learn-golang/Section02/07-oop/usecase"
)

func main() {
	// This is the main entry point of the application.
	// You can initialize your application here, such as setting up routes,
	// connecting to a database, or starting a server.

	userRepo := repository.NewUserRepository() // Assuming UserRepository is defined in the repository package

	// For example, you might want to create an instance of RegisterUseCase
	// and call its RegisterUser method to register a new user.

	registerUseCase := RegisterUseCase.NewRegisterUseCase(userRepo)
	err := registerUseCase.RegisterUser("John Doe", "john.doe@email.com", "password123")
	if err != nil {
		fmt.Println("Error registering user:", err)
	} else {
		fmt.Println("User registered successfully")
	}

	// Similarly, you can create an instance of LoginUseCase
	// and call its LoginUser method to log in a user.
	loginUseCase := LoginUseCase.NewLoginUseCase(userRepo)
	token, err := loginUseCase.LoginUser("john.doe@email.com", "password123")
	if err != nil {
		fmt.Println("Error logging in user:", err)
	} else {
		fmt.Println("User logged in successfully, token:", token)
	}
}
