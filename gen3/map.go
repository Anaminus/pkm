package gen3

import (
	"github.com/anaminus/pkm"
	"strings"
)

var (
	structMapHeader = makeStruct(
		4, // 00 Map data
		4, // 01 Event data
		4, // 02 Map scripts
		4, // 03 Connections
		2, // 04 Music index
		2, // 05 Map pointer index
		1, // 06 Label index
		1, // 07 Visibility
		1, // 08 Weather
		1, // 09 Map type
		2, // 10 Unknown
		1, // 11 Show label on entry
		1, // 12 In-battle field model id
	)
	structMapLabel = makeStruct(
		1, // 0 Unknown
		1, // 1 Unknown
		1, // 2 Unknown
		1, // 3 Unknown
		4, // 4 Pointer to map name
	)
	structMapLayoutData = makeStruct(
		4, // 0 Width
		4, // 1 Height
		4, // 2 Border
		4, // 3 Map data / Tile structure
		4, // 4 Global tileset
		4, // 5 Local tileset
		1, // 6 Border width
		1, // 7 Border height
	)
	structTilesetHeader = makeStruct(
		1, // 0 Compressed
		1, // 1 Is primary
		1, // 2 Unknown
		1, // 3 Unknown
		4, // 4 Pointer to tileset image
		4, // 5 Pointer to color palettes
		4, // 6 Pointer to blocks
		4, // 7 Pointer to animation routine
		4, // 8 Pointer to behavior and background bytes
	)
	structConnHeader = makeStruct(
		4, // 0 Amount of map connections
		4, // 1 Pointer to connection data
	)
	structConnData = makeStruct(
		4, // 0 Connection direction
		4, // 1 Offset
		1, // 2 Map Bank
		1, // 3 Map Number
		2, // 4 Padding
	)
	structEncounterPtrs = makeStruct(
		1, // 0 Bank
		1, // 1 Map
		2, // 2 Padding
		4, // 3 Pointer to encounter table (grass)
		4, // 4 Pointer to encounter table (water)
		4, // 5 Pointer to encounter table (rock)
		4, // 6 Pointer to encounter table (rod)
	)
	structEncounterHeader = makeStruct(
		1, // 0 Encounter rate
		3, // 1 Padding
		4, // 2 Pointer to encounter table
	)
	structEncounter = makeStruct(
		1, // 0 MinLevel
		1, // 1 MaxLevel
		2, // 2 Species
	)
)

type Bank struct {
	v *Version
	i int
}

func (b Bank) Index() int {
	return b.i
}

func (b Bank) MapIndexSize() int {
	return b.v.sizeMapTable[b.i]
}

func (b Bank) Maps() []pkm.Map {
	maps := make([]pkm.Map, 0, 64)
	for i := 0; i < b.MapIndexSize(); i++ {
		maps = append(maps, Map{v: b.v, b: b.i, i: i})
	}
	return maps
}

func (b Bank) MapByIndex(index int) pkm.Map {
	if index < 0 || index >= b.MapIndexSize() {
		panic("bank index out of bounds")
	}
	return Map{v: b.v, b: b.i, i: index}
}

func (b Bank) MapByName(name string) pkm.Map {
	name = strings.ToUpper(name)
	for _, m := range b.Maps() {
		if name == m.Name() {
			return m
		}
	}
	return nil
}

type Map struct {
	v    *Version
	b, i int
}

func (m Map) BankIndex() int {
	return m.b
}

func (m Map) Index() int {
	return m.i
}

func (m Map) Encounters() []pkm.EncounterList {
	ptrs := [4]ptr{}
	for p := 0; p < len(ptrs); p++ {
		for i := 0; ; i++ {
			b := readStruct(
				m.v.ROM,
				m.v.AddrEncounterList,
				i,
				structEncounterPtrs,
				0, 1,
			)
			if b[0] == 0xFF && b[1] == 0xFF {
				break
			} else if int(b[0]) == m.BankIndex() && int(b[1]) == m.Index() {
				b := readStruct(
					m.v.ROM,
					m.v.AddrEncounterList,
					i,
					structEncounterPtrs,
					p+3,
				)
				ptrs[p] = decPtr(b)
				break
			}
		}
	}
	return []pkm.EncounterList{
		EncounterGrass{v: m.v, p: ptrs[0]},
		EncounterWater{v: m.v, p: ptrs[1]},
		EncounterRock{v: m.v, p: ptrs[2]},
		EncounterRod{v: m.v, p: ptrs[3]},
	}
}

func (m Map) Name() string {
	m.v.ROM.Seek(m.v.AddrBanksPtr.ROM(), 0)
	b := readStruct(
		m.v.ROM,
		readPtr(m.v.ROM),
		m.b,
		structPtr,
	)
	b = readStruct(
		m.v.ROM,
		decPtr(b),
		m.i,
		structPtr,
	)
	b = readStruct(
		m.v.ROM,
		decPtr(b),
		0,
		structMapHeader,
		6,
	)
	b = readStruct(
		m.v.ROM,
		m.v.AddrMapLabel,
		int(b[0]),
		structMapLabel,
	)
	m.v.ROM.Seek(decPtr(b[4:8]).ROM(), 0)
	return readTextString(m.v.ROM)
}

////////////////////////////////////////////////////////////////

type Encounter struct {
	v *Version
	p ptr
	i int
}

func (e Encounter) MinLevel() int {
	b := readStruct(
		e.v.ROM,
		e.p,
		e.i,
		structEncounter,
		0,
	)
	return int(b[0])
}

func (e Encounter) MaxLevel() int {
	b := readStruct(
		e.v.ROM,
		e.p,
		e.i,
		structEncounter,
		1,
	)
	return int(b[0])
}

func (e Encounter) Species() pkm.Species {
	b := readStruct(
		e.v.ROM,
		e.p,
		e.i,
		structEncounter,
		2,
	)
	return Species{v: e.v, i: int(decUint16(b))}
}

////////////////////////////////////////////////////////////////

func encounterRate(v *Version, p ptr) byte {
	if !p.ValidROM() {
		return 0
	}
	b := readStruct(
		v.ROM,
		p,
		0,
		structEncounterHeader,
		0,
	)
	return b[0]
}

func encounters(v *Version, p ptr, s int) []pkm.Encounter {
	if !p.ValidROM() {
		return nil
	}
	b := readStruct(
		v.ROM,
		p,
		0,
		structEncounterHeader,
		2,
	)
	p = decPtr(b)
	encounters := make([]pkm.Encounter, s)
	for i := range encounters {
		encounters[i] = Encounter{v: v, p: p, i: i}
	}
	return encounters
}

func encounter(v *Version, p ptr, s, index int) pkm.Encounter {
	if index < 0 || index >= s {
		panic("encounter index out of bounds")
	}
	if !p.ValidROM() {
		return nil
	}
	b := readStruct(
		v.ROM,
		p,
		0,
		structEncounterHeader,
		2,
	)
	return Encounter{v: v, p: decPtr(b), i: index}
}

////////////////////////////////////////////////////////////////

type EncounterGrass struct {
	v *Version
	p ptr
}

func (e EncounterGrass) Name() string {
	return "Grass"
}

func (e EncounterGrass) Populated() bool {
	return e.p.ValidROM()
}

func (e EncounterGrass) EncounterIndexSize() int {
	return 12
}

func (e EncounterGrass) EncounterRate() byte {
	return encounterRate(e.v, e.p)
}

func (e EncounterGrass) Encounters() []pkm.Encounter {
	return encounters(e.v, e.p, e.EncounterIndexSize())
}

func (e EncounterGrass) Encounter(index int) pkm.Encounter {
	return encounter(e.v, e.p, e.EncounterIndexSize(), index)
}

func (e EncounterGrass) SpeciesRate(index int) (rate float32) {
	if index < 0 || index >= e.EncounterIndexSize() {
		panic("species rate index out of bounds")
	}
	switch index {
	case 0, 1:
		rate = 0.2
	case 2, 3, 4, 5:
		rate = 0.1
	case 6, 7:
		rate = 0.05
	case 8, 9:
		rate = 0.04
	case 10, 11:
		rate = 0.01
	}
	return
}

////////////////////////////////////////////////////////////////

type EncounterWater struct {
	v *Version
	p ptr
}

func (e EncounterWater) Name() string {
	return "Water"
}

func (e EncounterWater) Populated() bool {
	return e.p.ValidROM()
}

func (e EncounterWater) EncounterIndexSize() int {
	return 5
}

func (e EncounterWater) EncounterRate() byte {
	return encounterRate(e.v, e.p)
}

func (e EncounterWater) Encounters() []pkm.Encounter {
	return encounters(e.v, e.p, e.EncounterIndexSize())
}

func (e EncounterWater) Encounter(index int) pkm.Encounter {
	return encounter(e.v, e.p, e.EncounterIndexSize(), index)
}

func (e EncounterWater) SpeciesRate(index int) (rate float32) {
	if index < 0 || index >= e.EncounterIndexSize() {
		panic("species rate index out of bounds")
	}
	switch index {
	case 0:
		rate = 0.6
	case 1:
		rate = 0.3
	case 2:
		rate = 0.05
	case 3:
		rate = 0.04
	case 4:
		rate = 0.01
	}
	return
}

////////////////////////////////////////////////////////////////

type EncounterRock struct {
	v *Version
	p ptr
}

func (e EncounterRock) Name() string {
	return "Rock"
}

func (e EncounterRock) Populated() bool {
	return e.p.ValidROM()
}

func (e EncounterRock) EncounterIndexSize() int {
	return 5
}

func (e EncounterRock) EncounterRate() byte {
	return encounterRate(e.v, e.p)
}

func (e EncounterRock) Encounters() []pkm.Encounter {
	return encounters(e.v, e.p, e.EncounterIndexSize())
}

func (e EncounterRock) Encounter(index int) pkm.Encounter {
	return encounter(e.v, e.p, e.EncounterIndexSize(), index)
}

func (e EncounterRock) SpeciesRate(index int) (rate float32) {
	if index < 0 || index >= e.EncounterIndexSize() {
		panic("species rate index out of bounds")
	}
	switch index {
	case 0:
		rate = 0.6
	case 1:
		rate = 0.3
	case 2:
		rate = 0.05
	case 3:
		rate = 0.04
	case 4:
		rate = 0.01
	}
	return
}

////////////////////////////////////////////////////////////////

type EncounterRod struct {
	v *Version
	p ptr
}

func (e EncounterRod) Name() string {
	return "Rod"
}

func (e EncounterRod) Populated() bool {
	return e.p.ValidROM()
}

func (e EncounterRod) EncounterIndexSize() int {
	return 10
}

func (e EncounterRod) EncounterRate() byte {
	return encounterRate(e.v, e.p)
}

func (e EncounterRod) Encounters() []pkm.Encounter {
	return encounters(e.v, e.p, e.EncounterIndexSize())
}

func (e EncounterRod) Encounter(index int) pkm.Encounter {
	return encounter(e.v, e.p, e.EncounterIndexSize(), index)
}

func (e EncounterRod) SpeciesRate(index int) (rate float32) {
	if index < 0 || index >= e.EncounterIndexSize() {
		panic("species rate index out of bounds")
	}
	switch index {
	case 0:
		rate = 0.7
	case 1:
		rate = 0.3
	case 2:
		rate = 0.6
	case 3:
		rate = 0.2
	case 4:
		rate = 0.2
	case 5:
		rate = 0.4
	case 6:
		rate = 0.4
	case 7:
		rate = 0.15
	case 8:
		rate = 0.04
	case 9:
		rate = 0.01
	}
	return
}

func (e EncounterRod) RodType(index int) (rod Rod) {
	if index < 0 || index >= e.EncounterIndexSize() {
		panic("rod type index out of bounds")
	}
	switch index {
	case 0, 1:
		rod = OldRod
	case 2, 3, 4:
		rod = GoodRod
	case 5, 6, 7, 8, 9:
		rod = SuperRod
	}
	return
}

type Rod byte

const (
	OldRod Rod = iota
	GoodRod
	SuperRod
)

func (r Rod) String() string {
	switch r {
	case OldRod:
		return "Old"
	case GoodRod:
		return "Good"
	case SuperRod:
		return "Super"
	}
	return ""
}
