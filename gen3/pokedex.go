package gen3

import (
	"github.com/anaminus/pkm"
)

var (
	structDex = makeStruct(
		2, // 0 Species
	)
)

type pokedexData struct {
	Name    string
	Size    int
	Address ptr
}

type Pokedex struct {
	v *Version
	i int
}

func (p Pokedex) Name() string {
	return p.v.pokedex[p.i].Name
}

func (p Pokedex) Size() int {
	return p.v.pokedex[p.i].Size
}

func (p Pokedex) Species(number int) pkm.Species {
	if number <= 0 || number > p.Size() {
		panic("species number out of bounds")
	}
	p.v.ROM.Seek(p.v.pokedex[p.i].Address.ROM(), 0)
	var species pkm.Species
	for i, q := 1, make([]byte, 2); i <= indexSizeSpecies; i++ {
		p.v.ROM.Read(q)
		if int(decUint16(q)) == number {
			species = Species{v: p.v, i: i}
			break
		}
	}
	return species
}

func (p Pokedex) AllSpecies() []pkm.Species {
	a := make([]pkm.Species, p.Size())
	p.v.ROM.Seek(p.v.pokedex[p.i].Address.ROM(), 0)
	for i, q := 1, make([]byte, 2); i < p.Size(); i++ {
		p.v.ROM.Read(q)
		if n := int(decUint16(q)); n <= p.Size() {
			a[n-1] = Species{v: p.v, i: i}
		}
	}
	return a
}

func (p Pokedex) SpeciesNumber(species pkm.Species) int {
	b := readStruct(
		p.v.ROM,
		p.v.pokedex[p.i].Address,
		species.Index()-1,
		structDex,
	)
	return int(decUint16(b))
}
