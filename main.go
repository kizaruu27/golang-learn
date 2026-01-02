package main

import "fmt"

func main() {
	// -- map[] --
	// ? seperti slice maupun array, namun map menggunakan "key" sebagai penanda tiap value di dalamnya -> slice/array menggunakan index
	var myHobby map[string] string = map[string] string {
		"Bambang": "Nyabu",
		"Wahyu": "Ngudud",
		"Suyatmo": "Berantem",
	}
	fmt.Println(myHobby["Suyatmo"])

	// -- membuat map[] menggunakan make --
	// ? Jika mencoba mengakses key yang tidak ada, maka yang tampil adalah zero value dari tipe data map tersebut
	var score map[string] int = make(map[string] int)
	score["Wahyu"] = 100
	score["Bambang"] = 70
	score["Jose"] = 85

	fmt.Println(score["Ayu"])

	// -- Mengecek keberadaan value dari map --
	// ? value dari map sendiri me-return 2 jenis value
	// ? value pertama merupakan nilai dari key tersebut, dan value kedua merupakan boolean ada atau tidaknya key tersebut
	namaSiswa := "Bahlil"
	if score, isAvailable := score[namaSiswa]; isAvailable {
		fmt.Println("Nilai wahyu adalah:", score)
	} else {
		fmt.Printf("Nama %s tidak ditemukan \n", namaSiswa)
	}

	// -- Iterasi menggunakan map --
	fmt.Printf("--- Iterasi menggunakan map --- \n")
	for key, value := range score {
		fmt.Printf("%s mendapatkan score %d \n", key, value)
	}
}
