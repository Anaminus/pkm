package gen3

import (
	"fmt"
	"github.com/anaminus/pkm"
)

var (
	structMoveName = makeStruct(
		13, // 0 Name
	)
	structMoveData = makeStruct(
		1, // 0 Effect
		1, // 1 BasePower
		1, // 2 Type
		1, // 3 Accuracy
		1, // 4 PP
		1, // 5 EffectAccuracy
		1, // 6 Affectee
		1, // 7 Priority
		1, // 8 Flags
		3, // 9 Padding
	)
	structTMMove = makeStruct(
		2, // 0 Move
	)
)

type Move struct {
	v Version
	i int
}

func (m Move) Name() string {
	b := readStruct(
		m.v.ROM,
		m.v.AddrMoveName,
		m.i,
		structMoveName,
	)
	return decodeTextString(b)
}

func (m Move) Index() int {
	return m.i
}

func (m Move) Description() string {
	b := readStruct(
		m.v.ROM,
		m.v.AddrMoveDescPtr,
		m.i-1,
		structPtr,
	)
	m.v.ROM.Seek(int64(decPtr(b)), 0)
	return readTextString(m.v.ROM)
}

func (m Move) Type() pkm.Type {
	b := readStruct(
		m.v.ROM,
		m.v.AddrMoveData,
		m.i,
		structMoveData,
		2,
	)
	return pkm.Type(b[0])
}

func (m Move) BasePower() byte {
	b := readStruct(
		m.v.ROM,
		m.v.AddrMoveData,
		m.i,
		structMoveData,
		1,
	)
	return b[0]
}

func (m Move) Accuracy() byte {
	b := readStruct(
		m.v.ROM,
		m.v.AddrMoveData,
		m.i,
		structMoveData,
		3,
	)
	return b[0]
}

func (m Move) Effect() pkm.Effect {
	b := readStruct(
		m.v.ROM,
		m.v.AddrMoveData,
		m.i,
		structMoveData,
		0,
	)
	return pkm.Effect(b[0])
}

func (m Move) EffectAccuracy() byte {
	b := readStruct(
		m.v.ROM,
		m.v.AddrMoveData,
		m.i,
		structMoveData,
		5,
	)
	return b[0]
}

func (m Move) Affectee() pkm.Affectee {
	b := readStruct(
		m.v.ROM,
		m.v.AddrMoveData,
		m.i,
		structMoveData,
		6,
	)
	return pkm.Affectee(b[0])
}

func (m Move) Priority() int8 {
	b := readStruct(
		m.v.ROM,
		m.v.AddrMoveData,
		m.i,
		structMoveData,
		7,
	)
	return int8(b[0])
}

func (m Move) Flags() pkm.MoveFlags {
	b := readStruct(
		m.v.ROM,
		m.v.AddrMoveData,
		m.i,
		structMoveData,
		8,
	)
	return pkm.MoveFlags(b[0])
}

type TM struct {
	v Version
	i int
}

func (tm TM) Name() string {
	if tm.i > 49 {
		return "HM" + fmt.Sprintf("%02d", tm.i-49)
	}
	return "TM" + fmt.Sprintf("%02d", tm.i+1)
}

func (tm TM) Index() int {
	return tm.i
}

func (tm TM) Move() pkm.Move {
	b := readStruct(
		tm.v.ROM,
		tm.v.AddrTMMove,
		tm.i,
		structTMMove,
	)
	return Move{v: tm.v, i: int(decUint16(b))}
}
