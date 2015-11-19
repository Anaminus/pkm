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
	// TODO
	return ""
}

func (a Ability) Index() int {
	// TODO
	return a.i
}

func (a Ability) Description() string {
	// TODO
	return ""
}
