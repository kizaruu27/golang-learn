package main

import "fmt"

func main() {
	// --- Slice ---
	var mySlice []string = []string {"Value 1", "Value 2", "Value 3"}
	fmt.Println("Nilai mySlice:", mySlice)
	fmt.Println("len() mySlice:", len(mySlice)) // ? Untuk mendapatkan length dari slice
	fmt.Println("cap() mySlice:", cap(mySlice)) // ? Untuk mendapatkan capacity dari slice

	// --- Slice dengan menggunakan make ---
	// ? make membutuhkan 3 parameter pada umumnya => ([]tipeData, length, capacity)
	// ? make juga bisa diberikan 2 parameter, nanti parameter ke-2 akan otomatis menjadi length dan capacity-nya

	// ---- make() menggunakan 3 parameter
	var names []string = make([]string, 5, 7)
	names[0] = "Sumanto"
	names[1] = "Bambang"
	names[2] = "Supriyanto"
	names[3] = "Samsudin"
	fmt.Println("Names:", names)
	fmt.Println("len() Names:", len(names))
	fmt.Println("cap() Names:", cap(names))

	// ---- make() menggunakan 2 parameter 
	myNumSlice := make([]int, 10) // jika valuenya blm diinisiasi, default value akan menggunakan zero value berdasarkan tipe data
	fmt.Println("myNumSlice:", myNumSlice)
	fmt.Println("len() myNumSlice:", len(myNumSlice))
	fmt.Println("cap() myNumSlice:", cap(myNumSlice))

	// --- Zero value di slice ---
	myZeroSlice := []string {}
	fmt.Println("len() myZeroSlice:", len(myZeroSlice))
	// fmt.Println(myZeroSlice[0])
	myZeroSlice = append(myZeroSlice, "slice1", "slice2", "slice3")
	fmt.Println("myZeroSlice:", myZeroSlice)

	// ---- Zero value slice menggunakan make()
	usernames := make([]string, 0)
	fmt.Println("len() usernames:", len(usernames))
	// fmt.Println(usernames[0]) // ? Jika mengakses slice yang zero value, maka akan terjadi error karna panjangnya kosong

	// --- Menambahkan value pada slice menggunakan append() ---
	usernames = append(usernames, "kizaruu88", "bambang123", "supriyanto") // ? append harus dipanggil sebagai value (variable = append(variable, ..., ...))
	fmt.Println("usernames:", usernames)

	// --- Menambahkan slice ke slice lain menggunakan append() ---
	myString1 := []string {"bambang", "sumanto", "yanto"}
	myString2 := []string {"rahmat", "subagja"}

	myString1 = append(myString1, myString2...) // ? gunakan elips ketika menambahkan slice lain
	fmt.Println("myString:", myString1)

	// ---- Membuat slice baru menggunakan append()
	myNums := []int {10, 30, 40}
	myUpdatedNums := append(myNums, 20, 50, 100) // ! Value tambahan harus memiliki tipe data yang sama
	fmt.Println("myUpdatedNums:", myUpdatedNums)

	// --- Memotong / Menghapus slice ---
	warni := []string {"merah", "kuning", "hijau", "ungu"}
	// ? angka pertama adalah mau potong dari index ke berapa
	// ? angka kedua adalah mau potong sampai index ke berapa
	// ! value pada index terakhir tidak akan ter-include
	
	warni = warni[0:3]
	fmt.Println(warni)  
	
	// ! hati-hati karna di sini memotong slice yang sudah dipotong
	warni = warni[1:2]
	fmt.Println(warni)  

	// ---- Memotong slice dengan hanya 1 value
	namaNama := []string {"Ujang", "Supriyanto", "Prabowo", "Anies"}
	// ? jika : di depan, maka akan mengambil dari index paling awal
	fmt.Println(namaNama[:2])

	// ? jika : di belakang, maka akan mengambil sampai index paling akhir
	fmt.Println(namaNama[2:])

	// ? jika hanya menuliskan : saja, maka akan mengambil seluruh nilai dari slice
	fmt.Println(namaNama[:])

	// --- Memotong / menghapus nilai yang ada ditengah pada slice
	// ? Menggunakan append jika ingin menghapus nilai yang ada di tengah slice
	namaNama = append(namaNama[0:1], namaNama[2:]...)
	fmt.Println(namaNama)

	// --- Membuat slice dari array ---
	// ? Sebenarnya slice itu adalah sebuah pointer dari sebuah array
	arrValues := [3]string {"Yanto", "Yanti", "Bambang"}
	sliceValues := arrValues[:]

	// ? jika mengubah value pada slice yang diberi nilai array yg sudah ada, maka nilai array juga akan ikut berubah
	// ? konsepnya sama seperti pointer
	arrValues[1] = "Sumanto" 

	fmt.Println("sliceValues: ", sliceValues)
	fmt.Println("arrValues: ", arrValues)

	// --- Realokasi slice ---
	arrNames := [3]string {"Budi", "Yanto", "Rusdi"}
	sliceNames := arrNames[:]
	fmt.Println("cap() sliceNames before:", cap(sliceNames))
	
	sliceNames = append(sliceNames, "Susan", "Rahmat") // ? ketika sudah di re-alokasi menggunakan append(), maka array sudah tidak terhubung lagi dengan slice
	fmt.Println("cap() sliceNames after:", cap(sliceNames))

	fmt.Println("arrNames:", arrNames) // ? array tidak ikut bertambah meskipun value slice bertambah
	fmt.Println("sliceNames:", sliceNames)

	// ? Capacity akan bertambah secara eksponensial ketika ada penambahan value pada slice
	// ? Capacity tidak akan berkurang jika ada pengurangan pada value pada slice
	// ! Sebaiknya capacity sudah ditentukan sejak awal
	fmt.Println("")
	fmt.Println("--- Realokasi slice capacity ---")
	warnaWarni := make([]string, 10, 16)
	warnaWarni = append(warnaWarni, "Merah") // ? append juga menambah length dari slice
	fmt.Println("--- Merah ---")
	fmt.Println("len() warna:", len(warnaWarni))
	fmt.Println("cap() warna:", cap(warnaWarni))

	warnaWarni = append(warnaWarni, "Biru")
	fmt.Println("--- Biru ---")
	fmt.Println("len() warna:", len(warnaWarni))
	fmt.Println("cap() warna:", cap(warnaWarni))

	warnaWarni = append(warnaWarni, "Kuning")
	fmt.Println("--- Kuning ---")
	fmt.Println("len() warna:", len(warnaWarni))
	fmt.Println("cap() warna:", cap(warnaWarni))

	warnaWarni = append(warnaWarni, "Jingga")
	fmt.Println("--- Jingga ---")
	fmt.Println("len() warna:", len(warnaWarni))
	fmt.Println("cap() warna:", cap(warnaWarni))

	warnaWarni = append(warnaWarni, "Marun")
	fmt.Println("--- Marun ---")
	fmt.Println("len() warna:", len(warnaWarni))
	fmt.Println("cap() warna:", cap(warnaWarni))
}
