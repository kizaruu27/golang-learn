package main

import "fmt"

func main() {
	// [CONDITIONAL STATEMENT]
	// -- if statement --
	x := 50
	y := 30

	if x == y {
		fmt.Println("X sama dengan Y")
	} else if x > y {
		fmt.Println("X lebih dari Y")
	} else if x < y {
		fmt.Println("X kurang dari Y")
	} else {
		fmt.Println("Nilai tidak valid")
	}

	a := "Pria"
	b := "Wanita"
	if a == "Pria" && b == "Wanita" {
		fmt.Println("Gender sudah sesuai")
	} else {
		fmt.Println("Gender tidak sesuai")
	}

	currentBalance := 10000
	minimumBalance := 50000
	if currentBalance > minimumBalance {
		fmt.Println("Transaksi berhasil")
	} else {
		fmt.Println("Saldo tidak mencukupi")
	}

	// -- Deklarasi variable di if statement --
	cekSaldo := func() int {
		return 1000000
	}

	// ? Variable currentSaldo dapat diakses di dalam block if, else if, maupun else semala masih dalam scope if statement yang sama
	// ? Pendeklarasian seperti membuat variable biasa yang ditulis di sebelah if lalu diakhiri dengan ";"
	// ! Pendeklarasian variable di if statement harus menggunakan shorthand variable, dan tidak bisa menggunakan pendeklarasian variable lainnya !
	if currentSaldo := cekSaldo(); currentSaldo > minimumBalance {
		fmt.Printf("Transaksi berhasil, sisa saldo sekarang: %d\n", currentSaldo-100000)
	} else {
		fmt.Printf("Saldo tidak mencukupi, saldo saat ini: %d\n", currentSaldo)
	}

	if kelamin, nama := "Bisexual", "Yanto"; kelamin == "Pria" {
		fmt.Printf("%s adalah seorang %s \n", nama, kelamin)
	} else if kelamin == "Wanita" {
		fmt.Printf("Lohh kelamin %s %s? \n", nama, kelamin)
	} else {
		fmt.Printf("Kelamin %s tidak jelas lohh yaa \n", nama)
	}

	// -- Switch Case --
	var menuOption int = 2

	switch menuOption {
		case 0:
			fmt.Println("Opsi menu 0")
		case 1:
			fmt.Println("Opsi menu 1")
		case 2:
			fmt.Println("Opsi menu 2")
		case 3:
			fmt.Println("Opsi menu 3")
		default:
			fmt.Println("Menu tidak ditemukan")
	}

	// --- Switch multiple case ---
	var namaKota string = "Ciawi"
	switch namaKota {
		case "Bandung", "Bogor", "Depok", "Cimahi": // ? Implementasi seperti penggunaan OR di if statement -> jika salah satu terpenuhi maka akan dijalankan
			fmt.Println("Jawa Barat")
		case "Jakarta":
			fmt.Println("DKI Jakarta")
		default:
			fmt.Println("Kota belum terdaftar")
	}

	// --- Switch case tanpa nilai ---
	// ? Biasanya digunakan jika ingin melakukan comparison/pembandingan pada tiap kondisi di switch case
	score := 50
	switch {
		case score == 100:
			fmt.Println("Nilai Sempurna")
		case score >= 80:
			fmt.Println("Nilai Bagus")
		case score >= 70:
			fmt.Println("Cukup")
		case score < 70:
			fmt.Println("Kurang")
	}

	// --- Switch case with fallthrough ---
	// ? Akan melanjutkan ke case selanjutnya dan tidak akan mengecek kondisi dari case tersebut
	// ? fallthrough pertama akan dieksekusi pada block yang true terlebih dahulu dan akan stop di case/block yang tidak memiliki fallthrough
	fmt.Println("--- Using Fallthrough ---")
	switch {
	case score == 100:
		fmt.Println("Score sempurna")
		fallthrough
	case score == 50:
		fmt.Println("Score 50")
		fallthrough
	case score == 20:
		fmt.Println("Score 20")
	}
}
