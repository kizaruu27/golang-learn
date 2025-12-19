package main

import "fmt"

func main() {
	const appName string = "JajaninApp"
	const version float64 = 1.0
	const releaseYear int = 2022

	// ? %.1f artinya banyaknya value yang ditampilkan di belakang koma
	fmt.Printf("Nama aplikasi ini adalah %s, saat ini berada di versi %.1f. App ini dirilis pada tahun %d lalu\n", appName, version, releaseYear)
}