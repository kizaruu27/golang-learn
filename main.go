package main

import (
	"errors"
	"fmt"
)
var errZero error = errors.New("value can't be devided with 0")
func devideValue(a, b int) (int, error) {
	if b == 0 {
		// [Error Handling]
		// ? Gunakan errors.New() untuk mengeset error message
		// return 0, errors.New("value can't be devided with 0")

		// ? Return error sebagai variable
		return 0, errZero
	} else if b < 0 {
		return 0, errors.New("can't accept value that less than 0")
	}

	total := a / b
	return total, nil
}

func configUsername(username string) string {
	if username == "" {
		// -- Panic Error --
		// ? Sebuah error event yang akan menghentikan program jika dipanggil atau terjadi error
		// ? Ketika panic dipanggil, maka program di bawahnya tidak akan dieksekusi
		panic("username can't be an empty string")
	}

	return fmt.Sprintf("Selamat pagi, %s", username)
}

func main() {
	// -- Recover panic error menggunakan defer func --
	// ? fungsi defer digunakan untuk menandai sebuah func yang akan dieksekusi terakhir
	// ? Bisa juga seperti callback dimana dia akan menunggu semua function dieksekusi, baru dia akan dieksekusi paling akhir
	// ? Meskipun dia dituliskan paling awal, dia akan tetap dipanggil terakhir
	defer func ()  {
		// ? recover dapat memungkinkan program untuk dihentikan secara halus dan tidak secara paksa jika terjadi panic
		// ? recover akan menangkap error panic message -> mirip seperti catch pada try catch
		r := recover()
		if r != nil {
			fmt.Println("Error recovered:", r)
		}
	}()

	// --- Urutan pemanggilan defer ---
	// ? defer akan memanggil yang paling bawah terlebih dahulu, baru naik ke atas
	defer fmt.Println("Defer 1") // ? Dijalankan ketiga
	defer fmt.Println("Defer 2") // ? Dijalankan kedua
	defer fmt.Println("Defer 3") // ? Dijalankan pertama

	// ? Variable ini menggunakan panic error
	// ? Jika error, maka program di bawahnya tidak akan dieksekusi
	greet := configUsername("Yanto")
	fmt.Println(greet)


	value, err := devideValue(20, 0)

	// ? Perlu mengecek apakah error yang direturn ada atau tidak
	if err != nil {
		// ? Jika ada, panggil fungsi err.Error() untuk menampilkan pesan error yang direturn berupa "string"
		fmt.Printf("There is error in your code: %s \n", err)

		// ? Gunakan errors.Is() untuk membandingkan error message terentu yang ingin dibandingkan
		if errors.Is(err, errZero) {
			fmt.Println("devided with 0 will cause an infinite value")
		}
	} else {
		fmt.Printf("Hasil perhitungan: %d \n", value)
	}


	errGender := errors.New("orang ini bisexual")
	var checkGender = func(gender string) (string, error) {
		switch gender {
		case "Pria", "Wanita", "Laki-laki", "Perempuan":
			return "Gender Valid!", nil
		case "Bisex":
			return "", errGender
		default:
			return "", errors.New("gender tidak valid")
		}
	}

	message, genderErr := checkGender("Bisex")
	if genderErr != nil {
		if errors.Is(genderErr, errGender) {
			fmt.Println("Dihh, dasar kamoeh kaum pelangi")
		}
		fmt.Printf("Error: %s \n", genderErr.Error())
		} else {
		fmt.Println(message)
	}
}
