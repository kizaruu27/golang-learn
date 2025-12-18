package main

import "fmt"

// [FUNCTION]
// -- Function Init --
// ? Function yang akan dipanggil sebelum menjalankan function main() / entry point
// ? Jika terdapat lebih dari 1 init, maka akan dieksekusi secara berurutan
func init() {
	fmt.Println("Hello, from init() 1")
}

func init() {
	fmt.Println("Hello, from init() 2")
}

// -- Function main() --
// ? Function yang merupakan entry point ketika program dijalankan
func main() {
	printHello()
	sumValue(10, 50)

	// Memanggil return function
	generatedMessage := generateMessage()
	fmt.Println(generatedMessage)

	// Memanggil return function dengan parameter
	speed := calculateSpeed(4.0, 9.8, 320.0)
	fmt.Println(speed)

	userMessage := getUsername("kizaruu77")
	fmt.Println(userMessage)

	// Memanggil func dengan multiple return
	sum, subtract, times, divide := calculateValue(32, 10, 5.0, 3.5)
	fmt.Println("Penjumlahan: ", sum)
	fmt.Println("Pengurangan: ", subtract)
	fmt.Println("Perkalian: ", times)
	fmt.Println("Pembagian: ", divide)

	// -- Blank identifier --
	// ? Digunakan jika hanya ingin memanggil return value yang dibutuhkan saja, yang tidak dibutuhkan bisa ditulis dengan _
	usedSum, _, _, _ := calculateValue(10, 10, 0, 0)
	fmt.Printf("Value yang digunakan %d \n", usedSum)

	_, message := generateGreeting("Bahlil")
	fmt.Println(message)
}

// -- Function Declaration --
// ! func tidak bisa dideklarasikan di dalam func lain !
func printHello() {
	fmt.Println("Hello, Duniah")
}

// -- Function with Parameter --
func sumValue(value1 int, value2 int) {
	sum := value1 + value2
	fmt.Println("Sum value: ", sum)
}

// -- Return Function --
func generateMessage() string {
	return "Hai, selamat datang di konoha"
}

// -- Return Function with Parameter --
func calculateSpeed(velocity float64, gravity float64, time float64) string {
	velocityTime := velocity * time
	speed := velocityTime / gravity

	return fmt.Sprintf("Total speed is: %f", speed)
}

func getUsername(username string) string {
	name := fmt.Sprintf("Welcome back, %s", username)

	return name
}

//  -- Multiple Return Function --
func calculateValue(value1 int, value2 int, value3 float64, value4 float64) (int, int, float64, float64) {
	sum := value1 + value2
	subtract := value1 - value2
	times := value3 * value4
	divide := value3 / value4

	return sum, subtract, times, divide
}

// -- Named Return Function --
// ? Memberikan nama pada return type
// ? Jika return kosong, maka yang direturn adalah type yang sudah diberi nama
// ? Jika menambahkan value lagi setelah return, maka type yang diberi nama akan ter-override
func generateGreeting(name string) (greet string, message string) {
	greet = fmt.Sprintf("Good morning, %s", name)
	message = fmt.Sprintf("Semoga harimu menyenangkan yaa %s sayang", name)
	return 
}
