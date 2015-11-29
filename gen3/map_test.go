package gen3_test

import (
	"github.com/anaminus/pkm/gen3"
	"testing"
)

func TestBank(t *testing.T) {
	ver := gen3.OpenROM(ROM(t))
	if ver == nil {
		t.Fatalf("failed to open ROM")
	}

	ver.ScanBanks()
	b := ver.BankByIndex(24)
	if v := b.Index(); v != 24 {
		t.Errorf("Index: unexpected result %d", v)
	}
	if v := b.MapIndexSize(); v != 108 {
		t.Errorf("MapIndexSize: unexpected result %d", v)
	}
	if v := b.Maps(); len(v) != 108 {
		t.Errorf("Maps: unexpected result length %d", len(v))
	}
	if v := b.MapByIndex(4); v == nil {
		t.Errorf("MapByIndex: unexpected result <nil>")
	} else if v.Index() != 4 {
		t.Errorf("MapByIndex: unexpected result %d (%s)", v.Index(), v.Name())
	}
	ExpectPanic(t, "MapByIndex", func() {
		b.MapByIndex(108)
	})
	if v := b.MapByName("RUSTURF TUNNEL"); v == nil {
		t.Errorf("MapByIndex: unexpected result <nil>")
	} else if v.Index() != 4 {
		t.Errorf("MapByName: unexpected result %d (%s)", v.Index(), v.Name())
	} else if v := b.MapByName("Rusturf Tunnel"); v == nil {
		t.Errorf("MapByName: not case-insensitive")
	}
	if v := b.MapByName("unknown"); v != nil {
		t.Errorf("MapByName: expected nil result")
	}
}

func TestMap(t *testing.T) {
	ver := gen3.OpenROM(ROM(t))
	if ver == nil {
		t.Fatalf("failed to open ROM")
	}

	ver.ScanBanks()
	m := ver.BankByIndex(24).MapByIndex(4)
	if v := m.BankIndex(); v != 24 {
		t.Errorf("BankIndex: unexpected result %d", v)
	}
	if v := m.Index(); v != 4 {
		t.Errorf("Index: unexpected result %d", v)
	}
	if v := m.Name(); v != "RUSTURF TUNNEL" {
		t.Errorf("Name: unexpected result \"%s\"", v)
	}
}

func TestEncounters(t *testing.T) {
	ver := gen3.OpenROM(ROM(t))
	if ver == nil {
		t.Fatalf("failed to open ROM")
	}

	ver.ScanBanks()
	for i, e := range ver.BankByIndex(1).MapByIndex(0).Encounters() {
		if e.Populated() {
			t.Errorf("Encounters: unexpected populated EncounterList #%d (%s) of map 1.0", i, e.Name())
		}
	}
	{
		e := ver.BankByIndex(0).MapByIndex(0).Encounters()[0]
		if v := e.Name(); v != "Grass" {
			t.Errorf("EncounterList.Name: unexpected result \"%s\"", v)
		}
		if v := e.Populated(); v {
			t.Errorf("EncounterList.Populated: unexpected result %t", v)
		}
		if v := e.EncounterRate(); v != 0 {
			t.Errorf("EncounterList.EncounterRate: unexpected result %d", v)
		}
		if v := e.EncounterIndexSize(); v != 12 {
			t.Errorf("EncounterList.EncounterIndexSize: unexpected result %d", v)
		}
		if v := e.Encounters(); v != nil {
			t.Errorf("EncounterList.Encounters: unexpected non-nil result", v)
		}
		if v := e.Encounter(0); v != nil {
			t.Errorf("EncounterList.Encounter: unexpected non-nil result", v)
		}
		ExpectPanic(t, "EncounterList.Encounter", func() {
			e.Encounter(e.EncounterIndexSize())
		})
		if v := e.SpeciesRate(0); v != 0.2 {
			t.Errorf("EncounterList.SpeciesRate: unexpected result %g", v)
		}
	}

	m := ver.BankByIndex(0).MapByIndex(26)
	e := m.Encounters()
	if len(e) != 4 {
		t.Errorf("Encounters: unexpected result length %d", len(e))
	}
	if e, ok := e[0].(gen3.EncounterGrass); !ok {
		t.Errorf("type assertion of EncounterList #0 failed: expected gen3.EncounterGrass")
	} else {
		ExpectPanic(t, "EncounterGrass.SpeciesRate", func() {
			e.SpeciesRate(e.EncounterIndexSize())
		})
		if v := e.SpeciesRate(0); v != 0.2 {
			t.Errorf("EncounterGrass.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(2); v != 0.1 {
			t.Errorf("EncounterGrass.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(6); v != 0.05 {
			t.Errorf("EncounterGrass.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(8); v != 0.04 {
			t.Errorf("EncounterGrass.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(10); v != 0.01 {
			t.Errorf("EncounterGrass.SpeciesRate: unexpected result %g", v)
		}
	}
	if e, ok := e[1].(gen3.EncounterWater); !ok {
		t.Errorf("type assertion of EncounterList #1 failed: expected gen3.EncounterWater")
	} else {
		ExpectPanic(t, "EncounterWater.SpeciesRate", func() {
			e.SpeciesRate(e.EncounterIndexSize())
		})
		if v := e.SpeciesRate(0); v != 0.6 {
			t.Errorf("EncounterWater.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(1); v != 0.3 {
			t.Errorf("EncounterWater.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(2); v != 0.05 {
			t.Errorf("EncounterWater.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(3); v != 0.04 {
			t.Errorf("EncounterWater.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(4); v != 0.01 {
			t.Errorf("EncounterWater.SpeciesRate: unexpected result %g", v)
		}
	}
	if e, ok := e[2].(gen3.EncounterRock); !ok {
		t.Errorf("type assertion of EncounterList #2 failed: expected gen3.EncounterRock")
	} else {
		ExpectPanic(t, "EncounterRock.SpeciesRate", func() {
			e.SpeciesRate(e.EncounterIndexSize())
		})
		if v := e.SpeciesRate(0); v != 0.6 {
			t.Errorf("EncounterRock.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(1); v != 0.3 {
			t.Errorf("EncounterRock.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(2); v != 0.05 {
			t.Errorf("EncounterRock.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(3); v != 0.04 {
			t.Errorf("EncounterRock.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(4); v != 0.01 {
			t.Errorf("EncounterRock.SpeciesRate: unexpected result %g", v)
		}
	}
	if e, ok := e[3].(gen3.EncounterRod); !ok {
		t.Errorf("type assertion of EncounterList #3 failed: expected gen3.EncounterRod")
	} else {
		ExpectPanic(t, "EncounterRod.SpeciesRate", func() {
			e.SpeciesRate(e.EncounterIndexSize())
		})
		if v := e.SpeciesRate(0); v != 0.7 {
			t.Errorf("EncounterRod.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(1); v != 0.3 {
			t.Errorf("EncounterRod.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(2); v != 0.6 {
			t.Errorf("EncounterRod.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(3); v != 0.2 {
			t.Errorf("EncounterRod.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(4); v != 0.2 {
			t.Errorf("EncounterRod.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(5); v != 0.4 {
			t.Errorf("EncounterRod.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(6); v != 0.4 {
			t.Errorf("EncounterRod.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(7); v != 0.15 {
			t.Errorf("EncounterRod.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(8); v != 0.04 {
			t.Errorf("EncounterRod.SpeciesRate: unexpected result %g", v)
		}
		if v := e.SpeciesRate(9); v != 0.01 {
			t.Errorf("EncounterRod.SpeciesRate: unexpected result %g", v)
		}

		ExpectPanic(t, "EncounterRod.RodType", func() {
			e.RodType(e.EncounterIndexSize())
		})
		if v := e.RodType(0); v != gen3.OldRod {
			t.Errorf("EncounterRod.RodType: unexpected result %s", v)
		}
		if v := e.RodType(2); v != gen3.GoodRod {
			t.Errorf("EncounterRod.RodType: unexpected result %s", v)
		}
		if v := e.RodType(5); v != gen3.SuperRod {
			t.Errorf("EncounterRod.RodType: unexpected result %s", v)
		}
	}
	name := [4]string{"Grass", "Water", "Rock", "Rod"}
	size := [4]int{12, 5, 5, 10}
	erate := [4]byte{10, 4, 20, 30}
	srate := [4]float32{0.2, 0.6, 0.6, 0.7}
	for i, e := range e {
		if v := e.Name(); v != name[i] {
			t.Errorf("EncounterList.Name: %d: unexpected result \"%s\"", i, v)
		}
		if v := e.Populated(); !v {
			t.Errorf("EncounterList.Populated: %d: unexpected result %t", i, v)
		}
		if v := e.EncounterRate(); v != erate[i] {
			t.Errorf("EncounterList.EncounterRate: %d: unexpected result %d", i, v)
		}
		if v := e.EncounterIndexSize(); v != size[i] {
			t.Errorf("EncounterList.EncounterIndexSize: %d: unexpected result %d", i, v)
		}
		if v := e.Encounters(); len(v) != size[i] {
			t.Errorf("EncounterList.Encounters: %d: unexpected result length %d", i, len(v))
		}
		if v := e.Encounter(0); v == nil {
			t.Errorf("EncounterList.Encounter: %d: unexpected result <nil>", i)
		}
		ExpectPanic(t, "EncounterList.Encounter", func() {
			e.Encounter(e.EncounterIndexSize())
		})
		if v := e.SpeciesRate(0); v != srate[i] {
			t.Errorf("EncounterList.SpeciesRate: %d: unexpected result %g", i, v)
		}
	}
	{
		e := e[0].Encounter(0)
		if v := e.MinLevel(); v != 20 {
			t.Errorf("Encounter.MinLevel: unexpected result %d", v)
		}
		if v := e.MaxLevel(); v != 20 {
			t.Errorf("Encounter.MaxLevel: unexpected result %d", v)
		}
		if v := e.Species(); v.Index() != 27 {
			t.Errorf("Encounter.Species: unexpected result %d (%s)", v.Index(), v.Name())
		}
	}
}

func TestRod(t *testing.T) {
	if v := gen3.OldRod.String(); v != "Old" {
		t.Errorf("OldRod.String: unexpected result \"%s\"", v)
	}
	if v := gen3.GoodRod.String(); v != "Good" {
		t.Errorf("GoodRod.String: unexpected result \"%s\"", v)
	}
	if v := gen3.SuperRod.String(); v != "Super" {
		t.Errorf("SuperRod.String: unexpected result \"%s\"", v)
	}
}
