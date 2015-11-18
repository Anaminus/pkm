package emerald

import (
	"bytes"
	"github.com/anaminus/pkm"
	"io"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////

type Version struct {
	ROM io.ReadSeeker
}

var _ = pkm.Version(&Version{})

func OpenROM(rom io.ReadSeeker) pkm.Version {
	return &Version{
		ROM: rom,
	}
}

func (v Version) Name() string {
	return "Emerald"
}

func (v Version) Query() pkm.Query {
	return &Query{v: v}
}

func (v Version) Codecs() []pkm.Codec {
	return []pkm.Codec{
		codecUTF8,
		codecASCII,
		codecString,
		codecPUA,
	}
}
func (v Version) DefaultCodec() pkm.Codec {
	return defaultCodec
}

func (v Version) SpeciesIndexSize() int {
	return indexSizeSpecies
}

func (v Version) SpeciesByIndex(index int) pkm.Species {
	if index < 0 || index >= indexSizeSpecies {
		panic("species index out of bounds")
	}
	return Species{v: v, i: index}
}

func (v Version) SpeciesByName(name string) pkm.Species {
	encName := encodeText(strings.ToUpper(name))
	b := make([]byte, speciesNameLength)
	v.ROM.Seek(addrSpeciesName, 0)
	for i := 0; i < indexSizeSpecies; i++ {
		v.ROM.Read(b)
		if bytes.Equal(encName, truncateText(b)) {
			return Species{v: v, i: i}
		}
	}
	return nil
}

func (v Version) Pokedex() []pkm.Pokedex {
	// TODO
	return nil
}

func (v Version) PokedexByName(name string) pkm.Pokedex {
	// TODO
	return nil
}

func (v Version) ItemIndexSize() int {
	return indexSizeItem
}

func (v Version) Items() []pkm.Item {
	a := make([]pkm.Item, indexSizeItem)
	for i := range a {
		a[i] = Item{v: v, i: i}
	}
	return a
}

func (v Version) ItemByIndex(index int) pkm.Item {
	if index < 0 || index >= indexSizeItem {
		panic("item index out of bounds")
	}
	return Item{v: v, i: index}
}

func (v Version) ItemByName(name string) pkm.Item {
	encName := encodeText(strings.ToUpper(name))
	b := make([]byte, itemNameLength)
	for i := 0; i < indexSizeSpecies; i++ {
		v.ROM.Seek(addrItemData+int64(i*itemDataLength), 0)
		v.ROM.Read(b)
		if bytes.Equal(encName, truncateText(b)) {
			return Item{v: v, i: i}
		}
	}
	return nil
}

func (v Version) AbilityIndexSize() int {
	return indexSizeAbility
}

func (v Version) Abilities() []pkm.Ability {
	a := make([]pkm.Ability, indexSizeAbility)
	for i := range a {
		a[i] = Ability{v: v, i: i}
	}
	return a
}

func (v Version) AbilityByIndex(index int) pkm.Ability {
	if index < 0 || index >= indexSizeAbility {
		panic("ability index out of bounds")
	}
	return Ability{v: v, i: index}
}

func (v Version) AbilityByName(name string) pkm.Ability {
	encName := encodeText(strings.ToUpper(name))
	b := make([]byte, abilityNameLength)
	v.ROM.Seek(addrAbilityName, 0)
	for i := 0; i < indexSizeAbility; i++ {
		v.ROM.Read(b)
		if bytes.Equal(encName, truncateText(b)) {
			return Ability{v: v, i: i}
		}
	}
	return nil
}

func (v Version) MoveIndexSize() int {
	return indexSizeMove
}

func (v Version) Moves() []pkm.Move {
	a := make([]pkm.Move, indexSizeMove)
	for i := range a {
		a[i] = Move{v: v, i: i}
	}
	return a
}

func (v Version) MoveByIndex(index int) pkm.Move {
	if index < 0 || index >= indexSizeMove {
		panic("move index out of bounds")
	}
	return Move{v: v, i: index}
}

func (v Version) MoveByName(name string) pkm.Move {
	encName := encodeText(strings.ToUpper(name))
	b := make([]byte, moveNameLength)
	v.ROM.Seek(addrMoveName, 0)
	for i := 0; i < indexSizeMove; i++ {
		v.ROM.Read(b)
		if bytes.Equal(encName, truncateText(b)) {
			return Move{v: v, i: i}
		}
	}
	return nil
}

func (v Version) TMIndexSize() int {
	return indexSizeTM
}

func (v Version) TMs() []pkm.TM {
	a := make([]pkm.TM, indexSizeTM)
	for i := range a {
		a[i] = TM{v: v, i: i}
	}
	return a
}

func (v Version) TMByIndex(index int) pkm.TM {
	if index < 0 || index >= indexSizeTM {
		panic("TM index out of bounds")
	}
	return TM{v: v, i: index}
}

func (v Version) TMByName(name string) pkm.TM {
	name = strings.ToUpper(name)
	if len(name) != 4 ||
		name[1] != 'M' ||
		'0' > name[2] || name[2] > '9' ||
		'0' > name[3] || name[3] > '9' {
		return nil
	}
	size, off := 0, 0
	switch name[0] {
	case 'T':
		size, off = 50, -1
	case 'H':
		size, off = 8, 49
	default:
		return nil
	}
	n, _ := strconv.ParseInt(name[2:4], 10, 8)
	if int(n) <= 0 || int(n) > size {
		return nil
	}
	return TM{v: v, i: int(n) + off}
}

func (v Version) BankIndexSize() int {
	return indexSizeBank
}

func (v Version) Banks() []pkm.Bank {
	a := make([]pkm.Bank, indexSizeBank)
	for i := range a {
		a[i] = Bank{v: v, i: i}
	}
	return a
}

func (v Version) BankByIndex(index int) pkm.Bank {
	if index < 0 || index >= indexSizeBank {
		panic("bank index out of bounds")
	}
	return Bank{v: v, i: index}
}

func (v Version) AllMaps() []pkm.Map {
	// TODO
	return nil
}

func (v Version) MapByName(name string) pkm.Map {
	// TODO
	return nil
}
