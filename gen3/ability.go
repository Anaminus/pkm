package gen3

var (
	structAbilityName = makeStruct(
		13, // 0 Name
	)
)

type Ability struct {
	v Version
	i int
}

func (a Ability) Name() string {
	b := readStruct(
		a.v.ROM,
		a.v.AddrAbilityName,
		a.i,
		structAbilityName,
	)
	return decodeTextString(b)
}

func (a Ability) Index() int {
	return a.i
}

func (a Ability) Description() string {
	b := readStruct(
		a.v.ROM,
		a.v.AddrAbilityDescPtr,
		a.i,
		structPtr,
	)
	a.v.ROM.Seek(int64(decPtr(b)), 0)
	return readTextString(a.v.ROM)
}
