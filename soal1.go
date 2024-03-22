package main

import "fmt"

// struct lumba-lumba
type Lumba_lumba struct {
	skor_lumba, skorLBonus1, skorLBonus2 []int
}

// struct koala
type Koala struct {
	skor_koala, skorKBonus1, skorKBonus2 []int
}

// function mencari skor rata2 dari lumba2 dan koala
func meanSkor(skorL, skorK []int) (int, int) {
	var sumL int
	var sumK int

	// mengambil value dari data lumba2
	for _, skor := range skorL {
		sumL += skor
	}

	// mengambil value dari data koala
	for _, skor := range skorK {
		sumK += skor
	}

	return sumL / 3, sumK / 3 // rata2 adalah total skor / banyaknya pertandingan (3)
}

// cek hasil pertandingan
func hasilLomba(meanL, meanK int) string {
	if meanL < 100 && meanK < 100 {
		return "Tidak ada yang memenangkan trofi"
	} else if (meanL >= 100 && meanK >= 100) && (meanL == meanK) {
		return "Seri"
	} else {
		if meanL > meanK {
			return "Lumba-lumba memenangkan trofi!"
		} else {
			return "Koala memenangkan trofi!"
		}
	}
}

func main() {
	// memasukan data lumba2
	dataL1 := []int{96, 108, 89}
	dataLBonus1 := []int{97, 112, 101}
	dataLBonus2 := []int{97, 112, 101}

	// memasukan data koala
	dataK1 := []int{88, 91, 110}
	dataKBonus1 := []int{109, 95, 123}
	dataKBonus2 := []int{109, 95, 106}

	// membuat object dari struct lumba-lumba dan koala
	lumba_lumba := Lumba_lumba{dataL1, dataLBonus1, dataLBonus2}
	koala := Koala{dataK1, dataKBonus1, dataKBonus2}

	// memanggil function meanSkor untuk mengetahui setiap mean dari setiap data
	meanL, meanK := meanSkor(lumba_lumba.skor_lumba, koala.skor_koala)
	meanLBonus1, meanKBonus1 := meanSkor(lumba_lumba.skorLBonus1, koala.skorKBonus1)
	meanLBonus2, meanKBonus2 := meanSkor(lumba_lumba.skorLBonus2, koala.skorKBonus2)

	// memanggil function hasilLomba untuk tau hasil dari siapa yang menang
	hasil1 := hasilLomba(meanL, meanK)
	hasilBonus1 := hasilLomba(meanLBonus1, meanKBonus1)
	hasilBonus2 := hasilLomba(meanLBonus2, meanKBonus2)

	// hasil dari Data 1
	fmt.Printf("Skor lumba-lumba : %d dan skor koala %d\n", meanL, meanK)
	fmt.Println("Hasil Data 1		: ", hasil1)

	// hasil dari Data Bonus 1
	fmt.Printf("\nSkor lumba-lumba : %d dan skor koala %d\n", meanLBonus1, meanKBonus1)
	fmt.Println("Hasil Data Bonus 1	: ", hasilBonus1)

	// hasil dari Data Bonus 2
	fmt.Printf("\nSkor lumba-lumba : %d dan skor koala %d\n", meanLBonus2, meanKBonus2)
	fmt.Println("Hasil Data Bonus 2	: ", hasilBonus2)
}
