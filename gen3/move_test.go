package gen3_test

import (
	"github.com/anaminus/pkm"
	"github.com/anaminus/pkm/gen3"
	"testing"
)

func TestMove(t *testing.T) {
	ver := gen3.OpenROM(ROM(t))
	if ver == nil {
		t.Fatalf("failed to open ROM")
	}

	move := ver.MoveByIndex(1)
	if v := move.Index(); v != 1 {
		t.Errorf("Index: unexpected result \"%s\"", v)
	}
	if v := move.Name(); v != "POUND" {
		t.Errorf("Name: unexpected result \"%s\"", v)
	}
	if v := move.Description(); v != "Pounds the foe with\nforelegs or tail." {
		t.Errorf("Description: unexpected result \"%s\"", v)
	}
	if v := move.Type(); v != pkm.TypeNormal {
		t.Errorf("Type: unexpected result \"%s\"", v)
	}
	if v := move.BasePower(); v != 40 {
		t.Errorf("BasePower: unexpected result %d", v)
	}
	if v := move.Accuracy(); v != 100 {
		t.Errorf("Accuracy: unexpected result %d", v)
	}
	if v := move.Effect(); v != 0 {
		t.Errorf("Effect: unexpected result %d", v)
	}
	if v := move.EffectAccuracy(); v != 0 {
		t.Errorf("EffectAccuracy: unexpected result %d", v)
	}
	if v := move.Affectee(); v != 0 {
		t.Errorf("Affectee: unexpected result %d", v)
	}
	if v := move.Priority(); v != 0 {
		t.Errorf("Priority: unexpected result %d", v)
	}
	if v := move.Flags(); v != pkm.Contact|pkm.Protect|pkm.MirrorMove|pkm.KingsRock {
		t.Errorf("Flags: unexpected result %s", v)
	}
}

func TestTM(t *testing.T) {
	ver := gen3.OpenROM(ROM(t))
	if ver == nil {
		t.Fatalf("failed to open ROM")
	}

	tm := ver.TMByIndex(0)
	if v := tm.Index(); v != 0 {
		t.Errorf("Index: unexpected result %d", v)
	}
	if v := tm.Name(); v != "TM01" {
		t.Errorf("Name: unexpected result \"%s\"", v)
	}
	if v := tm.Move(); v.Index() != 264 {
		t.Errorf("Move: unexpected result %d", v.Index())
	}

	tm = ver.TMByIndex(35)
	if v := tm.Index(); v != 35 {
		t.Errorf("Index: unexpected result %d", v)
	}
	if v := tm.Name(); v != "TM36" {
		t.Errorf("Name: unexpected result \"%s\"", v)
	}
	if v := tm.Move(); v.Index() != 188 {
		t.Errorf("Move: unexpected result %d", v.Index())
	}

	tm = ver.TMByIndex(50)
	if v := tm.Index(); v != 50 {
		t.Errorf("Index: unexpected result %d", v)
	}
	if v := tm.Name(); v != "HM01" {
		t.Errorf("Name: unexpected result \"%s\"", v)
	}
	if v := tm.Move(); v.Index() != 15 {
		t.Errorf("Move: unexpected result %d", v.Index())
	}
}
