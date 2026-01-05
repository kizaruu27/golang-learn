package main

import (
	"fmt"
	"strconv"
)

// --- Type Alias ---
// ? Sebuah tipe data yang dibuat dari tipe data yang sudah ada
// ? Biasanya digunakan untuk memperjelas kegunaan dari tipe data tersebut
type Makanan = string
type Age = int

func main() {
	var burger Makanan = "Burger"
	var umurYanto Age = 20
	PrintAll(burger, umurYanto)

	// --- tipe data any ---
	// ? Sebuah tipe data dinamis yang bisa menampung semua jenis tipe data
	arrData := [5]any{"Bambang", 10, 25.5, false, "Yanto"}
	mapData := map[string]any{
		"Sumanto":  1,
		"Yahya":    "Jelek",
		"Samsudin": false,
	}

	// --- tipe data interface {} ---
	mySlice := []interface{}{"Udin", 1, false}

	fmt.Println(arrData, mapData, mySlice)

	// --- Type conversion ---
	// ? Mengconversi suatu tipe data ke tipe data yg lain
	xValue := 65   // int
	yValue := 50.5 // float

	floatCalculation := float64(xValue) / yValue // ? Mengubah int ke float
	intCalculation := xValue * int(yValue)       // ? Mengubah float ke int
	fmt.Println("Hasil float:", floatCalculation, "dan Hasil int:", intCalculation)

	// --- String convertion ---
	// -- string to int ---
	stringValue := "20"
	intValue := 10

	// ? me-return 2 value -> value yg diconvert dan error
	// ? Atoi -> string(A) ke int(i)
	convertedValue, _ := strconv.Atoi(stringValue)
	fmt.Println(convertedValue + intValue)

	// -- int to string --
	// ? Itoa -> int(I) ke string(a)
	// ? hanya me-return 1 value -> value yg diconvert
	convertedString := strconv.Itoa(intValue)
	fmt.Printf("Tipe data dari %s, adalah %T\n", convertedString, convertedString)

	// --------------------------
	calculateNumbers(1, 2, 3, 4)

	// ? Jika tipe data any tidak sesuai dengan yg di assertion, maka default value-nya adalah zero value dari tipe data yang ingin di assert
	var myValue any = "1233"
	assertedValue, ok := myValue.(string)

	if !ok {
		fmt.Println("Data type not match")
	} else {
		fmt.Println(assertedValue)
	}
}

func PrintAll(makanan Makanan, age Age) {
	fmt.Println("Nama makanan:", makanan, "Umur orang ini:", age)
}

// ! Keterbatasan tipe data any
func calculateNumbers(values ...any) {
	totalValue := 0
	for _, n := range values {
		// ! Operasi ini tidak bisa dilakukan karna any tidak bisa melakukan operasi matematika
		// totalValue += n

		// -- Type assertion --
		// ? menandakan tipe data any menjadi tipe data yang kita mau -> int, string, float, boolean, dll
		// ? assertion me-return 2 value -> value yg di assert dan boolean status
		value, valid := n.(int)
		if !valid {
			fmt.Println("Value tidak valid")
		} else {
			totalValue += value
			fmt.Println(totalValue)
		}
	}
}
