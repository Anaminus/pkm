package gen3

import (
	"github.com/anaminus/pkm"
)

const (
	addrEncounterPtr = 0x00552d48
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

func (m Map) Encounters() {
	// TODO
	return
}

func (m Map) Name() string {
	m.v.ROM.Seek(int64(m.v.AddrBanksPtr), 0)
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
	m.v.ROM.Seek(int64(decPtr(b[4:8])), 0)
	return readTextString(m.v.ROM)
}
