package gen3_test

import (
	"bytes"
	"github.com/anaminus/pkm/gen3"
	"io"
	"os"
	"testing"
)

const ROMLocation = "rom.gba"

var rom io.ReadSeeker

func ROM(t *testing.T) io.ReadSeeker {
	if rom != nil {
		return rom
	}
	t.Logf("Note: tests require `%s` file in the current directory, whose contents are a ROM dump of Pokemon Emerald (BPEE)", ROMLocation)
	f, err := os.Open(ROMLocation)
	if err != nil {
		t.Fatal("failed to open ROM: %s", err)
	}
	defer f.Close()
	s, err := f.Seek(0, 2)
	if err != nil {
		t.Fatal("failed to seek ROM: %s", err)
	}
	b := make([]byte, s)
	_, err = f.Seek(0, 0)
	if err != nil {
		t.Fatal("failed to seek ROM: %s", err)
	}
	_, err = f.Read(b)
	if err != nil {
		t.Fatal("failed to read ROM: %s", err)
	}
	return bytes.NewReader(b)
}

func ExpectPanic(t *testing.T, s string, f func()) {
	defer func() {
		if v := recover(); v != nil {
			return
		}
	}()
	f()
	t.Errorf("%s: expected panic", s)
}

////////////////////////////////////////////////////////////////

func TestOpenROM(t *testing.T) {
	ver := gen3.OpenROM(ROM(t))
	if ver == nil {
		t.Fatalf("OpenROM: failed to open ROM")
	}
	if ver.GameCode() != gen3.CodeEmeraldEN {
		t.Fatalf("expected version game code `%s`, got `%s`", gen3.CodeEmeraldEN, ver.GameCode())
	}

	if gen3.OpenROM(bytes.NewReader([]byte{})) != nil {
		t.Fatalf("OpenROM: expected no Version")
	}
}
