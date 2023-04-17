package main

import (
	"fmt"

	"github.com/imron16/app-hello/belajar"
)

var pan, le, ti, x int
var fruits = [4]string{"Pisang", "Apel", "Durian", "Alpukat"}
var fruity = make([]string, 2)
var fruitsA = []string{"Pisang", "Apel", "Durian", "Alpukat"}
var fruitsB = append(fruitsA, "Jeruk", "Pir")
var newFruits = fruitsA[0:2]
var newFruitsA = fruitsA[0:2:2]

func main() {
	//inheritance
	// Create a world cup
	worldCup := belajar.WorldCup{Orgnizer: "FIFA", Country: "Qatar", CountryPlayer: 32}
	worldCup.StartGame()

	// Create a tournament world cup
	tournamentWorldCup := belajar.TournamentWorldCup{WorldCup: belajar.WorldCup{Orgnizer: "FIFA", Country: "QATAR", CountryPlayer: 4}, Winner: "Argentina", Prizes: "$42,000,000"}
	tournamentWorldCup.StartTournament()

	fmt.Println("")

	//encapsulation
	user := belajar.NewUser("Imron", 26, "IT", "DANAMAS")
	fmt.Println(user.GetUser())

	fmt.Println("")

	//polymorphism
	belajar.Menghitung("Kotak")

	fmt.Printf("Masukan Panjang	: ")
	fmt.Scan(&pan)
	fmt.Printf("Masukan Lebar	: ")
	fmt.Scan(&le)
	fmt.Printf("Masukan tinggi : ")
	fmt.Scan(&ti)

	kotak := belajar.NewKotak("Kotak", pan, le, ti)
	fmt.Println(kotak.GetBangunan())

	x = pan * le * ti

	fmt.Println("Hasil Volume nya adalah 	: ", x)

	if x >= 1000 {
		fmt.Println("Volume Kotak Besar")
	} else if x < 1000 && x >= 500 {
		fmt.Println("Volume Kotak Sedang")
	} else {
		fmt.Println("Volume Kotak Kecil")
	}

	for _, fruit := range fruits {
		fmt.Printf("Ini Buah %s. \n", fruit)
	}

	fruity[0] = "apel"
	fruity[1] = "mangga"

	fmt.Println(fruity)

	fmt.Println(fruitsA[1])

	fmt.Println(fruits)

	fmt.Println(newFruits)

	fmt.Println(len(newFruits))

	fmt.Println(cap(newFruits))

	fruitsA[1] = "Jeruk"

	fmt.Println(fruitsA)

	fmt.Println(fruitsB)

}
