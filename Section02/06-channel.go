package main

import (
	"fmt"
	"time"
)

func getOngkirJNE(ch chan string) {
	//Simulasi get data dari API JNE
	// Biasanya kita akan melakukan request HTTP ke API JNE untuk mendapatkan ongkir
	// Namun, untuk contoh ini kita akan mensimulasikan dengan mengirimkan nilai ongkir melalui channel
	time.Sleep(2 * time.Second) // Simulasi delay untuk request API
	// Misalnya, ongkir JNE adalah 10.000
	ch <- "10000" // Simulasi ongkir JNE
}

func getOngkirJNT(ch chan string) {
	// Simulasi get data dari API JNT
	// Biasanya kita akan melakukan request HTTP ke API JNT untuk mendapatkan ongkir
	// Namun, untuk contoh ini kita akan mensimulasikan dengan mengirimkan nilai ongkir melalui channel
	time.Sleep(3 * time.Second) // Simulasi delay untuk request API
	// Misalnya, ongkir JNT adalah 12000
	ch <- "12000" // Simulasi ongkir JNT
}

func getOngkirTIKI(ch chan string) {
	// Simulasi get data dari API TIKI
	// Biasanya kita akan melakukan request HTTP ke API TIKI untuk mendapatkan ongkir
	// Namun, untuk contoh ini kita akan mensimulasikan dengan mengirimkan nilai ongkir melalui channel
	time.Sleep(1 * time.Second) // Simulasi delay untuk request API
	// Misalnya, ongkir TIKI adalah 8000
	ch <- "8000" // Simulasi ongkir TIKI
}

func main() {
	start := time.Now() // Mencatat waktu mulai eksekusi
	fmt.Println("Mulai mendapatkan ongkir...")
	// Contoh penggunaan channel dalam Go
	// Channel memungkinkan kita untuk mengirim dan menerima data antar goroutine
	// Ini sangat berguna untuk komunikasi antar goroutine dan sinkronisasi
	// Kita bisa membuat channel dengan menggunakan make(chan TipeData)
	// Channel juga mendukung buffered channel, yang memungkinkan kita untuk mengirim
	// beberapa data sekaligus tanpa harus menunggu penerima
	// Contoh penggunaan channel dalam Go:
	channelJNE := make(chan string)
	channelJNT := make(chan string)
	channelTIKI := make(chan string)

	go getOngkirJNE(channelJNE)
	go getOngkirJNT(channelJNT)
	go getOngkirTIKI(channelTIKI)

	// Menerima ongkir dari masing-masing channel
	ongkirJNEValue := <-channelJNE
	ongkirJNTValue := <-channelJNT
	ongkirTIKIValue := <-channelTIKI

	// Setelah menerima ongkir dari masing-masing channel, kita bisa menggunakan nilainya
	fmt.Println("Ongkir JNE:", ongkirJNEValue)
	fmt.Println("Ongkir JNT:", ongkirJNTValue)
	fmt.Println("Ongkir TIKI:", ongkirTIKIValue)

	elapsed := time.Since(start) // Menghitung waktu yang dibutuhkan untuk mendapatkan ongkir
	fmt.Printf("Waktu yang dibutuhkan untuk mendapatkan ongkir: %s\n", elapsed)
	// Dengan menggunakan channel, kita bisa mendapatkan ongkir dari beberapa API secara bersamaan
	// Ini membuat aplikasi kita lebih responsif dan efisien dalam penggunaan sumber daya
}
