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

func main() {
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
