package main

import "fmt"

// -- Interface --
// ? Sebuah tipe data yang mendefinisikan bentuk function kosong
type BangunRuang interface {
	luas() float64
	// ? Jika function ini ditambahkan ke interface, maka function ini wajib diimplementasikan pada struct
	// ! Jika tidak diimplementasikan, maka struct akan dianggap tidak mengimplementasikan interface tersebut
	// ! Sehingga tiap ada perubahan pada interface, maka struct juga harus berubah
	// keliling() float64
}

// ? Solusi agar struct tidak harus selalu berubah, kita perlu menggunakan interface dengan function sesedikit mungkin untuk meminimalisir perubahan pada struct
type Keliling interface {
	keliling() float64
}

type Persegi struct {
	sisi float64
}

func (p Persegi) keliling() float64 {
	return p.sisi * p.sisi * p.sisi
}

func hitungKeliling(keliling Keliling) {
	fmt.Println(keliling.keliling())
}

// ? Struct yang memiliki format function yang sama, akan dianggap mengimplementasikan function dari suatu interface
func (persegi Persegi) luas() float64 {
	return persegi.sisi * persegi.sisi
}

// ? interface bisa dijadikan parameter untuk memanggil function milik interface
func hitungLuasPersegi(bangunRuang BangunRuang) float64 {
	return bangunRuang.luas()
}

// --------------------------
type RestaurantMenu interface {
	price() int
	menuType() string
}

type Food struct {
	foodName    string
	basePrice   int
	composition string
}

func (f Food) price() int {
	return f.basePrice * 14 / 100
}

func (f Food) menuType() string {
	return fmt.Sprintf("Makanan ini mengandung %s", f.composition)
}

type Drink struct {
	drinkName string
	basePrice int
	iced      bool
}

func (d Drink) price() int {
	return d.basePrice * 4 / 100
}

func (d Drink) menuType() string {
	if d.iced {
		return "Minuman ini dingin"
	} else {
		return "Minuman ini panas"
	}
}

func restoPrice(restoMenu RestaurantMenu) {
	price := restoMenu.price()
	fmt.Println("Price:", price)
}

func checkMenu(restoMenu RestaurantMenu) {
	fmt.Println(restoMenu.menuType())
}

func main() {
	persegiPanjang := Persegi{sisi: 10}

	// ? Kita bisa memasukkan struct ke parameter function meskipun parameternya adalah interfacenya
	// ! Ini hanya bisa dilakukan jika struct tersebut mengimplementasikan function dari interface
	luasPersegi := hitungLuasPersegi(persegiPanjang)

	fmt.Println("Luas persegi panjang:", luasPersegi)
	hitungKeliling(persegiPanjang)

	food := Food{foodName: "Bakso", basePrice: 10000, composition: "Daging"}
	drink := Drink{drinkName: "Es teh manis", basePrice: 2000, iced: true}

	restoPrice(food)
	restoPrice(drink)

	checkMenu(food)
	checkMenu(drink)

}
