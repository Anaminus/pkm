package gen3_test

import (
	"github.com/anaminus/pkm/gen3"
	"testing"
)

func TestPokedex(t *testing.T) {
	ver := gen3.OpenROM(ROM(t))
	if ver == nil {
		t.Fatalf("failed to open ROM")
	}

	dex := ver.Pokedex()[0]
	if v := dex.Name(); v != "National" {
		t.Errorf("Name: unexpected result \"%s\"", v)
	}
	if v := dex.Size(); v != 386 {
		t.Errorf("Size: unexpected result \"%s\"", v)
	}

	type sp struct {
		name string
		num  int
	}
	for _, s := range []sp{{"Bulbasaur", 1}, {"Mew", 151}, {"Treecko", 252}, {"Deoxys", 386}} {
		species := ver.SpeciesByName(s.name)
		if v := dex.Species(s.num); v != species {
			if v == nil {
				t.Errorf("Species: %s: unexpected result <nil>", species.Name())
			} else {
				t.Errorf("Species: %s: unexpected result %d (%s)", species.Name(), v.Index(), v.Name())
			}
		}
		if v := dex.SpeciesNumber(species); v != s.num {
			t.Errorf("SpeciesNumber: %s: unexpected result %d", species.Name(), v)
		}
	}

	ExpectPanic(t, "Species", func() {
		dex.Species(-1)
	})
	ExpectPanic(t, "Species", func() {
		dex.Species(0)
	})
	ExpectPanic(t, "Species", func() {
		dex.Species(dex.Size() + 1)
	})

	if v := dex.AllSpecies(); len(v) != dex.Size() {
		t.Errorf("AllSpecies: unexpected length of result %d", len(v))
	}

}
