package gen3_test

import (
	"github.com/anaminus/pkm/gen3"
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	ver := gen3.OpenROM(ROM(t))
	if ver == nil {
		t.Fatalf("failed to open ROM")
	}

	if !strings.Contains(ver.Name(), "Emerald") {
		t.Errorf("Name: unexpected name `%s`", ver.Name())
	}
	if ver.GameCode() != gen3.CodeEmeraldEN {
		t.Errorf("GameCode: got game code %s, expected %s", ver.GameCode(), gen3.CodeEmeraldEN)
	}
	if ver.Query() == nil {
		t.Errorf("Query: expected Query")
	}
	if len(ver.Codecs()) == 0 {
		t.Errorf("Codecs: expected at least one codec")
	}
	if ver.DefaultCodec() == nil {
		t.Errorf("DefaultCodec: expected default codec")
	}
	if ver.Codecs()[0] != ver.DefaultCodec() {
		t.Errorf("Codecs: first codec should be default codec")
	}

	ExpectPanic(t, "SpeciesByIndex", func() {
		ver.SpeciesByIndex(-1)
	})
	ExpectPanic(t, "SpeciesByIndex", func() {
		ver.SpeciesByIndex(ver.SpeciesIndexSize())
	})
	ExpectPanic(t, "SpeciesByIndex", func() {
		ver.SpeciesByIndex(ver.SpeciesIndexSize() + 1)
	})
	for i := 0; i < ver.SpeciesIndexSize(); i++ {
		s := ver.SpeciesByIndex(i)
		if s.Index() != i {
			t.Errorf("SpeciesByIndex: returned index %d, expected %d", s.Index(), i)
			break
		}
	}
	if ver.SpeciesByName("BULBASAUR") == nil {
		t.Errorf("SpeciesByName: returned nil species")
	} else {
		if ver.SpeciesByName("BULBASAUR").Index() != 1 {
			t.Errorf("SpeciesByName: expected index 1")
		}
		if ver.SpeciesByName("Bulbasaur") == nil {
			t.Errorf("SpeciesByName: not case-insensitive")
		}
	}
	if ver.SpeciesByName("") != nil {
		t.Errorf("SpeciesByName: expected nil")
	}
	if ver.SpeciesByName(ver.SpeciesByIndex(ver.SpeciesIndexSize()-1).Name()) == nil {
		t.Errorf("SpeciesByName: returned nil last species")
	}

	if len(ver.Pokedex()) == 0 {
		t.Errorf("Pokedex: expected at least one pokedex")
	}
	if ver.PokedexByName("National") == nil {
		t.Errorf("PokedexByName: expected pokedex")
	}
	if ver.PokedexByName("Standard") == nil {
		t.Errorf("PokedexByName: expected pokedex")
	}
	if ver.PokedexByName("") != nil {
		t.Errorf("PokedexByName: unexpected pokedex")
	}

	if len(ver.Items()) != ver.ItemIndexSize() {
		t.Errorf("Items: expected length %d, got %d", ver.ItemIndexSize(), len(ver.Items()))
	}
	ExpectPanic(t, "ItemByIndex", func() {
		ver.ItemByIndex(-1)
	})
	ExpectPanic(t, "ItemByIndex", func() {
		ver.ItemByIndex(ver.ItemIndexSize())
	})
	ExpectPanic(t, "ItemByIndex", func() {
		ver.ItemByIndex(ver.ItemIndexSize() + 1)
	})
	for i := 0; i < ver.ItemIndexSize(); i++ {
		s := ver.ItemByIndex(i)
		if s.Index() != i {
			t.Errorf("ItemByIndex: returned index %d, expected %d", s.Index(), i)
			break
		}
	}
	if ver.ItemByName("MASTER BALL") == nil {
		t.Errorf("ItemByName: returned nil item")
	} else {
		if ver.ItemByName("MASTER BALL").Index() != 1 {
			t.Errorf("ItemByName: expected index 1")
		}
		if ver.ItemByName("Master Ball") == nil {
			t.Errorf("ItemByName: not case-insensitive")
		}
	}
	if ver.ItemByName("") != nil {
		t.Errorf("ItemByName: expected nil")
	}
	if ver.ItemByName(ver.ItemByIndex(ver.ItemIndexSize()-1).Name()) == nil {
		t.Errorf("ItemByName: returned nil last species")
	}

	if len(ver.Abilities()) != ver.AbilityIndexSize() {
		t.Errorf("Abilities: expected length %d, got %d", ver.AbilityIndexSize(), len(ver.Abilities()))
	}
	ExpectPanic(t, "AbilityByIndex", func() {
		ver.AbilityByIndex(-1)
	})
	ExpectPanic(t, "AbilityByIndex", func() {
		ver.AbilityByIndex(ver.AbilityIndexSize())
	})
	ExpectPanic(t, "AbilityByIndex", func() {
		ver.AbilityByIndex(ver.AbilityIndexSize() + 1)
	})
	for i := 0; i < ver.AbilityIndexSize(); i++ {
		s := ver.AbilityByIndex(i)
		if s.Index() != i {
			t.Errorf("AbilityByIndex: returned index %d, expected %d", s.Index(), i)
			break
		}
	}
	if ver.AbilityByName("STENCH") == nil {
		t.Errorf("AbilityByName: returned nil item")
	} else {
		if ver.AbilityByName("STENCH").Index() != 1 {
			t.Errorf("AbilityByName: expected index 1")
		}
		if ver.AbilityByName("Stench") == nil {
			t.Errorf("AbilityByName: not case-insensitive")
		}
	}
	if ver.AbilityByName("") != nil {
		t.Errorf("AbilityByName: expected nil")
	}
	if ver.AbilityByName(ver.AbilityByIndex(ver.AbilityIndexSize()-1).Name()) == nil {
		t.Errorf("AbilityByName: returned nil last species")
	}

	if len(ver.Moves()) != ver.MoveIndexSize() {
		t.Errorf("Moves: expected length %d, got %d", ver.MoveIndexSize(), len(ver.Moves()))
	}
	ExpectPanic(t, "MoveByIndex", func() {
		ver.MoveByIndex(-1)
	})
	ExpectPanic(t, "MoveByIndex", func() {
		ver.MoveByIndex(ver.MoveIndexSize())
	})
	ExpectPanic(t, "MoveByIndex", func() {
		ver.MoveByIndex(ver.MoveIndexSize() + 1)
	})
	for i := 0; i < ver.MoveIndexSize(); i++ {
		s := ver.MoveByIndex(i)
		if s.Index() != i {
			t.Errorf("MoveByIndex: returned index %d, expected %d", s.Index(), i)
			break
		}
	}
	if ver.MoveByName("POUND") == nil {
		t.Errorf("MoveByName: returned nil item")
	} else {
		if ver.MoveByName("POUND").Index() != 1 {
			t.Errorf("MoveByName: expected index 1")
		}
		if ver.MoveByName("Pound") == nil {
			t.Errorf("MoveByName: not case-insensitive")
		}
	}
	if ver.MoveByName("") != nil {
		t.Errorf("MoveByName: expected nil")
	}
	if ver.MoveByName(ver.MoveByIndex(ver.MoveIndexSize()-1).Name()) == nil {
		t.Errorf("MoveByName: returned nil last species")
	}

	if len(ver.TMs()) != ver.TMIndexSize() {
		t.Errorf("TMs: expected length %d, got %d", ver.TMIndexSize(), len(ver.TMs()))
	}
	ExpectPanic(t, "TMByIndex", func() {
		ver.TMByIndex(-1)
	})
	ExpectPanic(t, "TMByIndex", func() {
		ver.TMByIndex(ver.TMIndexSize())
	})
	ExpectPanic(t, "TMByIndex", func() {
		ver.TMByIndex(ver.TMIndexSize() + 1)
	})
	for i := 0; i < ver.TMIndexSize(); i++ {
		s := ver.TMByIndex(i)
		if s.Index() != i {
			t.Errorf("TMByIndex: returned index %d, expected %d", s.Index(), i)
			break
		}
	}
	if ver.TMByName("TM01") == nil {
		t.Errorf("TMByName: returned nil item")
	} else {
		if ver.TMByName("TM01").Index() != 0 {
			t.Errorf("TMByName: expected index 1")
		}
		if ver.TMByName("tm01") == nil {
			t.Errorf("TMByName: not case-insensitive")
		}
	}
	if ver.TMByName("") != nil {
		t.Errorf("TMByName: expected nil")
	}
	if ver.TMByName("AAAA") != nil {
		t.Errorf("TMByName: expected nil")
	}
	if ver.TMByName("MM01") != nil {
		t.Errorf("TMByName: expected nil")
	}
	if ver.TMByName("TM99") != nil {
		t.Errorf("TMByName: expected nil")
	}
	if ver.TMByName("AAAAA") != nil {
		t.Errorf("TMByName: expected nil")
	}
	if ver.TMByName(ver.TMByIndex(ver.TMIndexSize()-1).Name()) == nil {
		t.Errorf("TMByName: returned nil last species")
	}

	ExpectPanic(t, "BankIndexSize", func() {
		ver.BankIndexSize()
	})
	ExpectPanic(t, "Banks", func() {
		ver.Banks()
	})
	ExpectPanic(t, "BankByIndex", func() {
		ver.BankByIndex(0)
	})
	ExpectPanic(t, "AllMaps", func() {
		ver.AllMaps()
	})
	ExpectPanic(t, "MapByName", func() {
		ver.MapByName("")
	})

	ver.ScanBanks()

	ExpectPanic(t, "BankByIndex", func() {
		ver.BankByIndex(257)
	})

	if v := ver.BankIndexSize(); v != 34 {
		t.Errorf("BankIndexSize: unexpected result %d", v)
	}
	if v := ver.Banks(); len(v) != 34 {
		t.Errorf("Banks: unexpected result length %d", len(v))
	}
	if v := ver.BankByIndex(0); v == nil || v.Index() != 0 {
		t.Errorf("BankByIndex: unexpected result %d", v.Index)
	}
	if v := ver.AllMaps(); len(v) != 518 {
		t.Errorf("AllMaps: unexpected result length %d", v)
	}
	if m := ver.MapByName("RUSTURF TUNNEL"); m == nil {
		t.Errorf("MapByName: unexpected result <nil>")
	} else if m.BankIndex() != 24 || m.Index() != 4 {
		t.Errorf("MapByName: unexpected result %d.%d (%s)", m.BankIndex(), m.Index(), m.Name())
	} else {
		if v := ver.MapByName("Rusturf Tunnel"); v != m {
			t.Errorf("MapByName: not case-insensitive")
		}
	}
	if v := ver.MapByName("unknown"); v != nil {
		t.Errorf("MapByName: expected nil result")
		t.Logf("Result: %d.%d", v.BankIndex(), v.Index())
	}
}
