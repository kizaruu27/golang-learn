package main

import "fmt"

type Siswa struct {
	name string
	address string
	age int
	score int
}

type Computer struct {
	cpuName string
	core int
	ram int
	storage int
}

// --- Embeded Struct ---
type SchoolResult struct {
	Siswa // ? struct Siswa langsung diembed ke dalam struct SchoolResult
	totalScore int
	pass bool
}

type ComputerResult struct {
	Computer
	totalPrice int
	qty int
}

// -- function dengan struct --
// ? Function ini akan menjadi bagian dari struct receivernya
func (siswa Siswa) calculateScore() int {
	return siswa.score * 100
}

func (comp Computer) calculatePrice() (int, int) {
	return (comp.core + comp.storage + comp.ram) * 100, (comp.core + comp.storage + comp.ram) / 100
}

// ? Function menggunakan pointer receiver -> berguna untuk mengubah value pada field sebuah struct
// ? Jika tidak menggunakan pointer, maka data tidak bisa diubah
func (siswa *Siswa) changeName(name string) {
	siswa.name = name
}

func (comp *Computer) editComputer(cpuName string, core int) {
	comp.cpuName = cpuName
	comp.core = core
}

func main() {
	// ? Jika iniasiasi struct tanpa menggunakan nama field, maka semua value dari field harus diisi dan harus berurutan
	dataSiswa := Siswa{"Samsudin", "Jl. Mangga", 20, 80}
	fmt.Println(dataSiswa.name)
	fmt.Println(dataSiswa.address)
	fmt.Println(dataSiswa.age)
	fmt.Println(dataSiswa.score)

	fmt.Println("--- After change name ---")
	dataSiswa.changeName("Gibran Subianto")
	fmt.Println("Nama setelah dimodifikasi:", dataSiswa.name)

	calculatedScore := dataSiswa.calculateScore()
	fmt.Println("Calculated score:", calculatedScore)

	// ? Jika inisiasi struct menggunakan nama field, maka bisa tidak mengisi semua field dan bisa tidak berurutan -> field yg kosong akan berisi zero value dari tipe data secara default
	pcComponent := Computer{cpuName: "Intel", core: 12, ram: 32, storage: 128}
	fmt.Printf("CPU: %s\nCore: %d\nRAM: %d\nStorage: %d\n", pcComponent.cpuName, pcComponent.core, pcComponent.ram, pcComponent.storage)
	pcComponent.editComputer("AMD", 8)
	
	fmt.Println("--- Comp after modified ---")
	fmt.Printf("CPU: %s\nCore: %d\nRAM: %d\nStorage: %d\n", pcComponent.cpuName, pcComponent.core, pcComponent.ram, pcComponent.storage)

	
	maxPrice, _ := pcComponent.calculatePrice()
	fmt.Println("Maximum pc price:", maxPrice)

	// --- Pointer struct ---
	var pointerSiswa *Siswa = &Siswa{name: "Wahyu", score: 20}
	fmt.Println((*pointerSiswa).name)
	fmt.Println(pointerSiswa.score) // ? Tidak wajib menggunakan * ketika ingin mengakses nilai dari pointer struct

	// --- Anonymus struct ---
	// ? Struct langsung disimpan di dalam variable dan langsung diisi nilainya di tiap field
	tshirt := struct {
		brand string
		color string
	} {
		brand: "Uniqlo",
		color: "red",
	}
	fmt.Println(tshirt.brand, tshirt.color)

	schoolResult := SchoolResult{totalScore: 85, pass: true, Siswa: Siswa{name: "Prabowo", address: "Jl. Sawit", age: 19, score: 75}}
	fmt.Println(schoolResult.Siswa.name)

	computerResult := ComputerResult{totalPrice: 2000000, qty: 10, Computer: pcComponent}
	fmt.Println(computerResult)

	// -- Array of structs ---
	computerArray := []Computer {
		{cpuName: "AMD", core: 12, ram: 32},
		{cpuName: "Intel", core: 8, ram: 16},
	}
	
	for _, comp := range computerArray {
		fmt.Printf("CPU: %s, core: %d, RAM: %d \n", comp.cpuName, comp.core, comp.ram)
	}
}
