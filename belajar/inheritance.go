package belajar

import (
	"fmt"
)

// inheritance
type WorldCup struct {
	Orgnizer      string
	Country       string
	CountryPlayer int
}

func (g *WorldCup) StartGame() {
	fmt.Printf("Telah selesai piala dunia yang diselenggarakan oleh %s di negara %s yang diikuti oleh %d negara.\n",
		g.Orgnizer, g.Country, g.CountryPlayer)
}

type TournamentWorldCup struct {
	WorldCup
	Winner string
	Prizes string
}

func (t *TournamentWorldCup) StartTournament() {
	fmt.Printf("Piala dunia yang diselenggarakan oleh %s di negara %s yang diikuti oleh %d negara telah dimenangkan oleh %s dengan total hadiah %s.\n",
		t.Orgnizer, t.Country, t.CountryPlayer, t.Winner, t.Prizes)
}
