package gen3

import (
	"fmt"
	"github.com/anaminus/pkm"
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
	structSpeciesTM = makeStruct(
		8, // 0 TMs
	)
	structEvolution = makeStruct(
		2, // 0 Method
		2, // 1 Parameter
		2, // 2 Target
	)
)

type Species struct {
	v *Version
	i int
}

func (s Species) Name() string {
	b := readStruct(
		s.v.ROM,
		s.v.AddrSpeciesName,
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
		s.v.AddrPokedexData,
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
		s.v.AddrPokedexData,
		s.i,
		structDexData,
		0,
	)
	return decodeTextString(b)
}

func (s Species) Height() pkm.Height {
	b := readStruct(
		s.v.ROM,
		s.v.AddrPokedexData,
		s.i,
		structDexData,
		1,
	)
	return pkm.Height(decUint16(b))
}

func (s Species) Weight() pkm.Weight {
	b := readStruct(
		s.v.ROM,
		s.v.AddrPokedexData,
		s.i,
		structDexData,
		1,
	)
	return pkm.Weight(decUint16(b))
}

func (s Species) BaseStats() pkm.Stats {
	b := readStruct(
		s.v.ROM,
		s.v.AddrSpeciesData,
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
		s.v.AddrSpeciesData,
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
		s.v.AddrSpeciesData,
		s.i,
		structSpeciesData,
		8,
	)
	return b[0]
}

func (s Species) ExpYield() byte {
	b := readStruct(
		s.v.ROM,
		s.v.AddrSpeciesData,
		s.i,
		structSpeciesData,
		9,
	)
	return b[0]
}

func (s Species) EffortPoints() pkm.EffortPoints {
	b := readStruct(
		s.v.ROM,
		s.v.AddrSpeciesData,
		s.i,
		structSpeciesData,
		10,
	)
	return pkm.EffortPoints(decUint16(b))
}

func (s Species) HeldItem() [2]pkm.Item {
	b := readStruct(
		s.v.ROM,
		s.v.AddrSpeciesData,
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
		s.v.AddrSpeciesData,
		s.i,
		structSpeciesData,
		13,
	)
	return pkm.GenderRatio(b[0])
}

func (s Species) EggCycles() byte {
	b := readStruct(
		s.v.ROM,
		s.v.AddrSpeciesData,
		s.i,
		structSpeciesData,
		14,
	)
	return b[0]
}

func (s Species) BaseFriendship() byte {
	b := readStruct(
		s.v.ROM,
		s.v.AddrSpeciesData,
		s.i,
		structSpeciesData,
		15,
	)
	return b[0]
}

func (s Species) LevelType() pkm.LevelType {
	b := readStruct(
		s.v.ROM,
		s.v.AddrSpeciesData,
		s.i,
		structSpeciesData,
		16,
	)
	return pkm.LevelType(b[0])
}

func (s Species) EggGroup() [2]pkm.EggGroup {
	b := readStruct(
		s.v.ROM,
		s.v.AddrSpeciesData,
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
		s.v.AddrSpeciesData,
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
		s.v.AddrSpeciesData,
		s.i,
		structSpeciesData,
		21,
	)
	return b[0]
}

func (s Species) Color() pkm.Color {
	b := readStruct(
		s.v.ROM,
		s.v.AddrSpeciesData,
		s.i,
		structSpeciesData,
		22,
	)
	return pkm.Color(b[0] & 127)
}

func (s Species) Flipped() bool {
	b := readStruct(
		s.v.ROM,
		s.v.AddrSpeciesData,
		s.i,
		structSpeciesData,
		22,
	)
	return b[0]&128 != 0
}

func (s Species) LearnedMoves() []pkm.LevelMove {
	s.v.ROM.Seek(int64(s.v.AddrLevelMovePtr)+int64(s.i*4), 0)
	p := readPtr(s.v.ROM)
	s.v.ROM.Seek(int64(p), 0)
	q := make([]byte, 256)
	lms := make([]pkm.LevelMove, 0, 8)
	for j := 0; j < len(q); j += 2 {
		l := q[j+1]
		if q[j] == 0xFF || l == 0xFF {
			break
		}
		lms = append(lms, pkm.LevelMove{
			Level: l / 2,
			Move:  Move{v: s.v, i: int(q[j]) | int(l%2)<<8},
		})
	}
	return lms
}

func (s Species) CanLearnTM(tm pkm.TM) bool {
	b := make([]byte, 1)
	s.v.ROM.Seek(int64(s.v.AddrSpeciesTM)+int64(s.i*8+tm.Index()/8), 0)
	s.v.ROM.Read(b)
	return b[0]&(1<<uint(tm.Index()%8)) != 0
}

func (s Species) LearnableTMs() []pkm.TM {
	b := readStruct(
		s.v.ROM,
		s.v.AddrSpeciesTM,
		s.i,
		structSpeciesTM,
	)
	tms := make([]pkm.TM, 0, indexSizeTM)
	for i := 0; i < indexSizeTM; i++ {
		if b[i/8]&(1<<uint(i%8)) != 0 {
			tms = append(tms, TM{v: s.v, i: s.i})
		}
	}
	return tms
}

func (s Species) Evolutions() []pkm.Evolution {
	const evoIndexSize = 5

	evos := make([]pkm.Evolution, 0, evoIndexSize)
	for i := 0; i < evoIndexSize; i++ {
		b := readStruct(
			s.v.ROM,
			s.v.AddrSpeciesEvo+uint32(i*structEvolution.Size()),
			s.i,
			structEvolution,
			0,
		)
		if decUint16(b) == 0 {
			continue
		}
		evos = append(evos, Evolution{
			v: s.v,
			s: s.i,
			i: i,
		})
	}
	return evos
}

type Evolution struct {
	v *Version
	s int
	i int
}

func (e Evolution) Target() pkm.Species {
	b := readStruct(
		e.v.ROM,
		e.v.AddrSpeciesEvo+uint32(e.i*structEvolution.Size()),
		e.s,
		structEvolution,
		2,
	)
	return Species{v: e.v, i: int(decUint16(b))}
}

func (e Evolution) Method() uint16 {
	b := readStruct(
		e.v.ROM,
		e.v.AddrSpeciesEvo+uint32(e.i*structEvolution.Size()),
		e.s,
		structEvolution,
		0,
	)
	return decUint16(b)
}

func (e Evolution) Param() uint16 {
	b := readStruct(
		e.v.ROM,
		e.v.AddrSpeciesEvo+uint32(e.i*structEvolution.Size()),
		e.s,
		structEvolution,
		1,
	)
	return decUint16(b)
}

func (e Evolution) MethodString() string {
	b := readStruct(
		e.v.ROM,
		e.v.AddrSpeciesEvo+uint32(e.i*structEvolution.Size()),
		e.s,
		structEvolution,
		0, 1,
	)
	method := decUint16(b[0:2])
	param := decUint16(b[2:4])

	switch method {
	case 0x1:
		return fmt.Sprintf("Friendship")
	case 0x2:
		return fmt.Sprintf("Friendship (Day)")
	case 0x3:
		return fmt.Sprintf("Friendship (Night)")
	case 0x4:
		return fmt.Sprintf("Level %d", param)
	case 0x5:
		return fmt.Sprintf("Trade")
	case 0x6:
		return fmt.Sprintf("Trade holding %s", (Item{v: e.v, i: int(param)}).Name())
	case 0x7:
		return fmt.Sprintf("Use %s", (Item{v: e.v, i: int(param)}).Name())
	case 0x8:
		return fmt.Sprintf("Level %d if ATK > DEF", param)
	case 0x9:
		return fmt.Sprintf("Level %d if ATK = DEF", param)
	case 0xA:
		return fmt.Sprintf("Level %d if ATK < DEF", param)
	case 0xB:
		return fmt.Sprintf("Personality[1] (%d)", param)
	case 0xC:
		return fmt.Sprintf("Personality[2] (%d)", param)
	case 0xD:
		return fmt.Sprintf("Level %d (Spawns extra)", param)
	case 0xE:
		return fmt.Sprintf("Level %d (Spawned)", param)
	case 0xF:
		return fmt.Sprintf("Beauty (%d)", param)
	}
	return "None"
}
