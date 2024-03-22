package main

import "fmt"

// struct Mark
type Mark struct {
	Massa, Tinggi float64
}

// struct John
type John struct {
	Massa, Tinggi float64
}

// function menghitung BMI dari Mark dan John
func BMI(massaM, tinggiM, massaJ, tinggiJ float64) (float64, float64) {
	bmiM := massaM / (tinggiM * tinggiM)
	bmiJ := massaJ / (tinggiJ * tinggiJ)

	return bmiM, bmiJ
}

func main() {
	// object dari struct Mark
	mark1 := Mark{78, 1.69}
	mark2 := Mark{95, 1.88}

	// object dari struct John
	john1 := John{92, 1.95}
	john2 := John{85, 1.76}

	// memasukan massa dan tinggi mark ke variable
	massaM1, tinggiM1 := mark1.Massa, mark1.Tinggi
	massaM2, tinggiM2 := mark2.Massa, mark2.Tinggi

	// memasukan massa dan tinggi john ke variable
	massaJ1, tinggiJ1 := john1.Massa, john1.Tinggi
	massaJ2, tinggiJ2 := john2.Massa, john2.Tinggi

	// bmi data 1
	bmiM1, bmiJ1 := BMI(massaM1, tinggiM1, massaJ1, tinggiJ1)

	// bmi data 2
	bmiM2, bmiJ2 := BMI(massaM2, tinggiM2, massaJ2, tinggiJ2)

	// variable boolean untuk mengecek apakah BMI mark lebih tinggi dari pada john
	markHigherBMI := func(bmiM, bmiJ float64) bool {
		if bmiM > bmiJ {
			return true
		} else {
			return false
		}
	}

	// print hasil dari BMI dan apakah BMI Mark lebih tinggi dari pada John
	fmt.Printf("BMI Mark data 1: %.4f dan BMI John data 1: %.4f\n", bmiM1, bmiJ1)
	fmt.Println("Mark memiliki BMI lebih tinggi dari pada John?", markHigherBMI(bmiM1, bmiJ1))

	fmt.Printf("\nBMI Mark data 2: %.4f dan BMI John data 2: %.4f\n", bmiM2, bmiJ2)
	fmt.Println("Mark memiliki BMI lebih tinggi dari pada John?", markHigherBMI(bmiM2, bmiJ2))

}
