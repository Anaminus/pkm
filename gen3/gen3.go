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
	rom.Seek(addrGameCode, 0)
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
			{Name: "National", Size: 386, Address: 0xFFFFFF},
			{Name: "Standard", Size: 202, Address: 0xFFFFFF},
		},
		AddrAbilityName:    0xFFFFFF,
		AddrAbilityDescPtr: 0xFFFFFF,
		AddrBanksPtr:       0xFFFFFF,
		AddrItemData:       0xFFFFFF,
		AddrLevelMovePtr:   0xFFFFFF,
		AddrMapLabel:       0xFFFFFF,
		AddrMoveName:       0xFFFFFF,
		AddrMoveData:       0xFFFFFF,
		AddrMoveDescPtr:    0xFFFFFF,
		AddrPokedexData:    0xFFFFFF,
		AddrSpeciesData:    0xFFFFFF,
		AddrSpeciesEvo:     0xFFFFFF,
		AddrSpeciesName:    0xFFFFFF,
		AddrSpeciesTM:      0xFFFFFF,
		AddrTMMove:         0xFFFFFF,
	},
	CodeSapphireEN: Version{
		name: "Pokémon Sapphire Version",
		pokedex: []pokedexData{
			{Name: "National", Size: 386, Address: 0xFFFFFF},
			{Name: "Standard", Size: 202, Address: 0xFFFFFF},
		},
		AddrAbilityName:    0xFFFFFF,
		AddrAbilityDescPtr: 0xFFFFFF,
		AddrBanksPtr:       0xFFFFFF,
		AddrItemData:       0xFFFFFF,
		AddrLevelMovePtr:   0xFFFFFF,
		AddrMapLabel:       0xFFFFFF,
		AddrMoveName:       0xFFFFFF,
		AddrMoveData:       0xFFFFFF,
		AddrMoveDescPtr:    0xFFFFFF,
		AddrPokedexData:    0xFFFFFF,
		AddrSpeciesData:    0xFFFFFF,
		AddrSpeciesEvo:     0xFFFFFF,
		AddrSpeciesName:    0xFFFFFF,
		AddrSpeciesTM:      0xFFFFFF,
		AddrTMMove:         0xFFFFFF,
	},
	CodeEmeraldEN: Version{
		name: "Pokémon Emerald Version",
		pokedex: []pokedexData{
			{Name: "National", Size: 386, Address: 0x31DC82},
			{Name: "Standard", Size: 202, Address: 0x31D94C},
		},
		AddrAbilityName:    0x31B6DB,
		AddrAbilityDescPtr: 0x31BAD4,
		AddrBanksPtr:       0x084AA4,
		AddrItemData:       0x5839A0,
		AddrLevelMovePtr:   0x32937C,
		AddrMapLabel:       0x5A147C,
		AddrMoveName:       0x31977C,
		AddrMoveData:       0x31C898,
		AddrMoveDescPtr:    0x61C524,
		AddrPokedexData:    0x56B5B0,
		AddrSpeciesData:    0x3203CC,
		AddrSpeciesEvo:     0x32531C,
		AddrSpeciesName:    0x3185C8,
		AddrSpeciesTM:      0x31E898,
		AddrTMMove:         0x616040,
	},
	CodeFireRedEN: Version{
		name: "Pokémon Fire Red Version",
		pokedex: []pokedexData{
			{Name: "National", Size: 386, Address: 0xFFFFFF},
			{Name: "Standard", Size: 151, Address: 0xFFFFFF},
		},
		AddrAbilityName:    0xFFFFFF,
		AddrAbilityDescPtr: 0xFFFFFF,
		AddrBanksPtr:       0xFFFFFF,
		AddrItemData:       0xFFFFFF,
		AddrLevelMovePtr:   0xFFFFFF,
		AddrMapLabel:       0xFFFFFF,
		AddrMoveName:       0xFFFFFF,
		AddrMoveData:       0xFFFFFF,
		AddrMoveDescPtr:    0xFFFFFF,
		AddrPokedexData:    0xFFFFFF,
		AddrSpeciesData:    0xFFFFFF,
		AddrSpeciesEvo:     0xFFFFFF,
		AddrSpeciesName:    0xFFFFFF,
		AddrSpeciesTM:      0xFFFFFF,
		AddrTMMove:         0xFFFFFF,
	},
	CodeLeafGreenEN: Version{
		name: "Pokémon Leaf Green Version",
		pokedex: []pokedexData{
			{Name: "National", Size: 386, Address: 0xFFFFFF},
			{Name: "Standard", Size: 151, Address: 0xFFFFFF},
		},
		AddrAbilityName:    0xFFFFFF,
		AddrAbilityDescPtr: 0xFFFFFF,
		AddrBanksPtr:       0xFFFFFF,
		AddrItemData:       0xFFFFFF,
		AddrLevelMovePtr:   0xFFFFFF,
		AddrMapLabel:       0xFFFFFF,
		AddrMoveName:       0xFFFFFF,
		AddrMoveData:       0xFFFFFF,
		AddrMoveDescPtr:    0xFFFFFF,
		AddrPokedexData:    0xFFFFFF,
		AddrSpeciesData:    0xFFFFFF,
		AddrSpeciesEvo:     0xFFFFFF,
		AddrSpeciesName:    0xFFFFFF,
		AddrSpeciesTM:      0xFFFFFF,
		AddrTMMove:         0xFFFFFF,
	},
}
