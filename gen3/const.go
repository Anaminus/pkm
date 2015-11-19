package gen3

const (
	indexSizeSpecies = 412
	indexSizeItem    = 377
	indexSizeAbility = 78
	indexSizeMove    = 355
	indexSizeTM      = 58
	indexSizeBank    = 34
	indexSizeMapName = 213
)

var mapIndexSize = [indexSizeBank]int{
	56, 4, 4, 5, 6, 7, 8, 6,
	6, 13, 7, 16, 9, 22, 12, 14,
	14, 1, 1, 1, 2, 0, 0, 0,
	107, 60, 88, 1, 0, 12, 0, 0,
	2, 0,
}

const (
	speciesNameLength = 11
	itemNameLength    = 14
	itemDataLength    = 44
	abilityNameLength = 13
	moveNameLength    = 13
)
