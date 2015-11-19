package gen3

var (
	structItemData = makeStruct(
		14, // 00 Name
		2,  // 01 Index
		2,  // 02 Price
		1,  // 03 HoldEffect
		1,  // 04 Parameter
		4,  // 05 DescPtr
		2,  // 06 Mystery
		1,  // 07 Pocket
		1,  // 08 Type
		4,  // 09 FieldCodePtr
		4,  // 10 BattleUsage
		4,  // 11 BattleCodePtr
		4,  // 12 ExtraParam
	)
)

type Item struct {
	v Version
	i int
}

func (i Item) Name() string {
	b := readStruct(
		i.v.ROM,
		i.v.AddrItemData,
		i.i,
		structItemData,
		0,
	)
	return decodeTextString(b)
}

func (i Item) Index() int {
	return i.i
}

func (i Item) Description() string {
	b := readStruct(
		i.v.ROM,
		i.v.AddrItemData,
		i.i,
		structItemData,
		5,
	)
	i.v.ROM.Seek(int64(i.v.AddrItemDesc)+int64(decUint32(b)), 0)
	return decodeTextString(b)
}

func (i Item) Price() int {
	b := readStruct(
		i.v.ROM,
		i.v.AddrItemData,
		i.i,
		structItemData,
		2,
	)
	return int(decUint16(b))
}
