package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Fungsi ini menerima sync.WaitGroup sebagai parameter
// Sync.WaitGroup digunakan untuk menunggu beberapa goroutine selesai sebelum melanjutkan eksekusi
// Kita menambahkan 1 ke WaitGroup sebelum menjalankan goroutine ini
// Dengan defer wg.Done(), kita memastikan bahwa ketika fungsi ini selesai,
// kita akan mengurangi jumlah goroutine yang sedang ditunggu oleh WaitGroup.
// Ini penting untuk memastikan bahwa program utama tidak selesai sebelum semua goroutine selesai.
func taskA(wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done when it returns

	// Sample long running task
	// This function simulates a long-running task that takes 2 seconds to complete.
	// It uses a for loop to iterate 10 times, simulating some work being done
	// in each iteration. The time.Sleep function is used to pause the execution for 200 milliseconds
	// in each iteration, simulating a delay in processing.
	for i := 0; i < 10; i++ {
		// Simulate some work being done
		// In a real-world scenario, this could be any long-running task such as processing data,
		// making network requests, or performing complex calculations.
		// Here, we just use time.Sleep to simulate a delay.
		time.Sleep(200 * time.Millisecond)
	}
	// After the loop completes, we print a message indicating that the task is done.
	println("Task A completed")
}

func taskB(wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done when it returns

	// Sample long running task
	// This function simulates another long-running task that takes 3 seconds to complete.
	// Similar to taskA, it uses a for loop to iterate 15 times, simulating some work being done
	// in each iteration. The time.Sleep function is used to pause the execution for 200 milliseconds
	// in each iteration, simulating a delay in processing.
	for i := 0; i < 15; i++ {
		// Simulate some work being done
		time.Sleep(200 * time.Millisecond)
	}
	// After the loop completes, we print a message indicating that the task is done.
	println("Task B completed")
}

func main() {
	runtime.GOMAXPROCS(2) // Set the number of OS threads to use for executing goroutines
	// This line sets the maximum number of operating system threads that can execute Go code simultaneously.
	// By default, Go uses a single thread, but we can increase this number to allow
	// multiple goroutines to run concurrently on different threads.
	// This is useful for CPU-bound tasks where we want to utilize multiple cores of the CPU
	// to improve performance. In this case, we set it to 2, allowing two goroutines to run concurrently.
	fmt.Println("Starting tasks...")
	start := time.Now() // Record the start time of the tasks

	// Menggunakan goroutine untuk menjalankan fungsi secara bersamaan
	// Dalam contoh ini, kita menjalankan dua fungsi (taskA dan taskB) secara
	// bersamaan menggunakan goroutine. Goroutine adalah cara Go untuk menjalankan fungsi secara
	// asinkron, yang memungkinkan kita untuk menjalankan beberapa tugas secara bersamaan tanpa harus
	// menunggu satu sama lain selesai. Ini sangat berguna untuk meningkatkan performa aplikasi
	// dan mengoptimalkan penggunaan sumber daya.

	// Untuk memastikan program utama tidak langsung selesai sebelum goroutine selesai,
	// kita bisa menggunakan sync.WaitGroup atau time.Sleep.
	// Dalam contoh ini, kita menggunakan sync.WaitGroup untuk menunggu
	// kedua goroutine selesai sebelum program utama berakhir.
	fmt.Println("Waiting for tasks to complete...")
	wg := sync.WaitGroup{}
	wg.Add(2) // Menambahkan dua goroutine yang akan ditunggu

	go taskA(&wg) // Menjalankan taskA dalam goroutine
	go taskB(&wg) // Menjalankan taskB dalam goroutine

	wg.Wait()                    // Menunggu sampai semua goroutine selesai
	elapsed := time.Since(start) // Menghitung waktu yang dibutuhkan untuk menyelesaikan tugas
	fmt.Printf("All tasks completed in %s\n", elapsed)
}
