package gen3

import (
	"github.com/anaminus/pkm"
)

const (
	addrEncounterPtr = 0x00552d48
	addrMapName      = 0x005A1480
)

var (
	structMapHeader = makeStruct(
		4, // Map data
		4, // Event data
		4, // Map scripts
		4, // Connections
		2, // Music index
		2, // Map pointer index
		1, // Label index
		1, // Visibility
		1, // Weather
		1, // Map type
		2, // Unknown
		1, // Show label on entry
		1, // In-battle field model id
	)
	structMapLayoutData = makeStruct(
		4, // Width
		4, // Height
		4, // Border
		4, // Map data / Tile structure
		4, // Global tileset
		4, // Local tileset
		1, // Border width
		1, // Border height
	)
	structTilesetHeader = makeStruct(
		1, // Compressed
		1, // Is primary
		1, // Unknown
		1, // Unknown
		4, // Pointer to tileset image
		4, // Pointer to color palettes
		4, // Pointer to blocks
		4, // Pointer to animation routine
		4, // Pointer to behavior and background bytes
	)
	structConnHeader = makeStruct(
		4, // Amount of map connections
		4, // Pointer to connection data
	)
	structConnData = makeStruct(
		4, // Connection direction
		4, // Offset
		1, // Map Bank
		1, // Map Number
		2, // Padding
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
	// TODO
	return 0
}

func (b Bank) Maps() []pkm.Map {
	// TODO
	return nil
}

func (b Bank) MapByIndex(index int) pkm.Map {
	// TODO
	return nil
}

func (b Bank) MapByName(name string) pkm.Map {
	// TODO
	return nil
}

type Map struct {
	v *Version
	i int
}

func (m Map) BankIndex() int {
	// TODO
	return 0
}

func (m Map) Index() int {
	// TODO
	return 0
}

func (m Map) Encounters() {
	// TODO
	return
}
