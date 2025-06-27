package main

import "fmt"

type UserPointer struct {
	ID    int
	Name  string
	Email string
}

func main() {
	// Contoh penggunaan pointer dalam go
	// Pointer memungkinkan kita untuk mengakses dan mengubah nilai variabel tanpa membuat salinan
	// Ini sangat berguna untuk menghemat memori dan meningkatkan performa aplikasi
	// Kita bisa membuat pointer ke variabel, struct, atau tipe data lainnya
	// Pointer juga memungkinkan kita untuk mengoper data ke fungsi tanpa membuat salinan
	// Dengan pointer, kita bisa mengubah nilai variabel di dalam fungsi dan perubahan tersebut akan
	// terlihat di luar fungsi. Ini sangat berguna untuk menghemat memori dan meningkatkan performa aplikasi.
	// Contoh penggunaan pointer dalam go:
	var x int = 10
	var y *int = &x // Membuat pointer ke variabel x
	*y = 20         // Mengubah nilai x melalui pointer y
	// Setelah mengubah nilai x melalui pointer y, nilai x akan menjadi 20
	// Kita bisa mencetak nilai x untuk melihat perubahan
	println("Nilai x setelah diubah melalui pointer y:", x)
	// Output: Nilai x setelah diubah melalui pointer y: 20

	// Kita juga bisa membuat pointer ke struct
	// Pointer ke struct memungkinkan kita untuk mengakses dan mengubah nilai field struct tanpa membuat salinan
	user := UserPointer{ID: 1, Name: "John Doe", Email: "john.doe@email.com"}
	userPointer := &user          // Membuat pointer ke struct User
	userPointer.Name = "Jane Doe" // Mengubah field Name melalui pointer
	// Setelah mengubah field Name melalui pointer, field Name pada struct user akan menjadi "Jane Doe"
	// Kita bisa mencetak field Name untuk melihat perubahan
	println("Nama user setelah diubah melalui pointer:", user.Name)
	// Output: Nama user setelah diubah melalui pointer: Jane Doe

	// Contoh penggunaan pointer pada fungsi sebagai parameter
	// Kita bisa membuat fungsi yang menerima pointer sebagai parameter
	printUserInfoPointerV2(&user)
	// Fungsi ini menerima pointer ke struct User sebagai parameter
	// Ini memungkinkan kita untuk mengoper data user ke fungsi lain tanpa membuat salinan
	// Dengan cara ini, kita bisa menghemat memori dan meningkatkan performa aplikasi

	// Pointer menunjukkan alamat memori dari variabel atau struct
	// Kita bisa mencetak alamat memori dari variabel atau struct menggunakan %p
	fmt.Printf("Alamat memori dari userPointer: %p \n", userPointer)

	// Kita bisa mengakses nilai dari pointer dengan menggunakan tanda bintang (*)
	fmt.Println("User :", *userPointer)
}

func printUserInfoPointerV2(user *UserPointer) {
	// Fungsi ini menerima pointer ke struct User sebagai parameter
	// Ini memungkinkan kita untuk mengoper data user ke fungsi lain tanpa membuat salinan
	println("User ID:", user.ID)
	println("User Name:", user.Name)
	println("User Email:", user.Email)
	// Dengan pointer, kita bisa mengakses dan mengubah nilai field struct tanpa membuat salinan
	// Ini sangat berguna untuk menghemat memori dan meningkatkan performa aplikasi
}
