package main

import "fmt"

// Struct sangat berguna untuk mengelompokkan data yang terkait dan membuat kode lebih terstruktur
// Dengan menggunakan struct, kita bisa membuat kode lebih mudah dibaca dan di maintain
// Ini juga membantu kita mengorganisir data dalam aplikasi kita.
type User struct {
	ID    int
	Name  string
	Email string
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Order struct {
	ID        int
	UserID    int
	ProductID int
	Quantity  int
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// Struct dalam Go juga mendukung nested struct
// Kita bisa membuat struct yang memiliki field berupa struct lain
// Contoh struct dengan nested struct
type UserWithAddress struct {
	ID      int
	Name    string
	Email   string
	Address Address // Field Address yang merupakan struct lain
}

func main() {
	// Contoh penggunaan struct
	var user User
	user.ID = 1
	user.Name = "John Doe"
	user.Email = "john.doe@email.com"

	// Kita bisa membuat struct secara langsung dengan inisialisasi
	product := Product{1, "Laptop", 999.99}

	// Atau kita bisa menggunakan inisialisasi dengan nama field
	// Ini membuat kode lebih jelas dan mudah dibaca
	order := Order{
		ID:        1,
		UserID:    user.ID,
		ProductID: product.ID,
		Quantity:  1,
	}

	// Di sini kita bisa menggunakan user, product, dan order sesuai kebutuhan
	// Misalnya, kita bisa mencetak informasi user, product, dan order
	fmt.Println(user)
	fmt.Println(product)
	fmt.Println(order)

	// Kita juga bisa mengakses field dari struct
	fmt.Printf("User: %s, Email: %s\n", user.Name, user.Email)
	fmt.Printf("Product: %s, Price: %.2f\n", product.Name, product.Price)
	fmt.Printf("Order ID: %d, User: %s, Product: %s, Quantity: %d\n",
		order.ID, user.Name, product.Name, order.Quantity)

	// Kita bisa mengubah field dari struct
	// Ini memungkinkan kita untuk memperbarui data tanpa membuat struct baru
	fmt.Println("Original User Name:", user.Name)
	user.Name = "Jane Lee"
	fmt.Println("Updated User Name:", user.Name)

	// Kita juga bisa mengoper struct ke fungsi lain
	printUserInfo(user)
	// Fungsi ini menerima struct User sebagai parameter
	// Ini memungkinkan kita untuk mengoper data user ke fungsi lain
	// Dengan cara ini, kita bisa memisahkan logika dan membuat kode lebih modular
	// Misalnya, kita bisa membuat fungsi untuk mencetak informasi user
	// Ini membuat kode lebih terstruktur dan mudah dipahami

	// Contoh penggunaan nested struct
	address := Address{
		Street: "123 Main St",
		City:   "Springfield",
		State:  "IL",
		Zip:    "62701",
	}
	userWithAddress := UserWithAddress{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: address, // Menggunakan struct Address sebagai field
	}
	// Kita bisa mengakses field dari nested struct
	fmt.Printf("User with Address: %s, Email: %s, Address: %s, %s, %s, %s\n",
		userWithAddress.Name, userWithAddress.Email,
		userWithAddress.Address.Street, userWithAddress.Address.City,
		userWithAddress.Address.State, userWithAddress.Address.Zip)

	// Contoh penggunaan struct dengan pointer
	// Kita bisa membuat pointer ke struct untuk menghemat memori
	userPointer := &user // Membuat pointer ke struct User
	fmt.Printf("User Pointer: ID: %d, Name: %s, Email: %s\n",
		userPointer.ID, userPointer.Name, userPointer.Email)
	// Dengan pointer, kita bisa mengubah data struct tanpa membuat salinan
	userPointer.Name = "Alice Smith"
	fmt.Printf("Updated User Pointer Name: %s\n", userPointer.Name)
	fmt.Printf("User after pointer update: ID: %d, Name: %s, Email: %s\n",
		user.ID, user.Name, user.Email)

	// Kita juga bisa membuat fungsi yang menerima pointer ke struct
	printUserInfoPointer(userPointer)
	// Fungsi ini menerima pointer ke struct User sebagai parameter
	// Ini memungkinkan kita untuk mengoper data user ke fungsi lain tanpa membuat salinan
}

func printUserInfo(user User) {
	// Fungsi ini menerima struct User sebagai parameter
	// Ini memungkinkan kita untuk mengoper data user ke fungsi lain
	fmt.Printf("User ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
}

// Contoh handling error dengan struct
// Kita bisa membuat struct khusus untuk menangani error
type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func printUserInfoPointer(user *User) {
	// Fungsi ini menerima pointer ke struct User sebagai parameter
	// Ini memungkinkan kita untuk mengoper data user ke fungsi lain tanpa membuat salinan
	if user == nil {
		fmt.Println("User is nil")
		return
	}
	fmt.Printf("User Pointer: ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
}

func handleError(err error) {
	if err != nil {
		// Kita bisa menggunakan type assertion untuk memeriksa apakah error adalah CustomError
		if customErr, ok := err.(*CustomError); ok {
			fmt.Printf("Error Code: %d, Message: %s\n", customErr.Code, customErr.Message)
		} else {
			fmt.Println("An unexpected error occurred:", err)
		}
	}
}
