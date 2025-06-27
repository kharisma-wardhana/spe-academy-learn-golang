package main

import "time"

func main() {
	// Contoh penggunaan concurrency dalam Go
	// Concurrency memungkinkan kita untuk menjalankan beberapa tugas secara bersamaan
	// Ini sangat berguna untuk meningkatkan performa aplikasi dan mengoptimalkan penggunaan sumber daya
	// Kita bisa menggunakan goroutine untuk menjalankan fungsi secara bersamaan
	// Goroutine adalah fungsi yang berjalan secara asinkron dan tidak memblokir eksekusi program utama
	// Dengan goroutine, kita bisa menjalankan beberapa tugas secara bersamaan tanpa harus menunggu satu sama lain selesai
	// Ini membuat aplikasi kita lebih responsif dan efisien dalam penggunaan sumber daya.
	// Contoh penggunaan goroutine
	// Kita bisa membuat goroutine dengan menambahkan kata kunci 'go' sebelum pemanggilan fungsi
	go func() {
		// Fungsi ini akan berjalan secara asinkron
		// Kita bisa menambahkan logika yang ingin dijalankan secara bersamaan di sini
		println("Goroutine is running use anonymous function")
	}()

	go printMessage()

	// Kita bisa menambahkan logika lain di sini yang akan berjalan secara bersamaan dengan goroutine
	println("Main function is running")
	// Untuk memastikan goroutine selesai sebelum program utama berakhir,
	// kita bisa menggunakan time.Sleep atau sync.WaitGroup
	// Namun, dalam contoh ini kita tidak menggunakan itu, sehingga program utama akan segera selesai
	// Jika kita ingin menunggu goroutine selesai, kita bisa menggunakan channel atau sync.WaitGroup
	// Misalnya, kita bisa menggunakan time.Sleep untuk memberikan waktu bagi goroutine untuk selesai
	time.Sleep(time.Second * 1)
}

func printMessage() {
	// Fungsi ini akan dijalankan dalam goroutine
	println("Hello from printMessage function")
}
