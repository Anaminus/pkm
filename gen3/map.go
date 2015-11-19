package gen3

import (
	"github.com/anaminus/pkm"
)

const (
	addrBankPtr      = 0x00084AA4
	addrEncounterPtr = 0x00552d48
	addrMapName      = 0x005A1480
)

type Bank struct {
	v Version
	i int
}

func (b Bank) Index() int {
	// TODO
	return 0
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
	v Version
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
