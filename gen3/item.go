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
	// TODO
	return ""
}

func (i Item) Index() int {
	// TODO
	return 0
}

func (i Item) Description() string {
	// TODO
	return ""
}

func (i Item) Price() int {
	// TODO
	return 0
}
