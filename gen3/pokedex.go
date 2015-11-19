package gen3

import (
	"github.com/anaminus/pkm"
)

var (
	structDex = makeStruct(
		2, // 0 Species
	)
)

type PokedexNatl struct {
	v Version
}

func (p PokedexNatl) Name() string {
	return "National"
}

func (p PokedexNatl) Size() int {
	return 386
}

func (p PokedexNatl) Species(number int) pkm.Species {
	if number <= 0 || number > p.Size() {
		panic("species number out of bounds")
	}
	p.v.ROM.Seek(int64(p.v.AddrPokedexNatl), 0)
	for i, q := 1, make([]byte, 2); i < indexSizeSpecies; i++ {
		p.v.ROM.Read(q)
		if int(decUint16(q)) == number {
			return Species{v: p.v, i: i}
		}
	}
	return nil
}

func (p PokedexNatl) AllSpecies() []pkm.Species {
	a := make([]pkm.Species, p.Size())
	p.v.ROM.Seek(int64(p.v.AddrPokedexStd), 0)
	for i, q := 1, make([]byte, 2); i < indexSizeSpecies; i++ {
		p.v.ROM.Read(q)
		if n := int(decUint16(q)); n <= p.Size() {
			a[n-1] = Species{v: p.v, i: i}
		}
	}
	return a
}

func (p PokedexNatl) SpeciesNumber(species pkm.Species) int {
	b := readStruct(
		p.v.ROM,
		p.v.AddrPokedexNatl,
		species.Index(),
		structDex,
	)
	return int(decUint16(b))
}

type PokedexStd struct {
	v Version
	i int
}

func (p PokedexStd) Name() string {
	return "Standard"
}

func (p PokedexStd) Size() int {
	return 202
}

func (p PokedexStd) Species(number int) pkm.Species {
	if number <= 0 || number > p.Size() {
		panic("species number out of bounds")
	}
	p.v.ROM.Seek(int64(p.v.AddrPokedexStd), 0)
	for i, q := 1, make([]byte, 2); i < indexSizeSpecies; i++ {
		p.v.ROM.Read(q)
		if int(decUint16(q)) == number {
			return Species{v: p.v, i: i}
		}
	}
	return nil
}

func (p PokedexStd) AllSpecies() []pkm.Species {
	a := make([]pkm.Species, p.Size())
	p.v.ROM.Seek(int64(p.v.AddrPokedexStd), 0)
	for i, q := 1, make([]byte, 2); i < indexSizeSpecies; i++ {
		p.v.ROM.Read(q)
		if n := int(decUint16(q)); n <= p.Size() {
			a[n-1] = Species{v: p.v, i: i}
		}
	}
	return a
}

func (p PokedexStd) SpeciesNumber(species pkm.Species) int {
	b := readStruct(
		p.v.ROM,
		p.v.AddrPokedexStd,
		species.Index(),
		structDex,
	)
	return int(decUint16(b))
}
