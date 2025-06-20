package main

func main() {
	// Deklarasi array integer
	var angka [5]int // Deklarasi array bernama 'angka' dengan panjang 5

	// Inisialisasi array dengan nilai
	angka[0] = 10 // Elemen pertama diisi 10
	angka[1] = 20 // Elemen kedua diisi 20
	angka[2] = 30 // Elemen ketiga diisi 30
	angka[3] = 40 // Elemen keempat diisi 40
	angka[4] = 50 // Elemen kelima diisi 50

	// Cetak isi array
	for i, num := range angka {
		println("Angka di indeks", i, "adalah", num)
	}

	// Deklarasi dan inisialisasi array string
	kota := [3]string{"Jakarta", "Bandung", "Surabaya"}
	for i, k := range kota {
		println("Kota di indeks", i, "adalah", k)
	}

	// Deklarasi array 2 dimensi
	matriks := [2][3]int{
		{1, 2, 3}, // Baris pertama
		{4, 5, 6}, // Baris kedua
	}
	// Cetak isi array 2 dimensi
	for i, baris := range matriks {
		for j, nilai := range baris {
			println("Nilai di baris", i, "kolom", j, "adalah", nilai)
		}
	}

	// Deklarasi array dengan nilai awal 0
	var arrayKosong [5]int
	for i, num := range arrayKosong {
		println("Array kosong di indeks", i, "adalah", num)
	}

	// Deklarasi array string dengan ukuran tetap
	var buah = [4]string{"Apel", "Pisang", "Ceri", "Kurma"}
	for i, f := range buah {
		println("Buah di indeks", i, "adalah", f)
	}

	// Deklarasi array dengan cara singkat
	arrayPendek := [3]float64{1.1, 2.2, 3.3}
	for i, v := range arrayPendek {
		println("Nilai di indeks", i, "adalah", v)
	}

	// Deklarasi array boolean
	var arrayBool = [2]bool{true, false}
	for i, b := range arrayBool {
		println("Boolean di indeks", i, "adalah", b)
	}

	var array = [...]int{1, 2, 3, 4, 5}
	for i, v := range array {
		println("Array di indeks", i, "adalah", v)
	}

	// Menambah nilai ke array tidak bisa, gunakan slice
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	println("Slice setelah ditambah 4:", slice)

	slice = append(slice, 5)
	println("Slice setelah ditambah 5:", slice)

	// Penjelasan array vs slice
	// Array: ukuran tetap, tidak bisa diubah, tipe data nilai (copy saat dikirim ke fungsi)
	// Slice: ukuran fleksibel, bisa ditambah/dikurangi, tipe data referensi (perubahan mempengaruhi data asli)
	// Gunakan array jika jumlah data pasti, gunakan slice jika jumlah data bisa berubah

	// Catatan: Array di Go ukurannya tetap, jika butuh koleksi data yang bisa berubah gunakan slice.
}
