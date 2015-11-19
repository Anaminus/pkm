package gen3

import (
	"github.com/anaminus/pkm"
)

type Move struct {
	v Version
	i int
}

func (m Move) Name() string {
	// TODO
	return ""
}

func (m Move) Index() int {
	// TODO
	return 0
}

func (m Move) Description() string {
	// TODO
	return ""
}

func (m Move) Type() pkm.Type {
	// TODO
	return pkm.Type(0)
}

func (m Move) BasePower() byte {
	// TODO
	return 0
}

func (m Move) Accuracy() byte {
	// TODO
	return 0
}

func (m Move) Effect() pkm.Effect {
	// TODO
	return pkm.Effect(0)
}

func (m Move) EffectAccuracy() byte {
	// TODO
	return 0
}

func (m Move) Affectee() pkm.Affectee {
	// TODO
	return pkm.Affectee(0)
}

func (m Move) Priority() int8 {
	// TODO
	return 0
}

func (m Move) Flags() pkm.MoveFlags {
	// TODO
	return pkm.MoveFlags(0)
}

type TM struct {
	v Version
	i int
}

func (tm TM) Name() string {
	// TODO
	return ""
}

func (tm TM) Index() int {
	// TODO
	return 0
}

func (tm TM) Move() pkm.Move {
	// TODO
	return nil
}
