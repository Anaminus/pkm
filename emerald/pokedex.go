package emerald

import (
	"github.com/anaminus/pkm"
)

const (
	addrStdDex  = 0x0031D94C
	addrNatlDex = 0x0031DC82
)

const (
	pokedexStdSize  = 202
	pokedexNatlSize = 386
)

type Pokedex struct {
	v Version
	i int
}

func (p Pokedex) Name() string {
	// TODO
	return ""
}

func (p Pokedex) Size() int {
	// TODO
	return 0
}

func (p Pokedex) Species(number int) pkm.Species {
	// TODO
	return nil
}

func (p Pokedex) AllSpecies() []pkm.Species {
	// TODO
	return nil
}

func (p Pokedex) SpeciesNumber(species pkm.Species) int {
	// TODO
	return 0
}
