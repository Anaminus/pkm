package emerald

import (
	"github.com/anaminus/pkm"
)

const (
	addrSpeciesName  = 0x003185C8
	addrSpeciesData  = 0x003203CC
	addrPokedexData  = 0x0056B5B0
	addrLevelMove    = 0x003230DC
	addrLevelMovePtr = 0x0032937C
	addrSpeciesTM    = 0x0031E898
	addrEvoTable     = 0x0032531C
)

var (
	structSpeciesName = makeStruct(
		11, // 0 Name
	)
	structSpeciesData = makeStruct(
		1, // 00 HitPoints
		1, // 01 Attack
		1, // 02 Defense
		1, // 03 Speed
		1, // 04 SpAttack
		1, // 05 SpDefense
		1, // 06 Type1
		1, // 07 Type2
		1, // 08 CatchRate
		1, // 09 ExpYield
		2, // 10 EffortPoints
		2, // 11 HeldItem1
		2, // 12 HeldItem2
		1, // 13 GenderRatio
		1, // 14 EggCycles
		1, // 15 BaseFriendship
		1, // 16 LevelType
		1, // 17 EggGroup1
		1, // 18 EggGroup2
		1, // 19 Ability1
		1, // 20 Ability2
		1, // 21 SafariRate
		1, // 22 Color/Flip
		2, // 23 Padding
	)
	structDexData = makeStruct(
		12, // 0 Category
		2,  // 1 Height
		2,  // 2 Weight
		4,  // 3 DescPtr1
		4,  // 4 DescPtr2
		2,  // 5 PokémonScale
		2,  // 6 PokémonOffset
		2,  // 7 TrainerScale
		2,  // 8 TrainerOffset
	)
)

type Species struct {
	v Version
	i int
}

func (s Species) Name() string {
	b := readStruct(
		s.v.ROM,
		addrSpeciesName,
		s.i,
		structSpeciesName,
	)
	return decodeTextString(b)
}

func (s Species) Index() int {
	return s.i
}

func (s Species) Description() string {
	b := readStruct(
		s.v.ROM,
		addrPokedexData,
		s.i,
		structDexData,
		3,
	)
	s.v.ROM.Seek(int64(decPtr(b)), 0)
	return readTextString(s.v.ROM)
}

func (s Species) Category() string {
	b := readStruct(
		s.v.ROM,
		addrPokedexData,
		s.i,
		structDexData,
		0,
	)
	return decodeTextString(b)
}

func (s Species) Height() pkm.Height {
	b := readStruct(
		s.v.ROM,
		addrPokedexData,
		s.i,
		structDexData,
		1,
	)
	return pkm.Height(decUint16(b))
}

func (s Species) Weight() pkm.Weight {
	b := readStruct(
		s.v.ROM,
		addrPokedexData,
		s.i,
		structDexData,
		1,
	)
	return pkm.Weight(decUint16(b))
}

func (s Species) BaseStats() pkm.Stats {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		0, 1, 2, 3, 4, 5,
	)
	return pkm.Stats{
		HitPoints: b[0],
		Attack:    b[1],
		Defense:   b[2],
		Speed:     b[3],
		SpAttack:  b[4],
		SpDefense: b[5],
	}
}

func (s Species) Type() [2]pkm.Type {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		6, 7,
	)
	return [2]pkm.Type{
		pkm.Type(b[0]),
		pkm.Type(b[1]),
	}
}

func (s Species) CatchRate() byte {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		8,
	)
	return b[0]
}

func (s Species) ExpYield() byte {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		9,
	)
	return b[0]
}

func (s Species) EffortPoints() pkm.EffortPoints {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		10,
	)
	return pkm.EffortPoints(decUint16(b))
}

func (s Species) HeldItem() [2]pkm.Item {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		0,
	)
	return [2]pkm.Item{
		Item{v: s.v, i: int(decUint16(b[0:2]))},
		Item{v: s.v, i: int(decUint16(b[2:4]))},
	}
}

func (s Species) GenderRatio() pkm.GenderRatio {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		13,
	)
	return pkm.GenderRatio(b[0])
}

func (s Species) EggCycles() byte {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		14,
	)
	return b[0]
}

func (s Species) BaseFriendship() byte {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		15,
	)
	return b[0]
}

func (s Species) LevelType() pkm.LevelType {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		16,
	)
	return pkm.LevelType(b[0])
}

func (s Species) EggGroup() [2]pkm.EggGroup {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		17, 18,
	)
	return [2]pkm.EggGroup{
		pkm.EggGroup(b[0]),
		pkm.EggGroup(b[1]),
	}
}

func (s Species) Ability() [2]pkm.Ability {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		19, 20,
	)
	return [2]pkm.Ability{
		Ability{v: s.v, i: int(b[0])},
		Ability{v: s.v, i: int(b[1])},
	}
}

func (s Species) SafariRate() byte {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		21,
	)
	return b[0]
}

func (s Species) Color() pkm.Color {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		22,
	)
	return pkm.Color(b[0] & 127)
}

func (s Species) Flipped() bool {
	b := readStruct(
		s.v.ROM,
		addrSpeciesData,
		s.i,
		structSpeciesData,
		22,
	)
	return b[0]&128 != 0
}

func (s Species) LearnedMoves() []pkm.LevelMove {
	// TODO
	return nil
}

func (s Species) CanTeachTM(tm pkm.TM) bool {
	// TODO
	return false
}
