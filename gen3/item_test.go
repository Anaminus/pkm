package gen3_test

import (
	"github.com/anaminus/pkm/gen3"
	"testing"
)

func TestItem(t *testing.T) {
	ver := gen3.OpenROM(ROM(t))
	if ver == nil {
		t.Fatalf("failed to open ROM")
	}

	item := ver.ItemByIndex(1)
	if v := item.Index(); v != 1 {
		t.Errorf("Index: unexpected result %d", v)
	}
	if v := item.Name(); v != "MASTER BALL" {
		t.Errorf("Name: unexpected result \"%s\"", v)
	}
	if v := item.Description(); v != "The best BALL that\ncatches a POKéMON\nwithout fail." {
		t.Errorf("Description: unexpected result \"%s\"", v)
	}
	if v := item.Price(); v != 0 {
		t.Errorf("Price: unexpected result %d", v)
	}

	item = ver.ItemByIndex(13)
	if v := item.Index(); v != 13 {
		t.Errorf("Index: unexpected result %d", v)
	}
	if v := item.Name(); v != "POTION" {
		t.Errorf("Name: unexpected result \"%s\"", v)
	}
	if v := item.Description(); v != "Restores the HP of\na POKéMON by\n20 points." {
		t.Errorf("Description: unexpected result \"%s\"", v)
	}
	if v := item.Price(); v != 300 {
		t.Errorf("Price: unexpected result %d", v)
	}
}
