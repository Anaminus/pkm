package gen3

import (
	"encoding/binary"
	"github.com/anaminus/pkm"
	"image"
	"image/color"
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
		4, // 2 Pointer to border
		4, // 3 Pointer to tile structure
		4, // 4 Pointer to global tileset
		4, // 5 Pointer to local tileset
		1, // 6 Border width
		1, // 7 Border height
	)
	structTilesetHeader = makeStruct(
		1, // 0 Compressed
		1, // 1 Is primary
		2, // 2 Padding
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

// Get pointer to map header.
func (m Map) headerPtr() ptr {
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
	return decPtr(b)
}

func (m Map) Layout() pkm.Layout {
	b := readStruct(
		m.v.ROM,
		m.headerPtr(),
		0,
		structMapHeader,
		0,
	)
	b = readStruct(
		m.v.ROM,
		decPtr(b),
		0,
		structMapLayoutData,
		0, 1, 3,
	)

	width := int(decUint32(b[0:4]))
	height := int(decUint32(b[4:8]))
	l := _layout{
		width:  width,
		height: height,
		cells:  make([]byte, width*height*2),
	}
	m.v.ROM.Seek(decPtr(b[8:12]).ROM(), 0)
	m.v.ROM.Read(l.cells)
	return l
}

func (m Map) Border() pkm.Layout {
	b := readStruct(
		m.v.ROM,
		m.headerPtr(),
		0,
		structMapHeader,
		0,
	)
	b = readStruct(
		m.v.ROM,
		decPtr(b),
		0,
		structMapLayoutData,
		2, 6, 7,
	)

	var width, height int
	if gc := m.v.GameCode(); gc == CodeFireRedEN || gc == CodeLeafGreenEN {
		width, height = int(b[4]), int(b[5])
	} else {
		width, height = 2, 2
	}
	l := _layout{
		width:  width,
		height: height,
		cells:  make([]byte, width*height*2),
	}
	m.v.ROM.Seek(decPtr(b[0:4]).ROM(), 0)
	m.v.ROM.Read(l.cells)
	return l
}

func (m Map) Tileset() pkm.Tileset {
	b := readStruct(
		m.v.ROM,
		m.headerPtr(),
		0,
		structMapHeader,
		0,
	)
	b = readStruct(
		m.v.ROM,
		decPtr(b),
		0,
		structMapLayoutData,
		4, 5,
	)

	ts := &_tileset{id: [2]uint32{binary.LittleEndian.Uint32(b[0:4]), binary.LittleEndian.Uint32(b[4:8])}}
	m.readTileset(ts, decPtr(b[0:4]), 0)
	m.readTileset(ts, decPtr(b[4:8]), 1)
	return ts
}

func (m Map) Image() []*image.NRGBA {
	ts := m.Tileset()
	l := m.Layout()
	return []*image.NRGBA{
		pkm.DrawImage(l, ts, 0),
		pkm.DrawImage(l, ts, 1),
	}
}

func (m Map) BorderImage() []*image.NRGBA {
	ts := m.Tileset()
	l := m.Border()
	return []*image.NRGBA{
		pkm.DrawImage(l, ts, 0),
		pkm.DrawImage(l, ts, 1),
	}
}

func (m Map) TilesetImage(width int) []*image.NRGBA {
	if width < 1 {
		width = 0x400
	}
	b := readStruct(
		m.v.ROM,
		m.headerPtr(),
		0,
		structMapHeader,
		0,
	)
	b = readStruct(
		m.v.ROM,
		decPtr(b),
		0,
		structMapLayoutData,
		4, 5,
	)

	ts := m.Tileset()

	height := 0x400 / width
	if 0x400%width != 0 {
		height++
	}

	l := _layout{
		width:  width,
		height: height,
		cells:  make([]byte, width*height*2),
	}
	for i := 0; i < 0x400; i++ {
		binary.LittleEndian.PutUint16(l.cells[i*2:], uint16(i))
	}
	return []*image.NRGBA{
		pkm.DrawImage(l, ts, 0),
		pkm.DrawImage(l, ts, 1),
	}
}

func (m Map) BackgroundColor() color.NRGBA {
	// TODO: It isn't necessary to read the entire tileset.
	return m.Tileset().Palette(0).Color(0)
}

// Reads a single tileset from ROM into a given tileset.
//
// Tilesets come in pairs. When read into memory, each component of the
// tilesets are combined. That is, the global block list is read into the
// first half of the block address space, while the local block list is read
// into the second half. The same occurs with images and palettes.
func (m Map) readTileset(ts *_tileset, p ptr, off int) {
	header := readStruct(
		m.v.ROM,
		p,
		0,
		structTilesetHeader,
	)
	// Tileset image
	//
	// An image is a sequence of 8x8 sprites, to which tileset blocks refer to
	// create full 16x16 blocks. An image is usually compressed.
	//
	// The image from the global tileset is read into the first half of the
	// image address space, while the local image is read into the second
	// half.
	{
		const size = len(ts.image) / 2
		m.v.ROM.Seek(decPtr(header[4:8]).ROM(), 0)
		if header[0] == 1 {
			b, ok := readLZ77(m.v.ROM)
			if ok {
				copy(ts.image[off*size:], b)
			}
		} else {
			m.v.ROM.Read(ts.image[off*size : off*size+size])
		}
	}
	// Palette
	//
	// GBA has room for 16 palettes in RAM. Tilesets each point to a set of 16
	// palettes in ROM.
	//
	// Since there are two tilesets per map, only a portion of a tileset's
	// palettes in ROM are read into RAM. A tileset's `primary` byte appears
	// to determine which palettes are selected. 0 selects palettes 0-5, while
	// 1 selects 6-11.
	//
	// The selected palettes of the global tileset are set to palettes 0-5 in
	// RAM. The selected palettes of the local tileset are set to palettes
	// 6-11 in RAM.
	//
	// Palettes 12-15 in ROM appear to be unused by tilesets, but nonetheless
	// contain data (12 might be used, but I have not yet seen evidence). In
	// RAM, these palettes are likely reserved for other purposes.
	//
	// Color 0 in a given palette is always drawn as transparent, regardless
	// of color. Color 0 of palette 0 in RAM is used as the backdrop color; it
	// is drawn when no opaque colors have been drawn to a pixel.
	{
		const size = 32 * 6
		m.v.ROM.Seek(decPtr(header[8:12]).ROM()+32*6*int64(header[1]), 0)
		m.v.ROM.Read(ts.pal[off*size : off*size+size])
	}
	// Blocks
	{
		const size = len(ts.blocks) / 2
		m.v.ROM.Seek(decPtr(header[12:16]).ROM(), 0)
		m.v.ROM.Read(ts.blocks[off*size : off*size+size])
	}
}

// Tileset comprises a list of blocks, as well as an image and palette list.
// The full set is created from a global and local tileset.
type _tileset struct {
	id     [2]uint32
	blocks [16384]byte
	image  [32768]byte
	pal    [512]byte
}

func (t _tileset) ID() [2]uint32 {
	return t.id
}

func (t _tileset) Block(i int) pkm.Block {
	return _block(t.blocks[i*16 : i*16+16])
}

func (t _tileset) Sprite(i int) pkm.Sprite {
	return _sprite(t.image[i*32 : i*32+32])
}

func (t _tileset) Palette(i int) pkm.Palette {
	return _palette(t.pal[i*32 : i*32+32])
}

func (t _tileset) BlockLen() int {
	return 1024
}

func (t _tileset) SpriteLen() int {
	return 1024
}

func (t _tileset) PaletteLen() int {
	return 16
}

// Represents the layout of a map. The layout is a grid of cells. Each cell
// consists of an index that points to a block in some tileset, as well as the
// index of an attribute, which appears to be movement permissions.
type _layout struct {
	width  int
	height int
	cells  []byte
}

func (l _layout) Width() int {
	return l.width
}

func (l _layout) Height() int {
	return l.height
}

func (l _layout) Cell(i int) (block, attr int) {
	n := decUint16(l.cells[i*2 : i*2+2])
	block = int(n & 1023)
	attr = int(n & 64512 >> 10)
	return
}

func (l _layout) CellAt(x, y int) (block, attr int) {
	return l.Cell(y*l.width + x)
}

// A block is made up of two layers, with each layer containing 4 tiles,
// representing the four quadrants of the block.
type _block []byte

func (b _block) Tile(index, layer int) pkm.Tile {
	return pkm.Tile(decUint16(b[index*2+8*layer:]))
}

// A sprite is an 8x8 array of pixel data. Each value in the sprite is an
// index that points to a color in a palette.
type _sprite []byte

func (s _sprite) ColorIndex(i int) int {
	if i%2 == 0 {
		return int(s[i/2] & 15)
	} else {
		return int(s[i/2] & 240 >> 4)
	}
}

func (s _sprite) Len() int {
	return len(s) * 2
}

// A palette contains 16 colors.
type _palette []byte

func (p _palette) Color(i int) color.NRGBA {
	c := _color(decUint16(p[i*2 : i*2+2]))
	return color.NRGBA{R: c.R(), G: c.G(), B: c.B(), A: 255}
}

// 15-bit color, with each channel occupying 5 bits, or 32 values per channel.
// Values are in the range 0-248.
type _color uint16

func (c _color) R() byte {
	b := c & 31
	return byte(b * 8)
}

func (c _color) G() byte {
	b := c & 992 >> 5
	return byte(b * 8)
}

func (c _color) B() byte {
	b := c & 31744 >> 10
	return byte(b * 8)
}

////////////////////////////////////////////////////////////////

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
	b := readStruct(
		m.v.ROM,
		m.headerPtr(),
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

func encounterRate(v *Version, p ptr) float64 {
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
	return float64(b[0]) / 255
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

func (e EncounterGrass) EncounterRate() float64 {
	return encounterRate(e.v, e.p)
}

func (e EncounterGrass) Encounters() []pkm.Encounter {
	return encounters(e.v, e.p, e.EncounterIndexSize())
}

func (e EncounterGrass) Encounter(index int) pkm.Encounter {
	return encounter(e.v, e.p, e.EncounterIndexSize(), index)
}

func (e EncounterGrass) SpeciesRate(index int) (rate float64) {
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

func (e EncounterWater) EncounterRate() float64 {
	return encounterRate(e.v, e.p)
}

func (e EncounterWater) Encounters() []pkm.Encounter {
	return encounters(e.v, e.p, e.EncounterIndexSize())
}

func (e EncounterWater) Encounter(index int) pkm.Encounter {
	return encounter(e.v, e.p, e.EncounterIndexSize(), index)
}

func (e EncounterWater) SpeciesRate(index int) (rate float64) {
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

func (e EncounterRock) EncounterRate() float64 {
	return encounterRate(e.v, e.p)
}

func (e EncounterRock) Encounters() []pkm.Encounter {
	return encounters(e.v, e.p, e.EncounterIndexSize())
}

func (e EncounterRock) Encounter(index int) pkm.Encounter {
	return encounter(e.v, e.p, e.EncounterIndexSize(), index)
}

func (e EncounterRock) SpeciesRate(index int) (rate float64) {
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

func (e EncounterRod) EncounterRate() float64 {
	return encounterRate(e.v, e.p)
}

func (e EncounterRod) Encounters() []pkm.Encounter {
	return encounters(e.v, e.p, e.EncounterIndexSize())
}

func (e EncounterRod) Encounter(index int) pkm.Encounter {
	return encounter(e.v, e.p, e.EncounterIndexSize(), index)
}

func (e EncounterRod) SpeciesRate(index int) (rate float64) {
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

func (r Rod) String() (s string) {
	switch r {
	case OldRod:
		s = "Old"
	case GoodRod:
		s = "Good"
	case SuperRod:
		s = "Super"
	}
	return
}
