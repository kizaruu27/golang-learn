package main

import (
	"fmt"
)

func main() {
	// [TIPE DATA GOLANG]
	// -- Boolean --
	fmt.Println(true)
	fmt.Println(false)

	// -- String --
	fmt.Println("Hallo, Dunia!")
	fmt.Println(`Code with "Golang"`)

	// -- Numeric --
	fmt.Println(123)
	fmt.Println(10 + 10)

	// [STRING FORMATTING]
	fmt.Println("Umur saya saat ini ", 25)
	fmt.Printf("Saat ini saya berada di %v \n", "Jakarta")
	var message = fmt.Sprintf("Berat badan saya saat ini %dkg, Saat ini saya berusia %d tahun", 90, 20) // Sprintf dapat digunakan untuk menyimpan output value ke dalam sebuah variable
	fmt.Println(message)

	// -- Flag --
	// %v: Nilai dengan tipe data apapun
	fmt.Printf("Saya saat ini sudah memiliki %v anak, dan rumah di daerah %v \n", 10, "Pondok Indah")
	// %T: Untuk menampilkan nama tipe data dari suatu value/variable
	fmt.Printf("%T \n", false)
	// %t Menampilkan booelan
	fmt.Printf("Balon itu hijau: %t \n", false)
	// %s Menampilkan string
	fmt.Printf("Halo, nama saya: %s \n", "Eko Kamirudin")
	// %d Menampilkan integer
	fmt.Printf("Sekarang umur saya %d tahun \n", 10)
	// %b Menampilkan binary dari integer
	fmt.Printf("Value binary: %b \n", 100)
	// %o Menampilkan oktal dari integer
	fmt.Printf("Value oktal: %o \n", 100)
	// %x Menampilkan oktal dari integer
	fmt.Printf("Value hexadecimal: %x \n", 100000)
	// %e Menampilkan nilai eksponen dari float
	fmt.Printf("Value exponent: %e \n", 30.5)
	// %f Menampilkan float
	fmt.Printf("Jari-jari lingkaran: %f \n", 3.14)
	// %q Menampilkan quote/petik dua
	fmt.Printf("Selalu bilang %q \n", "Maaf")

	// -- ESCAPE CHARACTER --
	// \t Untuk memberikan 1 tab
	fmt.Printf(" \tKalimat dengan 1 tab\n")
	// \n Untuk membuat new line
	fmt.Printf("Ini adalah kalimat dengan 1 enter/new line \n")
	// /"" Untuk membuat quote/petik dua
	fmt.Printf("Nama saya \"Udin\"\n")
	// \\ Untuk membuat backslash
	fmt.Printf("\\ Ini adalah kalimat menggunakan backslash \\ \n")

	// [VARIABLE]
	// Di deklarasi dengan menggunakan keyword var
	// Value dari variable dapat diubah2 selama masih menggunakan tipe data yang sama
	// ! Variable wajib digunakan, jika ada variable yang unused/tidak digunakan maka akan muncul syntax error
	var username string
	username = "Jamal"
	fmt.Printf("Nama saya %s \n", username)

	var value int = 100
	fmt.Printf("Nilai ulangan kemarin adalah %d \n", value)

	// Deklarasi variable tanpa menuliskan tipe data
	// Variable harus langsung diassign value dan tidak bisa diberikan default value kosong
	var floatValue = 10.93
	fmt.Printf("Nilai decimal: %f\n", floatValue)

	// Multiple variable
	var (
		email  string
		height float32
		age    int = 30
	)
	email = "jumadi@google.com"
	height = 175.12

	fmt.Printf("Email saya %s, saya memiliki tinggi %f, dan saya berusia %d \n", email, height, age)

	// One line multiple variable
	// Bisa untuk bermacam2 tipe data
	// Harus di assign value di awal
	var value1, value2, value3 = 100, "Bambang", false
	value2 = "Sumanto"
	fmt.Printf("Saya %s, Umur %d, Gweh: %t \n", value2, value1, value3)

	// Short hand variable
	// ? Deklarasi variable seperti ini hanya bisa dilakukan di dalam function saja
	price := 10000000
	fmt.Printf("Harga barang ini sebesar %d \n", price)

	// [CONSTANT]
	// Variable yang nilainya konstan/tidak bisa diubah
	// Deklarasi tanpa tipe data
	// ? Kaidah penulisan const adalah dengan menggunakan huruf kapital semua
	const MONTH = "January"

	// Deklarasi dengan tipe data
	const phi float32 = 3.14

	// Multiple const variable
	const (
		title  string  = "Naruto"
		rating float64 = 9.5
	)

	const NETWORK_NAME, ADDRESS_PORT = "BSCTestNet", 115511

	// [POINTER]
	// Biasa digunakan untuk menyimpan address/alamat dari suatu variable
	// ! Yang disimpan hanyalah address/alamat, bukan value aslinya

	var myPointer *int
	var myPriceToPoint int = 15000
	myPointer = &myPriceToPoint

	fmt.Println("Address from pointer: ", myPointer) // Jika tidak menggunakan * maka akan menampilkan addressnya saja
	fmt.Println("Value from pointer: ", *myPointer) // Jika menggunakan * akan menampilkan isi value dari variable yang dipoint

	// -- Mengubah value melalui pointer --
	myUserNameToPoint := "dionovan7"
	var usernamePointer *string = &myUserNameToPoint
	*usernamePointer = "kizaruu77" // ! Jangan lupa tambahkan * di pointer jika ingin mengubah value pada variable

	fmt.Println("Pointed username: ", myUserNameToPoint)

	// -- Zero value pada pointer --
	// --- Deklarasi pointer menggunakan * ---
	var myPointer1 *string
	fmt.Printf("Value pointer1: %v \n", myPointer1)
	// fmt.Printf("Value pointer1: %v \n", *myPointer1) // ? Jika tidak ada address, maka akan menyebabkan panic error

	// --- Deklarasi menggunakan new() ---
	var myPointer2 = new(bool) 
	fmt.Printf("Value pointer2: %v \n", myPointer2)
	fmt.Printf("Value pointer2: %v \n", *myPointer2) // ? Jika tidak ada address, maka value dari pointer akan memakai default zero value dari tiap tipe data  (int = 0, string = "", bool = false, etc)

	// [ARITMATHIC OPERATION]
	fmt.Printf("\n")
	fmt.Printf("--- ARITMATHIC --- \n")

	var sum = 10 + 30
	fmt.Println("Penjumlahan: ", sum)

	var subtract = 10 - 2 
	fmt.Println("Pengurangan: ", subtract)

	var perkalian = 20 * 2
	fmt.Println("Perkalian: ", perkalian)

	var pembagian float64 = 10.0 / 9.0
	fmt.Println("Pembagian: ", pembagian)

	var modulus = 20 % 3
	fmt.Println("Modulus: ", modulus)

	var incValue = 10
	incValue++
	fmt.Println("Increment: ", incValue)

	var decrementValue = 10
	decrementValue--
	fmt.Println("Increament: ", decrementValue)

	// -- Multiple Operation --
	// ? Urutan => 1. Di dalam kurung, kali dan bagi, lalu kurang dan tambah
	var myMultipleOperation = 10 + (5 - 2) * 30 / 2
	fmt.Println("Multiple Operation: ", myMultipleOperation)

	// -- Operator Assignment --
	var myValue int = 20
	myValue += 5
	fmt.Println("Plus assignment: ", myValue)

	myValue -= 3
	fmt.Println("Subtract assignment: ", myValue)

	myValue *= 2
	fmt.Println("Perkalian assignment: ", myValue)

	myValue /= 5
	fmt.Println("Pembagian assignment: ", myValue)

	myValue %= 5
	fmt.Println("Modulus assignment: ", myValue)

	// -- Logical Operator --
	var isValid bool = true
	var isCorrect bool = false

	fmt.Println(isValid || isCorrect)
	fmt.Println(isValid && !isCorrect)

	// -- Comparison Operator --
	var myAnotherValue1 int = 50
	var myAnotherValue2 int = 50
	fmt.Println("== ", myAnotherValue1 == myAnotherValue2)
	fmt.Println("!= ", myAnotherValue1 != myAnotherValue2)
	fmt.Println("> ", myAnotherValue1 > myAnotherValue2)
	fmt.Println(">= ", myAnotherValue1 >= myAnotherValue2)
	fmt.Println("<= ", myAnotherValue1 <= myAnotherValue2)
	fmt.Println("< ", myAnotherValue1 < myAnotherValue2)

	// --- Comparison Using String ---
	// ? Untuk pembandingan menggunakan tipe data string sebagai pembanding, yang akan dibandingkan adalah urutan lexical dari ACII
	// ? Urutannya: 0-9 -> A-Z -> a-z
	var comparedString1 string = "AkuSeorangKapiten"
	var comparedString2 string = "mempunyaiPedangPanjang"
	fmt.Println(comparedString1 > comparedString2)
	fmt.Println(comparedString1 < comparedString2)
}
