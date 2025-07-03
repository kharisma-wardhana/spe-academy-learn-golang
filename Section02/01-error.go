package main

import "fmt"

// ValidationError adalah tipe error khusus yang mengimplementasikan interface error.
// Tipe ini berisi pesan yang menjelaskan kesalahan validasi.
// Dengan ini, kita bisa memberikan informasi lebih jelas saat terjadi error.
// Tipe error ini juga memudahkan kita membedakan jenis error dalam aplikasi.
// Sangat berguna untuk menangani error dan proses debugging.
type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

func validateInput(input string) error {
	if input == "" {
		return &ValidationError{Message: "Input cannot be empty"}
	}
	return nil
}

func main() {
	// Ambil input dari user
	// Misalnya, kita bisa menggunakan fmt.Scanln untuk mengambil input dari console.
	var input string
	fmt.Print("Enter input: ")
	_, err := fmt.Scanln(&input)

	// Check jika ada error saat membaca input.
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Validasi input menggunakan fungsi validateInput.
	// Jika input tidak valid, kita akan mendapatkan error dari fungsi ini.
	err = validateInput(input)
	if err != nil {
		// Jika error adalah ValidationError, kita bisa mengakses pesan errornya.
		// Ini memungkinkan kita memberikan feedback yang lebih spesifik kepada user.
		if ve, ok := err.(*ValidationError); ok {
			fmt.Println(ve.Message)
		} else {
			// Jika error bukan ValidationError, kita bisa menangani error lain.
			fmt.Println("An unexpected error occurred")
		}
		return
	}

	fmt.Println("Input is valid")
}
