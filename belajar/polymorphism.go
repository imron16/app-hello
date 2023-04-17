package belajar

import "fmt"

//polymorphism

func Menghitung(hitung string) {
	fmt.Println("Mari Belajar Menghitung : ", hitung)
}

type KotakService interface {
	GetBangunan() string
}

type kotak struct {
	bangun  string
	panjang int
	lebar   int
	tinggi  int
}

func (k kotak) getKotak() string {
	return k.bangun
}

func (k kotak) GetBangunan() string {
	return fmt.Sprintf("Menghitung : %s, Panjang nya : %d, lebar nya : %d, tinggi nya : %d",
		k.getKotak(), k.panjang, k.lebar, k.tinggi)
}

func NewKotak(x string, a, b, c int) KotakService {
	return kotak{bangun: x, panjang: a, lebar: b, tinggi: c}
}
