package gen3_test

import (
	"github.com/anaminus/pkm"
	"github.com/anaminus/pkm/gen3"
	"testing"
)

func TestSpecies(t *testing.T) {
	ver := gen3.OpenROM(ROM(t))
	if ver == nil {
		t.Fatalf("failed to open ROM")
	}

	species := ver.SpeciesByIndex(1)
	if v := species.Index(); v != 1 {
		t.Errorf("Index: unexpected result %d", v)
	}
	if v := species.Name(); v != "BULBASAUR" {
		t.Errorf("Name: unexpected result \"%s\"", v)
	}
	if v := species.Description(); v != "BULBASAUR can be seen napping in bright\nsunlight. There is a seed on its back.\nBy soaking up the sun’s rays, the seed\ngrows progressively larger." {
		t.Errorf("Description: unexpected result \"%s\"", v)
	}
	if v := species.Category(); v != "SEED" {
		t.Errorf("Category: unexpected result \"%s\"", v)
	}
	if v := species.Height(); v != 7 {
		t.Errorf("Height: unexpected result %d", v)
	}
	if v := species.Weight(); v != 7 {
		t.Errorf("Weight: unexpected result %d", v)
	}
	if v := species.BaseStats(); v != (pkm.Stats{45, 49, 49, 45, 65, 65}) {
		t.Errorf("BaseStats: unexpected result %#v", v)
	}
	if v := species.Type(); v != [2]pkm.Type{pkm.TypeGrass, pkm.TypePoison} {
		t.Errorf("Type: unexpected result %#v", v)
	}
	if v := species.CatchRate(); v != 45 {
		t.Errorf("CatchRate: unexpected result %d", v)
	}
	if v := species.ExpYield(); v != 64 {
		t.Errorf("ExpYield: unexpected result %d", v)
	}
	if v := species.EffortPoints(); v != 256 {
		t.Errorf("EffortPoints: unexpected result %d", v)
	}
	if v := species.HeldItem(); v != [2]pkm.Item{ver.ItemByIndex(0), ver.ItemByIndex(0)} {
		t.Errorf("HeldItem: unexpected result %#v", v)
	}
	if v := species.GenderRatio(); v != 31 {
		t.Errorf("GenderRatio: unexpected result %d", v)
	}
	if v := species.EggCycles(); v != 20 {
		t.Errorf("EggCycles: unexpected result %d", v)
	}
	if v := species.BaseFriendship(); v != 70 {
		t.Errorf("BaseFriendship: unexpected result %d", v)
	}
	if v := species.LevelType(); v != 3 {
		t.Errorf("LevelType: unexpected result %d", v)
	}
	if v := species.EggGroup(); v != [2]pkm.EggGroup{pkm.EggMonster, pkm.EggGrass} {
		t.Errorf("EggGroup: unexpected result %#v", v)
	}
	if v := species.Ability(); v != [2]pkm.Ability{ver.AbilityByIndex(65), ver.AbilityByIndex(0)} {
		t.Errorf("Ability: unexpected result %#v", v)
	}
	if v := species.SafariRate(); v != 0 {
		t.Errorf("SafariRate: unexpected result %d", v)
	}
	if v := species.Color(); v != pkm.ColorGreen {
		t.Errorf("Color: unexpected result %s", v)
	}
	{
		moves := []pkm.LevelMove{
			{Level: 1, Move: ver.MoveByName("TACKLE")},
			{Level: 4, Move: ver.MoveByName("GROWL")},
			{Level: 7, Move: ver.MoveByName("LEECH SEED")},
			{Level: 10, Move: ver.MoveByName("VINE WHIP")},
			{Level: 15, Move: ver.MoveByName("POISONPOWDER")},
			{Level: 15, Move: ver.MoveByName("SLEEP POWDER")},
			{Level: 20, Move: ver.MoveByName("RAZOR LEAF")},
			{Level: 25, Move: ver.MoveByName("SWEET SCENT")},
			{Level: 32, Move: ver.MoveByName("GROWTH")},
			{Level: 39, Move: ver.MoveByName("SYNTHESIS")},
			{Level: 46, Move: ver.MoveByName("SOLARBEAM")},
		}
		v := species.LearnedMoves()
		if len(v) != len(moves) {
			t.Errorf("LearnedMoves: unexpected result length %d", len(v))
		} else {
			for i, m := range v {
				if m.Level != moves[i].Level {
					t.Errorf("LearnedMoves: %d unexpected level %d", i, m.Level)
				}
				if m.Move != moves[i].Move {
					t.Errorf("LearnedMoves: %d unexpected move %d (%s)", i, m.Move.Index(), m.Move.Name())
				}
			}
		}
	}
	if v := species.CanLearnTM(ver.TMByIndex(1)); v {
		t.Errorf("CanLearnTM: unexpected result: %t", v)
	}
	if v := species.CanLearnTM(ver.TMByIndex(5)); !v {
		t.Errorf("CanLearnTM: unexpected result: %t", v)
	}
	{
		tms := []bool{
			0: false, 1: false, 2: false, 3: false, 4: false,
			5: true, 6: false, 7: false, 8: true, 9: true,
			10: true, 11: false, 12: false, 13: false, 14: false,
			15: false, 16: true, 17: false, 18: true, 19: false,
			20: true, 21: true, 22: false, 23: false, 24: false,
			25: false, 26: true, 27: false, 28: false, 29: false,
			30: false, 31: true, 32: false, 33: false, 34: false,
			35: true, 36: false, 37: false, 38: false, 39: false,
			40: false, 41: true, 42: true, 43: true, 44: true,
			45: false, 46: false, 47: false, 48: false, 49: false,
			50: true, 51: false, 52: false, 53: true, 54: true,
			55: true, 56: false, 57: false,
		}
		for _, tm := range species.LearnableTMs() {
			if !tms[tm.Index()] {
				t.Errorf("LearnableTMs: unexpected result %d (%s)", tm.Index(), tm.Name())
			}
		}
	}
	if evos := species.Evolutions(); len(evos) != 1 {
		t.Errorf("Evolutions: unexpected result length %d", len(evos))
	} else {
		if v := evos[0].Target(); v.Index() != 2 {
			t.Errorf("Evolutions: unexpected target %d (%s)", v.Index(), v.Name())
		}
		if v := evos[0].Method(); v != 4 {
			t.Errorf("Evolutions: unexpected method %d", v)
		}
		if v := evos[0].Param(); v != 16 {
			t.Errorf("Evolutions: unexpected param %d", v)
		}
	}
	species = ver.SpeciesByIndex(133)
	if evos := species.Evolutions(); len(evos) != 5 {
		t.Errorf("Evolutions: unexpected result length %d", len(evos))
	} else {
		target := [5]int{135, 134, 136, 196, 197}
		method := [5]uint16{7, 7, 7, 2, 3}
		param := [5]uint16{96, 97, 95, 0, 0}
		for i, evo := range evos {
			if v := evo.Target(); v.Index() != target[i] {
				t.Errorf("Evolutions: %d: unexpected target %d (%s)", i, v.Index(), v.Name())
			}
			if v := evo.Method(); v != method[i] {
				t.Errorf("Evolutions: %d: unexpected method %d", i, v)
			}
			if v := evo.Param(); v != param[i] {
				t.Errorf("Evolutions: %d: unexpected param %d", i, v)
			}
		}
	}
	if species, ok := species.(gen3.Species); ok {
		if v := species.Flipped(); v {
			t.Errorf("Flipped: unexpected result %t", v)
		}
	}
}
