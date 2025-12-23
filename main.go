package main

import (
	"fmt"
)

func main() {
	// -- ARRAY --
	var colors []string = []string {"Merah", "Kuning", "Hijau", "Biru", "Orange"}
	// --- Short hand ---
	numArr := [4]int {10, 20, 30, 40}
	// -- Another array initiation --
	var usernames = [3]string {"Bambang", "Mulyono", "Sumantod"}
	fmt.Println(usernames[2])
	// -- Array initiation pada index tertentu
	// ? Untuk index array yang tidak diisikan valuenya, maka secara default akan menggunakan zero value tergantung dari tipe data yang digunakan
	var scores = []int {1: 100, 3: 85, 6: 70}
	// ? Untuk assign value pada index tertentu jika sebelumnya array sudah di inisiasi
	scores[2] = 90
	fmt.Println("Nilai scores:", scores[2])

	// -- Array menggunakan elips --
	// ? Value length pada array yang menyesuaikan dengan isinya
	// ? Sebenarnya juga bisa tidak perlu diisi, karna panjang akan menyesuaikan dengan isinya
	var categories = [...]string {"Makanan", "Minuman", "Baju"}
	fmt.Println(categories)

	// -- Array yang ditulis dalam beberapa baris --
	rowedArr := []string {
		"Bambang",
		"Sumantao",
		"Prabowo", // ! Wajib ada koma di value terakhir !
	}
	fmt.Println(rowedArr)

	// --- len() ---
	// ? len() digunakan untuk mendapatkan value berupa panjang dari sebuah array
	valuesArr := [...]int {12, 12, 45, 60}
	fmt.Println(len(valuesArr))

	// -- Zero value pada array --
	// ? Jika nilai pada suatu index kosong, maka secara default akan diisi oleh zero value berdasarkan zero value tiap tipe data
	// ! Hanya bisa dilakukan jika panjang array sudah diinisiasikan
	damages := [5]string {}
	fmt.Println(damages)

	// -- Multidimention array --
	multiDimensionArr := [][]string {
		{"Ujang", "Bambang", "Prabowo", "Subianto"},
		{"Aji", "Gibran", "Suilistiyowati"},
	}
	fmt.Println(multiDimensionArr[0][0])


	// -- FOR LOOP --
	for i := 0; i < 3; i++ {
		fmt.Printf("Value colors index ke %d : %s \n", i, colors[i])
		fmt.Printf("Value numArr index ke %d : %d \n", i, numArr[i])
	}

	fmt.Printf("\n")
	fmt.Printf("--For loop menggunakan len()-- \n")

	// --- For loop menggunakan len() pada condition ---
	// ? len() merupakan panjang dari sebuah array
	for i := 0; i < len(colors); i++ {
		fmt.Printf("Warna: %s\n", colors[i])
	}

	fmt.Printf("\n")
	fmt.Printf("--For loop menggunakan range-- \n")
	// --- For loop menggunakan range ---
	// ? range digunakan untuk shorthand dari perulangan
	// ? range akan berisi value asli array dari tiap index
	// ? i adalah index dari tiap value yg ada di dalam array
	for i, value := range numArr {
		fmt.Printf("Pada index ke-%d, bernilai %d\n", i, value)
	}

	// ? Jika index tidak digunakan
	for _, value := range numArr {
		fmt.Printf("Value: %d\n", value)
	}

	// ? Jika hanya index yang digunakan
	for i := range numArr {
		fmt.Printf("Index: %d\n", i)
	}

	fmt.Println("--- For loop only condition ---")
	// --- for loop hanya dengan kondisi ---
	// ? Ini bentuk jika variable diinisiasikan
	counter := 0
	for counter != 10 {
		fmt.Println("Counter:", counter)
		counter++
	}

	// ? Ini bentuk jika loop hanya dengan condition, tapi variable diinisiasi di for loop
	// ? tapi scopenya hanya bisa digunakan di dalam for itu sendiri
	for index := 0; index < 5; { // ! hati-hati terjadi infinite loop
		fmt.Println("Index:", index)
		index++
	}

	// --- For tanpa kondisi ---
	// ? tidak menuliskan kondisi apapun pada for
	// ? for seperti ini akan terus berjalan jika tidak ada break/kondisi yang menghentikanya. bahkan meskipun tidak iterasinya
	index := 0
	for {
		if index < 10 {
			break // break digunakan untuk menghentikan sebuah proses looping
		}

		fmt.Println("New Index:", index)
		index++
	}

	// --- for menggunakan continue dan break ---
	newIndex := 0
	for {
		if newIndex % 2 != 0 {
			newIndex++ // ! usahakan juga menulis post statement ketika menggunakan continue agar tidak nyangkut
			continue // ? akan mengulang kembali perulangan dari awal
		} else if newIndex == 10 {
			fmt.Println("Looping berhenti di", newIndex)
			break // ? menghentikan perulangan
		}

		fmt.Println("New Index:", newIndex)
		newIndex++
	}

	myValue := 0
	for {
		if myValue % 2 == 0 {
			myValue++
			continue
		} else if myValue > 9 {
			fmt.Println("My value stops at:", myValue)
			break
		}
		
		fmt.Println("My Value:", myValue)
		myValue++
	}

	// -- Range pada string --
	fmt.Println("--- Range pada string ---")
	myString := "Dionovan"
	fmt.Println(len(myString)) // ? Mengambil panjang pada string

	for _, char := range myString {
		fmt.Printf("Tipe dari %c, adalah %T\n", char, char)

		if (char == 'v') {
			fmt.Printf("Char mengandung %c, harus dihentikan\n", char)
			break
		}
	}
}
