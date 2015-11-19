package pkm

import (
	"bytes"
	"fmt"
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
	if _, err := codec.Decode(&buf, bytes.NewReader(b)); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// EncodeText encodes a string into a slice of text data.
func EncodeText(codec Codec, s string) (b []byte, err error) {
	var buf bytes.Buffer
	if _, err := codec.Encode(&buf, strings.NewReader(s)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Version represents a single version of a pokemon game.
type Version interface {
	// Returns a the name of the version.
	Name() string

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

	// Returns a size that fits all bank indices (the maximum index + 1).
	BankIndexSize() int
	// Returns a list of banks. Array indices may not correspond to bank
	// indices.
	Banks() []Bank
	// Returns a bank by its index. Panics if the index exceeds BankIndexSize.
	BankByIndex(index int) Bank
	// Returns a list of all maps from every bank.
	AllMaps() []Map
	// Returns a map from any bank, by its name. The name is case-insensitive,
	// and uses the default codec. Returns nil if no map was found.
	MapByName(name string) Map
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
	// The name of the species.
	Name() string
	// The index of the species.
	Index() int

	Category() string
	Height() Height
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
	Color() Color
	// A list indicating which moves can be learned at which levels.
	LearnedMoves() []LevelMove
	// Returns whether a pokemon of this species can learn a move from a given
	// TM.
	CanTeachTM(tm TM) bool
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

// EffortPoints is the number of effort points a wild pokemon yields when
// defeated.
type EffortPoints uint16

func (ev EffortPoints) HP() byte        { return byte(ev & 3 >> 0) }
func (ev EffortPoints) Attack() byte    { return byte(ev & 12 >> 2) }
func (ev EffortPoints) Defense() byte   { return byte(ev & 48 >> 4) }
func (ev EffortPoints) Speed() byte     { return byte(ev & 192 >> 6) }
func (ev EffortPoints) SpAttack() byte  { return byte(ev & 768 >> 8) }
func (ev EffortPoints) SpDefense() byte { return byte(ev & 3072 >> 10) }

// GenderRatio indicates the chance that a wild pokemon will be of a certain
// gender, or genderless.
type GenderRatio byte

func (g GenderRatio) String() string {
	switch g {
	case 255:
		return "Genderless"
	default:
		return fmt.Sprintf("%.2f%% male / %.2f%% female",
			float32(g)/254*100,
			(1-float32(g)/254)*100,
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
	EggMonster      = 1
	EggWater1       = 2
	EggBug          = 3
	EggFlying       = 4
	EggField        = 5
	EggFairy        = 6
	EggGrass        = 7
	EggHumanLike    = 8
	EggWater3       = 9
	EggMineral      = 10
	EggAmorphous    = 11
	EggWater2       = 12
	EggDitto        = 13
	EggDragon       = 14
	EggUndiscovered = 15
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

// Color indicates the color of a species.
type Color byte

const (
	ColorRed    = 0
	ColorBlue   = 1
	ColorYellow = 2
	ColorGreen  = 3
	ColorBlack  = 4
	ColorBrown  = 5
	ColorPurple = 6
	ColorGray   = 7
	ColorWhite  = 8
	ColorPink   = 9
)

func (c Color) String() string {
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
	Level int
	Move  Move
}

type Height byte
type Weight byte

////////////////////////////////////////////////////////////////

// Item represents a single item for a Version.
type Item interface {
	Name() string
	Index() int
	Description() string
	Price() int
}

////////////////////////////////////////////////////////////////

// Ability represents a single pokemon ability in a Version.
type Ability interface {
	Name() string
	Index() int
	Description() string
}

////////////////////////////////////////////////////////////////

// Move represents a single pokemon move ina a Version.
type Move interface {
	Name() string
	Index() int
	Description() string
	Type() Type
	BasePower() byte
	Accuracy() byte
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

func (f MoveFlags) Contact() bool    { return f&1 != 0 }
func (f MoveFlags) Protect() bool    { return f&2 != 0 }
func (f MoveFlags) MagicCoat() bool  { return f&4 != 0 }
func (f MoveFlags) Snatch() bool     { return f&8 != 0 }
func (f MoveFlags) MirrorMove() bool { return f&16 != 0 }
func (f MoveFlags) KingsRock() bool  { return f&32 != 0 }

////////////////////////////////////////////////////////////////

// TM represents a single TM or HM in a Version.
type TM interface {
	Name() string
	Index() int
	Move() Move
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
	// default codec. Returns nil if no map was found.
	MapByName(name string) Map
}

// Map represents map data in a Version.
type Map interface {
	BankIndex() int
	Index() int
	Encounters()
}

////////////////////////////////////////////////////////////////

// Type indicates the type of a pokemon.
type Type byte

const (
	TypeNormal   = 0
	TypeFighting = 1
	TypeFlying   = 2
	TypePoison   = 3
	TypeGround   = 4
	TypeRock     = 5
	TypeBug      = 6
	TypeGhost    = 7
	TypeSteel    = 8
	TypeCurse    = 9
	TypeFire     = 10
	TypeWater    = 11
	TypeGrass    = 12
	TypeElectric = 13
	TypePsychic  = 14
	TypeIce      = 15
	TypeDragon   = 16
	TypeDark     = 17
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