package gen3

import (
	"bytes"
	"github.com/anaminus/pkm"
	"io"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////

var (
	structTypeEffect = makeStruct(
		1, // 0 Attack type
		1, // 1 Defend type
		1, // 2 Effect
	)
)

type Version struct {
	ROM                io.ReadSeeker
	name               string
	pokedex            []pokedexData
	sizeMapTable       []int
	AddrAbilityName    ptr // Table of ability names.
	AddrAbilityDescPtr ptr // Table of pointers to ability descriptions.
	AddrBanksPtr       ptr // Pointer to bank pointer table.
	AddrEncounterList  ptr // List of map references to encounter table pointers.
	AddrItemData       ptr // Table of item data.
	AddrLevelMovePtr   ptr // Table of pointers to learned-move data.
	AddrMapLabel       ptr // Table of map label data.
	AddrMoveName       ptr // Table of move names.
	AddrMoveData       ptr // Table of move data.
	AddrMoveDescPtr    ptr // Table of pointers to move descriptions.
	AddrPokedexData    ptr // Table of pokedex data.
	AddrSpeciesData    ptr // Table of species data.
	AddrSpeciesEvo     ptr // Table of species evolution data.
	AddrSpeciesName    ptr // Table of species names.
	AddrSpeciesTM      ptr // Table of species TM compatibility.
	AddrTypeEffect     ptr // List of type effectiveness.
	AddrTMMove         ptr // Table of TM move mappings.
}

var _ = pkm.Version(&Version{})

func (v *Version) Name() string {
	return v.name
}

func (v *Version) GameCode() (gc pkm.GameCode) {
	v.ROM.Seek(addrGameCode.ROM(), 0)
	v.ROM.Read(gc[:])
	return
}

func (v *Version) Query() pkm.Query {
	return &Query{v: v}
}

func (v *Version) Codecs() []pkm.Codec {
	return []pkm.Codec{
		CodecUTF8,
		CodecASCII,
		CodecString,
		CodecPUA,
	}
}
func (v *Version) DefaultCodec() pkm.Codec {
	return defaultCodec
}

func (v *Version) SpeciesIndexSize() int {
	return indexSizeSpecies
}

func (v *Version) SpeciesByIndex(index int) pkm.Species {
	if index < 0 || index >= indexSizeSpecies {
		panic("species index out of bounds")
	}
	return Species{v: v, i: index}
}

func (v *Version) SpeciesByName(name string) pkm.Species {
	encName := encodeText(strings.ToUpper(name))
	b := make([]byte, structSpeciesName.Size())
	v.ROM.Seek(v.AddrSpeciesName.ROM(), 0)
	for i := 0; i < indexSizeSpecies; i++ {
		v.ROM.Read(b)
		if bytes.Equal(encName, truncateText(b)) {
			return Species{v: v, i: i}
		}
	}
	return nil
}

func (v *Version) Pokedex() []pkm.Pokedex {
	a := make([]pkm.Pokedex, len(v.pokedex))
	for i := range v.pokedex {
		a[i] = Pokedex{v: v, i: i}
	}
	return a
}

func (v *Version) PokedexByName(name string) pkm.Pokedex {
	name = strings.ToUpper(name)
	for _, dex := range v.Pokedex() {
		if strings.ToUpper(dex.Name()) == name {
			return dex
		}
	}
	return nil
}

func (v *Version) ItemIndexSize() int {
	return indexSizeItem
}

func (v *Version) Items() []pkm.Item {
	a := make([]pkm.Item, indexSizeItem)
	for i := range a {
		a[i] = Item{v: v, i: i}
	}
	return a
}

func (v *Version) ItemByIndex(index int) pkm.Item {
	if index < 0 || index >= indexSizeItem {
		panic("item index out of bounds")
	}
	return Item{v: v, i: index}
}

func (v *Version) ItemByName(name string) pkm.Item {
	encName := encodeText(strings.ToUpper(name))
	for i := 0; i < indexSizeSpecies; i++ {
		b := readStruct(
			v.ROM,
			v.AddrItemData,
			i,
			structItemData,
			0,
		)
		if bytes.Equal(encName, truncateText(b)) {
			return Item{v: v, i: i}
		}
	}
	return nil
}

func (v *Version) AbilityIndexSize() int {
	return indexSizeAbility
}

func (v *Version) Abilities() []pkm.Ability {
	a := make([]pkm.Ability, indexSizeAbility)
	for i := range a {
		a[i] = Ability{v: v, i: i}
	}
	return a
}

func (v *Version) AbilityByIndex(index int) pkm.Ability {
	if index < 0 || index >= indexSizeAbility {
		panic("ability index out of bounds")
	}
	return Ability{v: v, i: index}
}

func (v *Version) AbilityByName(name string) pkm.Ability {
	encName := encodeText(strings.ToUpper(name))
	b := make([]byte, structAbilityName.Size())
	v.ROM.Seek(v.AddrAbilityName.ROM(), 0)
	for i := 0; i < indexSizeAbility; i++ {
		v.ROM.Read(b)
		if bytes.Equal(encName, truncateText(b)) {
			return Ability{v: v, i: i}
		}
	}
	return nil
}

func (v *Version) MoveIndexSize() int {
	return indexSizeMove
}

func (v *Version) Moves() []pkm.Move {
	a := make([]pkm.Move, indexSizeMove)
	for i := range a {
		a[i] = Move{v: v, i: i}
	}
	return a
}

func (v *Version) MoveByIndex(index int) pkm.Move {
	if index < 0 || index >= indexSizeMove {
		panic("move index out of bounds")
	}
	return Move{v: v, i: index}
}

func (v *Version) MoveByName(name string) pkm.Move {
	encName := encodeText(strings.ToUpper(name))
	b := make([]byte, structMoveName.Size())
	v.ROM.Seek(v.AddrMoveName.ROM(), 0)
	for i := 0; i < indexSizeMove; i++ {
		v.ROM.Read(b)
		if bytes.Equal(encName, truncateText(b)) {
			return Move{v: v, i: i}
		}
	}
	return nil
}

func (v *Version) TMIndexSize() int {
	return indexSizeTM
}

func (v *Version) TMs() []pkm.TM {
	a := make([]pkm.TM, indexSizeTM)
	for i := range a {
		a[i] = TM{v: v, i: i}
	}
	return a
}

func (v *Version) TMByIndex(index int) pkm.TM {
	if index < 0 || index >= indexSizeTM {
		panic("TM index out of bounds")
	}
	return TM{v: v, i: index}
}

func (v *Version) TMByName(name string) pkm.TM {
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

func validMapHeader(rom io.ReadSeeker, p ptr) bool {
	maph := make([]byte, structMapHeader.Size())
	rom.Seek(p.ROM(), 0)
	rom.Read(maph)
	if p := decPtr(maph[0:4]); !p.ValidROM() {
		return false
	}
	if p := decPtr(maph[4:8]); !p.ValidROM() {
		return false
	}
	if p := decPtr(maph[8:12]); !p.ValidROM() {
		return false
	}
	return true
}

func (v *Version) ScanBanks() {
	if len(v.sizeMapTable) > 0 {
		return
	}

	ps := structPtr.Size()

	// Find size of bank pointer table.
	size := 256
	banks := make([]byte, 256*ps)
	v.ROM.Seek(v.AddrBanksPtr.ROM(), 0)
	v.ROM.Seek(readPtr(v.ROM).ROM(), 0)
	v.ROM.Read(banks)
	for i := 0; i < len(banks); i += ps {
		bptr := decPtr(banks[i : i+ps])
		// Stop if pointer isn't valid.
		if !bptr.ValidROM() {
			size = i / ps
			break
		}
		// If a bank pointer points to an address within the supposed bank
		// pointer table, then the table must end at that location.
		if int(bptr-v.AddrBanksPtr)/ps < size {
			size = int(bptr-v.AddrBanksPtr) / ps
		}
	}

	bptrs := map[ptr]bool{}
	for i := 0; i < size; i++ {
		bptrs[decPtr(banks[i*ps:i*ps+ps])] = true
	}

	// Find size of each map table.
	v.sizeMapTable = make([]int, size)
	maps := make([]byte, 256*ps)
	for i := 0; i < size; i++ {
		bptr := decPtr(banks[i*ps : i*ps+ps])
		v.ROM.Seek(bptr.ROM(), 0)
		v.ROM.Read(maps)
		v.sizeMapTable[i] = 256
		for j := 0; j < len(maps); j += ps {
			mptr := decPtr(maps[j : j+ps])
			if !mptr.ValidROM() ||
				// Compare current address to bank pointers.
				(j > 0 && bptrs[bptr+ptr(j)]) ||
				// Compare map pointer to bank pointers.
				bptrs[mptr] ||
				// Check that data at map pointer looks like map data.
				!validMapHeader(v.ROM, mptr) {
				v.sizeMapTable[i] = j / ps
				break
			}
		}
	}
}

func (v *Version) BankIndexSize() int {
	if len(v.sizeMapTable) == 0 {
		panic("banks have not been scanned")
	}

	return len(v.sizeMapTable)
}

func (v *Version) Banks() []pkm.Bank {
	if len(v.sizeMapTable) == 0 {
		panic("banks have not been scanned")
	}

	a := make([]pkm.Bank, len(v.sizeMapTable))
	for i := range a {
		a[i] = Bank{v: v, i: i}
	}
	return a
}

func (v *Version) BankByIndex(index int) pkm.Bank {
	if len(v.sizeMapTable) == 0 {
		panic("banks have not been scanned")
	}

	if index < 0 || index >= len(v.sizeMapTable) {
		panic("bank index out of bounds")
	}
	return Bank{v: v, i: index}
}

func (v *Version) AllMaps() []pkm.Map {
	if len(v.sizeMapTable) == 0 {
		panic("banks have not been scanned")
	}

	maps := make([]pkm.Map, 0, 520)
	for b, size := range v.sizeMapTable {
		for i := 0; i < size; i++ {
			maps = append(maps, Map{v: v, b: b, i: i})
		}
	}
	return maps
}

func (v *Version) MapByName(name string) pkm.Map {
	if len(v.sizeMapTable) == 0 {
		panic("banks have not been scanned")
	}

	name = strings.ToUpper(name)
	for _, m := range v.AllMaps() {
		if name == m.Name() {
			return m
		}
	}
	return nil
}

func (v *Version) TypeEffectiveness(atk pkm.Type, def [2]pkm.Type) float64 {
	v.ROM.Seek(v.AddrTypeEffect.ROM(), 0)
	var mult byte = 4
	for q := make([]byte, structTypeEffect.Size()); ; {
		v.ROM.Read(q)
		if a, d := pkm.Type(q[0]), pkm.Type(q[1]); a == 0xFF || d == 0xFF {
			break
		} else if a == atk && (d == def[0] || d == def[1]) {
			switch q[2] {
			case 0:
				mult = 0
			case 5:
				mult >>= 1
			case 20:
				mult <<= 1
			}
		}
	}
	return float64(mult) / 4
}
