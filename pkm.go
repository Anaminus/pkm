package pkm

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io"
	"strings"
)

// Codec converts a game's text data to and from another format.
type Codec interface {
	// Name is the name of the codec.
	Name() string

	// Decode decodes game text data from src to another encoding in dst,
	// returning the number of bytes written to dst. Returns an error if the
	// codec does not support decoding.
	Decode(dst io.Writer, src io.Reader) (written int, err error)

	// Encode encodes characters from src into text data in dst, returning the
	// number of bytes written to dst. Returns an error if the codec does not
	// support encoding.
	Encode(dst io.Writer, src io.Reader) (written int, err error)
}

// DecodeText decodes a slice of text data into a string.
func DecodeText(codec Codec, b []byte) (s string, err error) {
	var buf bytes.Buffer
	if n, err := codec.Decode(&buf, bytes.NewReader(b)); err != nil {
		return "", err
	} else {
		return string(buf.Bytes()[:n]), nil
	}

}

// EncodeText encodes a string into a slice of text data.
func EncodeText(codec Codec, s string) (b []byte, err error) {
	var buf bytes.Buffer
	if n, err := codec.Encode(&buf, strings.NewReader(s)); err != nil {
		return nil, err
	} else {
		return buf.Bytes()[:n], nil
	}
}

// Version represents a single version of a pokemon game.
type Version interface {
	// Returns a the name of the version.
	Name() string

	// Returns the game code of the version.
	GameCode() GameCode

	// Returns a Query value that can be used to used for deep searching.
	Query() Query

	// Returns a list of codecs supported by the version.
	Codecs() []Codec
	// Returns the default codec used when text data is returned.
	DefaultCodec() Codec

	// Returns a size that fits all species indices (the maximum index + 1).
	SpeciesIndexSize() int
	// Returns the species at the given index. Not all species are valid.
	// Panics if the index exceeds SpeciesIndexSize.
	SpeciesByIndex(index int) Species
	// Returns a species by name. The name is case-insensitive, and uses the
	// default codec. Returns nil if no species was found.
	SpeciesByName(name string) Species

	// Returns a list of pokedexes.
	Pokedex() []Pokedex
	// Returns a pokedex by name. The name is case-insensitive, and uses the
	// default codec. Returns nil if no pokedex was found. An empty string
	// returns the pokedex that contains all species (national pokedex).
	PokedexByName(name string) Pokedex

	// Returns a size that fits all item indices (the maximum index + 1).
	ItemIndexSize() int
	// Returns a list of items. Array indices may not correspond to item indices.
	Items() []Item
	// Returns an item by its index. Panics if the index exceeds
	// ItemIndexSize.
	ItemByIndex(index int) Item
	// Returns an item by name. The name is case-insensitive, and uses the
	// default codec. Returns nil if no item was found.
	ItemByName(name string) Item

	// Returns a size that fits all ability indices (the maximum index + 1).
	AbilityIndexSize() int
	// Returns a list of abilities. Array indices may not correspond to
	// ability indices.
	Abilities() []Ability
	// Returns an ability by its index. Panics if the index exceeds
	// AbilityIndexSize.
	AbilityByIndex(index int) Ability
	// Returns an ability by name. The name is case-insensitive, and uses the
	// default codec. Returns nil if no ability was found.
	AbilityByName(name string) Ability

	// Returns a size that fits all move indices (the maximum index + 1).
	MoveIndexSize() int
	// Returns a list of moves. Array indices may not correspond to move
	// indices.
	Moves() []Move
	// Returns a move by its index. Panics if the index exceeds MoveIndexSize.
	MoveByIndex(index int) Move
	// Returns a move by name. The name is case-insensitive, and uses the
	// default codec. Returns nil if no move was found.
	MoveByName(name string) Move

	// Returns a size that fits all TM indices (the maximum index + 1).
	TMIndexSize() int
	// Returns a list of TMs. Array indices may not correspond to TM indices.
	TMs() []TM
	// Returns a TM by its index. Panics if the index exceeds TMIndexSize.
	TMByIndex(index int) TM
	// Returns a TM by name. The name is case-insensitive, and uses the
	// default codec. Returns nil if no TM was found.
	TMByName(name string) TM

	// Attempts to retrieve the sizes of the bank pointer table and map
	// pointer tables by scanning the ROM. Other map-related functions must be
	// called after this.
	ScanBanks()
	// Returns a size that fits all bank indices (the maximum index + 1).
	// Panics if ScanBanks has not been called.
	BankIndexSize() int
	// Returns a list of banks. Array indices may not correspond to bank
	// indices. Panics if ScanBanks has not been called.
	Banks() []Bank
	// Returns a bank by its index. Panics if the index exceeds BankIndexSize.
	// Panics if ScanBanks has not been called.
	BankByIndex(index int) Bank
	// Returns a list of all maps from every bank. Panics if ScanBanks has not
	// been called.
	AllMaps() []Map
	// Returns a map from any bank, by its name. The name is case-insensitive,
	// and uses the default codec. Returns nil if no map was found. Panics if
	// ScanBanks has not been called. Note that multiple maps may share the
	// same name, in which case the first map of the given name is returned.
	MapByName(name string) Map

	// TypeEffectiveness calculates the effectiveness of an attack of a given
	// type versus the types of a defending pokemon. If boths types of the
	// defender are the same, then they are counted as a single type.
	TypeEffectiveness(atk Type, def [2]Type) float64
}

////////////////////////////////////////////////////////////////

type GameCode [4]byte

func (gc GameCode) String() string {
	return "AGB-" + string(gc[:])
}

func (gc GameCode) Type() string {
	switch gc[0] {
	case 'A':
		return "Normal:A"
	case 'B':
		return "Normal:B"
	case 'C':
		return "Normal:C"
	case 'F':
		return "Famicom"
	case 'K':
		return "Acceleration sensor"
	case 'P':
		return "e-Reader"
	case 'R':
		return "Rumble/Gyro"
	case 'U':
		return "RTC/Solar sensor"
	case 'V':
		return "Rumble"
	default:
		return "Unknown:" + string(gc[0])
	}
	return string(gc[0])
}

func (gc GameCode) ID() string {
	return string(gc[1:3])
}

func (gc GameCode) Language() string {
	switch gc[3] {
	case 'J':
		return "Japanese"
	case 'E':
		return "English"
	case 'P':
		return "Europe"
	case 'G':
		return "German"
	case 'F':
		return "French"
	case 'I':
		return "Italian"
	case 'S':
		return "Spanish"
	default:
		return "Unknown:" + string(gc[3])
	}
}

////////////////////////////////////////////////////////////////

// Pokedex contains a list of species in a Version.
type Pokedex interface {
	// A name identifying the pokedex.
	Name() string
	// The number of species in the pokedex.
	Size() int
	// Returns the species of a given pokedex number. Note that the number
	// starts at 1.
	Species(number int) Species
	// Returns a list of all species in the pokedex. Note that array indices
	// may not correspond to pokedex numbers.
	AllSpecies() []Species
	// Returns the pokedex number for a given species. Returns 0 if the
	// species is not in the pokedex.
	SpeciesNumber(species Species) int
}

////////////////////////////////////////////////////////////////

// Species is a single species of pokemon in a Version.
type Species interface {
	// The index of the species.
	Index() int
	// The name of the species.
	Name() string
	// The pokedex category of the species.
	Category() string
	// The height of the species.
	Height() Height
	// The weight of the species.
	Weight() Weight
	// The pokedex description of the species.
	Description() string
	// The base stats of the species.
	BaseStats() Stats
	// The two types of the species. Both types are the same if the species
	// has only one type.
	Type() [2]Type
	// A value contributing to how easily the species can be captured.
	CatchRate() byte
	// A value contributing to how much experience the species yields in
	// battle.
	ExpYield() byte
	// The effort points the species yields in battle.
	EffortPoints() EffortPoints
	// Items that have a chance of being held by a wild pokemon of this
	// species. The first item has a 50% chance of being held, while the
	// second has a 5% chance of being held, or 100% chance if both items are
	// the same.
	HeldItem() [2]Item
	// The chance of a wild pokemon of this species being a certain gender.
	GenderRatio() GenderRatio
	// How long it takes for an egg of this species to hatch.
	EggCycles() byte
	// The starting friendship value of a caught pokemon of this species.
	BaseFriendship() byte
	// How a pokemon of this species gains experience.
	LevelType() LevelType
	// Which groups of pokemon this species is able to breed with.
	EggGroup() [2]EggGroup
	// The abilities that this species is able to have. A wild pokemon of this
	// species has a chance of having one or the other.
	Ability() [2]Ability
	// The rate at which wild pokemon of this species appear in the Safari
	// Zone.
	SafariRate() byte
	// The pokemon's color.
	Color() SpeciesColor
	// A list of moves that can be learned, at which levels, by a pokemon of
	// this species.
	LearnedMoves() []LevelMove
	// Returns whether a pokemon of this species can learn a move from a given
	// TM.
	CanLearnTM(tm TM) bool
	// Returns a list of the TMs that a pokemon of this species can learn.
	LearnableTMs() []TM
	// A list of species this species can evolve into, and by which methods.
	Evolutions() []Evolution
}

// Stats is the base stats of a species.
type Stats struct {
	HitPoints,
	Attack,
	Defense,
	Speed,
	SpAttack,
	SpDefense byte
}

func (s Stats) Total() int {
	return int(s.HitPoints) +
		int(s.Attack) +
		int(s.SpAttack) +
		int(s.Defense) +
		int(s.SpDefense) +
		int(s.Speed)
}

// EffortPoints is the number of effort points a wild pokemon yields when
// defeated.
type EffortPoints uint16

func (ep EffortPoints) Hitpoints() byte { return byte(ep & 3 >> 0) }
func (ep EffortPoints) Attack() byte    { return byte(ep & 12 >> 2) }
func (ep EffortPoints) Defense() byte   { return byte(ep & 48 >> 4) }
func (ep EffortPoints) Speed() byte     { return byte(ep & 192 >> 6) }
func (ep EffortPoints) SpAttack() byte  { return byte(ep & 768 >> 8) }
func (ep EffortPoints) SpDefense() byte { return byte(ep & 3072 >> 10) }
func (ep EffortPoints) Total() byte {
	return ep.Hitpoints() +
		ep.Attack() +
		ep.Defense() +
		ep.Speed() +
		ep.SpAttack() +
		ep.SpDefense()
}

// GenderRatio indicates the chance that a wild pokemon will be of a certain
// gender, or genderless.
type GenderRatio byte

func (g GenderRatio) Male() float64 {
	if g == 255 {
		return 0
	}
	return 1 - float64(g)/254
}

func (g GenderRatio) Female() float64 {
	if g == 255 {
		return 0
	}
	return float64(g) / 254
}

func (g GenderRatio) Genderless() bool {
	return g == 255
}

func (g GenderRatio) String() string {
	switch g {
	case 255:
		return "Genderless"
	default:
		return fmt.Sprintf("%.2f%% male / %.2f%% female",
			(1-float64(g)/254)*100,
			float64(g)/254*100,
		)
	}
}

// LevelType indicates how a species gains experience.
type LevelType byte

const (
	MediumFast  LevelType = 0
	Erratic               = 1
	Fluctuating           = 2
	MediumSlow            = 3
	Fast                  = 4
	Slow                  = 5
)

func (l LevelType) String() string {
	switch l {
	case MediumFast:
		return "Medium-fast"
	case Erratic:
		return "Erratic"
	case Fluctuating:
		return "Fluctuating"
	case MediumSlow:
		return "Medium-slow"
	case Fast:
		return "Fast"
	case Slow:
		return "Slow"
	}
	return "Unknown"
}

// EggGroup indicates species that can breed with one another.
type EggGroup byte

const (
	EggMonster      EggGroup = 1
	EggWater1                = 2
	EggBug                   = 3
	EggFlying                = 4
	EggField                 = 5
	EggFairy                 = 6
	EggGrass                 = 7
	EggHumanLike             = 8
	EggWater3                = 9
	EggMineral               = 10
	EggAmorphous             = 11
	EggWater2                = 12
	EggDitto                 = 13
	EggDragon                = 14
	EggUndiscovered          = 15
)

func (g EggGroup) String() string {
	switch g {
	case EggMonster:
		return "Monster"
	case EggWater1:
		return "Water1"
	case EggBug:
		return "Bug"
	case EggFlying:
		return "Flying"
	case EggField:
		return "Field"
	case EggFairy:
		return "Fairy"
	case EggGrass:
		return "Grass"
	case EggHumanLike:
		return "HumanLike"
	case EggWater3:
		return "Water3"
	case EggMineral:
		return "Mineral"
	case EggAmorphous:
		return "Amorphous"
	case EggWater2:
		return "Water2"
	case EggDitto:
		return "Ditto"
	case EggDragon:
		return "Dragon"
	case EggUndiscovered:
		return "Undiscovered"
	}
	return "Unknown"
}

// SpeciesColor indicates the color of a species.
type SpeciesColor byte

const (
	ColorRed    SpeciesColor = 0
	ColorBlue                = 1
	ColorYellow              = 2
	ColorGreen               = 3
	ColorBlack               = 4
	ColorBrown               = 5
	ColorPurple              = 6
	ColorGray                = 7
	ColorWhite               = 8
	ColorPink                = 9
)

func (c SpeciesColor) String() string {
	switch c {
	case ColorRed:
		return "Red"
	case ColorBlue:
		return "Blue"
	case ColorYellow:
		return "Yellow"
	case ColorGreen:
		return "Green"
	case ColorBlack:
		return "Black"
	case ColorBrown:
		return "Brown"
	case ColorPurple:
		return "Purple"
	case ColorGray:
		return "Gray"
	case ColorWhite:
		return "White"
	case ColorPink:
		return "Pink"
	}
	return "Unknown"
}

// LevelMove pairs a move with a level at which the move can be learned.
type LevelMove struct {
	Level byte
	Move  Move
}

type Height uint16

func (h Height) Centimeters() float64 {
	return h.Meters() * 100
}

func (h Height) Meters() float64 {
	return float64(h) / 10
}

func (h Height) Feet() float64 {
	return h.Meters() / 0.3048
}

func (h Height) Inches() float64 {
	return h.Feet() * 12
}

type Weight uint16

func (w Weight) Kilograms() float64 {
	return float64(w) / 10
}

func (w Weight) Pounds() float64 {
	return w.Kilograms() / 0.45359237
}

////////////////////////////////////////////////////////////////

// Item represents a single item for a Version.
type Item interface {
	Index() int
	Name() string
	Description() string
	Price() int
}

////////////////////////////////////////////////////////////////

// Ability represents a single pokemon ability in a Version.
type Ability interface {
	Index() int
	Name() string
	Description() string
}

////////////////////////////////////////////////////////////////

// Move represents a single pokemon move in a Version.
type Move interface {
	Index() int
	Name() string
	Description() string
	Type() Type
	BasePower() byte
	Accuracy() byte
	PowerPoints() byte
	Effect() Effect
	EffectAccuracy() byte
	Affectee() Affectee
	Priority() int8
	Flags() MoveFlags
}

// Effect indicates the type of effect that a move has.
type Effect byte

// Affectee indicates which pokemon are affected by a move in battle.
type Affectee byte

func (a Affectee) String() string {
	switch a {
	case 0x00:
		return "Selected target"
	case 0x01:
		return "Depends on the attack"
	case 0x02:
		return "Unused"
	case 0x04:
		return "Random target"
	case 0x08:
		return "Both foes"
	case 0x10:
		return "User"
	case 0x20:
		return "Both foes and partner"
	case 0x40:
		return "Opponent field"
	}
	return ""
}

// MoveFlags indicate various properties of a move.
type MoveFlags byte

const (
	Contact MoveFlags = 1 << iota
	Protect
	MagicCoat
	Snatch
	MirrorMove
	KingsRock
)

func (f MoveFlags) Contact() bool    { return f&Contact != 0 }
func (f MoveFlags) Protect() bool    { return f&Protect != 0 }
func (f MoveFlags) MagicCoat() bool  { return f&MagicCoat != 0 }
func (f MoveFlags) Snatch() bool     { return f&Snatch != 0 }
func (f MoveFlags) MirrorMove() bool { return f&MirrorMove != 0 }
func (f MoveFlags) KingsRock() bool  { return f&KingsRock != 0 }

////////////////////////////////////////////////////////////////

// TM represents a single TM or HM in a Version.
type TM interface {
	Index() int
	Name() string
	Move() Move
}

////////////////////////////////////////////////////////////////

// Evolution describes how a species evolves and what it evolves into.
type Evolution interface {
	// The species evolved into.
	Target() Species
	// The condition in which the evolution occurs.
	Method() uint16
	// A parameter applied to the evolution method.
	Param() uint16
	// Returns a string that combines the method and parameter to describe the
	// condition for evolution.
	MethodString() string
}

////////////////////////////////////////////////////////////////

// Bank comprises a number of Maps in a Version.
type Bank interface {
	// Returns the bank's index.
	Index() int
	// Returns a size that fits all map indices (the maximum index + 1).
	MapIndexSize() int
	// Returns a list of maps. Array indices may not correspond to map
	// indices.
	Maps() []Map
	// Returns a map by its index. Panics if the index exceeds MapIndexSize.
	MapByIndex(index int) Map
	// Returns a map by its name. The name is case-insensitive, and uses the
	// default codec. Returns nil if no map was found. Note that multiple maps
	// may share the same name, in which case the first map of the given name
	// is returned.
	MapByName(name string) Map
}

// Map represents map data in a Version.
type Map interface {
	// Returns the bank index of the map.
	BankIndex() int
	// Returns the index of the map.
	Index() int
	// Returns the name of the map.
	Name() string
	// Returns the cells that make up the map.
	Layout() Layout
	// Returns the cells that make up the border of the map.
	Border() Layout
	// Returns the tileset used to render the map.
	Tileset() Tileset
	// Render the map as an image. Returns an image for each layer in the map.
	Image() []*image.NRGBA
	// Render the map border as an image. Returns an image for each layer in
	// the map.
	BorderImage() []*image.NRGBA
	// Render the tileset used to draw the map. Returns an Image for each
	// layer in the map.
	//
	// The width determines how wide the image will be, in blocks (each being
	// 16x16 pixels), wrapping blocks that exceed the width to the next row. A
	// width < 1 will make the image as wide as the number of blocks in the
	// tileset.
	TilesetImage(width int) []*image.NRGBA
	// Returns the color that is drawn when no opaque colors have been drawn
	// to a pixel.
	BackgroundColor() color.NRGBA
	// Returns a list of all the areas in the map which may contain
	// encounters.
	Encounters() []EncounterList
}

// Represents the layout of a map. The layout is a grid of cells. Each cell
// consists of an index that points to a block in some tileset, as well as the
// index of an attribute, which appears to be movement permissions.
type Layout interface {
	Width() int
	Height() int
	Cell(i int) (block, attr int)
	CellAt(x, y int) (block, attr int)
}

// Tileset comprises a list of blocks, as well as an image and palette list.
// The full set is created from a global and local tileset.
type Tileset interface {
	Block(i int) Block
	Sprite(i int) Sprite
	Palette(i int) Palette
}

// A block is made up of two layers, with each layer containing 4 tiles,
// representing the four quadrants of the block.
type Block interface {
	Tile(index, layer int) Tile
}

// A tile contains a sprite index and a palette index, as well as whether the
// sprite is flipped on each axis.
type Tile uint16

func (t Tile) SpriteIndex() int {
	// 0000 0011 1111 1111
	return int(t & 1023)
}
func (t Tile) FlipX() bool {
	// 0000 0100 0000 0000
	return t&1024 != 0
}
func (t Tile) FlipY() bool {
	// 0000 1000 0000 0000
	return t&2048 != 0
}
func (t Tile) PaletteIndex() int {
	// 1111 0000 0000 0000
	return int(t & 61440 >> 12)
}

// Draws the tile to an image, given a tileset and an offset.
func (t Tile) DrawTo(ts Tileset, img *image.NRGBA, ox, oy int) {
	s := ts.Sprite(t.SpriteIndex())
	p := ts.Palette(t.PaletteIndex())
	for i := 0; i < 64; i++ {
		x, y := i%8, i/8
		if t.FlipX() {
			x = 7 - x
		}
		if t.FlipY() {
			y = 7 - y
		}
		ci := s.ColorIndex(i)
		var c color.NRGBA
		if ci > 0 {
			c = p.Color(ci)
		}
		img.SetNRGBA(ox+x, oy+y, c)
	}
}

// A sprite is an 8x8 array of pixel data. Each value in the sprite is an
// index that points to a color in a palette.
type Sprite interface {
	ColorIndex(i int) int
	Len() int
}

// A palette contains 16 colors.
type Palette interface {
	Color(i int) color.NRGBA
}

// Create an image from a tileset and layout.
func DrawImage(l Layout, ts Tileset, layer int) *image.NRGBA {
	w := l.Width()
	h := l.Height()
	img := image.NewNRGBA(image.Rect(0, 0, w*16, h*16))
	for i := 0; i < w*h; i++ {
		cx, cy := i%w, i/w
		bi, _ := l.Cell(i)
		block := ts.Block(bi)
		for j := 0; j < 4; j++ {
			sx, sy := j%2, j/2
			block.Tile(j, layer).DrawTo(ts, img, cx*16+sx*8, cy*16+sy*8)
		}
	}
	return img
}

func CombineLayers(layers ...*image.NRGBA) *image.NRGBA {
	bounds := layers[0].Bounds()
	for i := 1; i < len(layers); i++ {
		bounds = bounds.Union(layers[i].Bounds())
	}
	final := image.NewNRGBA(bounds)
	for _, layer := range layers {
		draw.Draw(
			final,
			layer.Bounds(),
			layer,
			image.Pt(0, 0),
			draw.Over,
		)
	}
	return final
}

////////////////////////////////////////////////////////////////

// EncounterList contains information about the species that can be
// encountered in a particular area type of a map.
type EncounterList interface {
	// Returns a name representing the type of area within a map where
	// encounters can occur.
	Name() string
	// Returns whether species can be encountered in the area (which implies
	// that the map contains the area type).
	Populated() bool
	// Returns the probability (0-1) that traversing a block in the area will
	// lead to an encounter. Returns 0 if the area is unpopulated.
	EncounterRate() float64
	// Returns the maximum size of the encounter table.
	EncounterIndexSize() int
	// Returns a list of possible species encounters. Returns nil if the area
	// is unpopulated.
	Encounters() []Encounter
	// Returns the species encounter at the given index. Returns nil if the
	// area is unpopulated.
	Encounter(index int) Encounter
	// Returns the probability (0-1) that the species encounter at the given
	// index will be selected.
	SpeciesRate(index int) float64
}

// Encounter contains information about a single encounter.
type Encounter interface {
	// Returns the minimum level at which the species may be encountered.
	MinLevel() int
	// Returns the maximum level at which the species may be encountered.
	MaxLevel() int
	// The encountered species.
	Species() Species
}

////////////////////////////////////////////////////////////////

// Type indicates the type of a pokemon.
type Type byte

const (
	TypeNormal   Type = 0
	TypeFighting      = 1
	TypeFlying        = 2
	TypePoison        = 3
	TypeGround        = 4
	TypeRock          = 5
	TypeBug           = 6
	TypeGhost         = 7
	TypeSteel         = 8
	TypeCurse         = 9
	TypeFire          = 10
	TypeWater         = 11
	TypeGrass         = 12
	TypeElectric      = 13
	TypePsychic       = 14
	TypeIce           = 15
	TypeDragon        = 16
	TypeDark          = 17
)

func (t Type) String() string {
	switch t {
	case TypeNormal:
		return "Normal"
	case TypeFighting:
		return "Fighting"
	case TypeFlying:
		return "Flying"
	case TypePoison:
		return "Poison"
	case TypeGround:
		return "Ground"
	case TypeRock:
		return "Rock"
	case TypeBug:
		return "Bug"
	case TypeGhost:
		return "Ghost"
	case TypeSteel:
		return "Steel"
	case TypeCurse:
		return "???"
	case TypeFire:
		return "Fire"
	case TypeWater:
		return "Water"
	case TypeGrass:
		return "Grass"
	case TypeElectric:
		return "Electric"
	case TypePsychic:
		return "Psychic"
	case TypeIce:
		return "Ice"
	case TypeDragon:
		return "Dragon"
	case TypeDark:
		return "Dark"
	}
	return "Unknown"
}

////////////////////////////////////////////////////////////////

// Query is used to extract interesting information from a Version.
type Query interface {
	// Returns a species by name. The name is case-insensitive, and uses the
	// default codec. Returns nil if no species was found.
	SpeciesByName(name string) Species
	// Returns an item by name. The name is case-insensitive, and uses the
	// default codec. Returns nil if no item was found.
	ItemByName(name string) Item
	// Returns an ability by name. The name is case-insensitive, and uses the
	// default codec. Returns nil if no ability was found.
	AbilityByName(name string) Ability
	// Returns a move by name. The name is case-insensitive, and uses the
	// default codec. Returns nil if no move was found.
	MoveByName(name string) Move
	// Returns a TM by name. The name is case-insensitive, and uses the
	// default codec. Returns nil if no TM was found.
	TMByName(name string) TM
	// Returns a map from any bank, by its name. The name is case-insensitive,
	// and uses the default codec. Returns nil if no map was found.
	MapByName(name string) Map

	// Returns species that learn a given move by leveling up.
	SpeciesLearningMove(move Move) []Species
	// Returns species that can learn a given TM.
	SpeciesLearningTM(tm TM) []Species
	// Returns locations from which a given species can be encountered in the
	// wild.
	SpeciesLocations(species Species) []Map
}
