package gen3

import (
	"github.com/anaminus/pkm"
	"io"
)

// OpemROM creates a pkm.Version that reads a GameBoy Advance ROM file. If the
// contents are identified as an unsupported version, then a nil value is
// returned.
func OpenROM(rom io.ReadSeeker) pkm.Version {
	var gc pkm.GameCode
	rom.Seek(addrGameCode.ROM(), 0)
	rom.Read(gc[:])
	if v, ok := versionLookup[gc]; ok {
		v.ROM = rom
		return &v
	}
	return nil
}

// Game codes that identify generation III versions.
var (
	// English
	CodeRubyEN      = pkm.GameCode{'A', 'X', 'V', 'E'}
	CodeSapphireEN  = pkm.GameCode{'A', 'X', 'P', 'E'}
	CodeEmeraldEN   = pkm.GameCode{'B', 'P', 'E', 'E'}
	CodeFireRedEN   = pkm.GameCode{'B', 'P', 'R', 'E'}
	CodeLeafGreenEN = pkm.GameCode{'B', 'P', 'G', 'E'}
)

var versionLookup = map[pkm.GameCode]Version{
	CodeRubyEN: Version{
		name: "Pokémon Ruby Version",
		pokedex: []pokedexData{
			{Name: "National", Size: 386, Address: 0xFFFFFFFF},
			{Name: "Standard", Size: 202, Address: 0xFFFFFFFF},
		},
		AddrAbilityName:    0xFFFFFFFF,
		AddrAbilityDescPtr: 0xFFFFFFFF,
		AddrBanksPtr:       0xFFFFFFFF,
		AddrEncounterList:  0xFFFFFFFF,
		AddrItemData:       0xFFFFFFFF,
		AddrLevelMovePtr:   0xFFFFFFFF,
		AddrMapLabel:       0xFFFFFFFF,
		AddrMoveName:       0xFFFFFFFF,
		AddrMoveData:       0xFFFFFFFF,
		AddrMoveDescPtr:    0xFFFFFFFF,
		AddrPokedexData:    0xFFFFFFFF,
		AddrSpeciesData:    0xFFFFFFFF,
		AddrSpeciesEvo:     0xFFFFFFFF,
		AddrSpeciesName:    0xFFFFFFFF,
		AddrSpeciesTM:      0xFFFFFFFF,
		AddrTMMove:         0xFFFFFFFF,
	},
	CodeSapphireEN: Version{
		name: "Pokémon Sapphire Version",
		pokedex: []pokedexData{
			{Name: "National", Size: 386, Address: 0xFFFFFFFF},
			{Name: "Standard", Size: 202, Address: 0xFFFFFFFF},
		},
		AddrAbilityName:    0xFFFFFFFF,
		AddrAbilityDescPtr: 0xFFFFFFFF,
		AddrBanksPtr:       0xFFFFFFFF,
		AddrEncounterList:  0xFFFFFFFF,
		AddrItemData:       0xFFFFFFFF,
		AddrLevelMovePtr:   0xFFFFFFFF,
		AddrMapLabel:       0xFFFFFFFF,
		AddrMoveName:       0xFFFFFFFF,
		AddrMoveData:       0xFFFFFFFF,
		AddrMoveDescPtr:    0xFFFFFFFF,
		AddrPokedexData:    0xFFFFFFFF,
		AddrSpeciesData:    0xFFFFFFFF,
		AddrSpeciesEvo:     0xFFFFFFFF,
		AddrSpeciesName:    0xFFFFFFFF,
		AddrSpeciesTM:      0xFFFFFFFF,
		AddrTMMove:         0xFFFFFFFF,
	},
	CodeEmeraldEN: Version{
		name: "Pokémon Emerald Version",
		pokedex: []pokedexData{
			{Name: "National", Size: 386, Address: 0x0831DC82},
			{Name: "Standard", Size: 202, Address: 0x0831D94C},
		},
		AddrAbilityName:    0x0831B6DB,
		AddrAbilityDescPtr: 0x0831BAD4,
		AddrBanksPtr:       0x08084AA4,
		AddrEncounterList:  0x08552D48,
		AddrItemData:       0x085839A0,
		AddrLevelMovePtr:   0x0832937C,
		AddrMapLabel:       0x085A147C,
		AddrMoveName:       0x0831977C,
		AddrMoveData:       0x0831C898,
		AddrMoveDescPtr:    0x0861C524,
		AddrPokedexData:    0x0856B5B0,
		AddrSpeciesData:    0x083203CC,
		AddrSpeciesEvo:     0x0832531C,
		AddrSpeciesName:    0x083185C8,
		AddrSpeciesTM:      0x0831E898,
		AddrTMMove:         0x08616040,
	},
	CodeFireRedEN: Version{
		name: "Pokémon Fire Red Version",
		pokedex: []pokedexData{
			{Name: "National", Size: 386, Address: 0xFFFFFFFF},
			{Name: "Standard", Size: 151, Address: 0xFFFFFFFF},
		},
		AddrAbilityName:    0xFFFFFFFF,
		AddrAbilityDescPtr: 0xFFFFFFFF,
		AddrBanksPtr:       0xFFFFFFFF,
		AddrEncounterList:  0xFFFFFFFF,
		AddrItemData:       0xFFFFFFFF,
		AddrLevelMovePtr:   0xFFFFFFFF,
		AddrMapLabel:       0xFFFFFFFF,
		AddrMoveName:       0xFFFFFFFF,
		AddrMoveData:       0xFFFFFFFF,
		AddrMoveDescPtr:    0xFFFFFFFF,
		AddrPokedexData:    0xFFFFFFFF,
		AddrSpeciesData:    0xFFFFFFFF,
		AddrSpeciesEvo:     0xFFFFFFFF,
		AddrSpeciesName:    0xFFFFFFFF,
		AddrSpeciesTM:      0xFFFFFFFF,
		AddrTMMove:         0xFFFFFFFF,
	},
	CodeLeafGreenEN: Version{
		name: "Pokémon Leaf Green Version",
		pokedex: []pokedexData{
			{Name: "National", Size: 386, Address: 0xFFFFFFFF},
			{Name: "Standard", Size: 151, Address: 0xFFFFFFFF},
		},
		AddrAbilityName:    0xFFFFFFFF,
		AddrAbilityDescPtr: 0xFFFFFFFF,
		AddrBanksPtr:       0xFFFFFFFF,
		AddrEncounterList:  0xFFFFFFFF,
		AddrItemData:       0xFFFFFFFF,
		AddrLevelMovePtr:   0xFFFFFFFF,
		AddrMapLabel:       0xFFFFFFFF,
		AddrMoveName:       0xFFFFFFFF,
		AddrMoveData:       0xFFFFFFFF,
		AddrMoveDescPtr:    0xFFFFFFFF,
		AddrPokedexData:    0xFFFFFFFF,
		AddrSpeciesData:    0xFFFFFFFF,
		AddrSpeciesEvo:     0xFFFFFFFF,
		AddrSpeciesName:    0xFFFFFFFF,
		AddrSpeciesTM:      0xFFFFFFFF,
		AddrTMMove:         0xFFFFFFFF,
	},
}
