package gen3

import (
	"github.com/anaminus/pkm"
)

type Query struct {
	v Version
}

func (q Query) SpeciesByName(name string) pkm.Species {
	// TODO
	return nil
}

func (q Query) ItemByName(name string) pkm.Item {
	// TODO
	return nil
}

func (q Query) AbilityByName(name string) pkm.Ability {
	// TODO
	return nil
}

func (q Query) MoveByName(name string) pkm.Move {
	// TODO
	return nil
}

func (q Query) TMByName(name string) pkm.TM {
	// TODO
	return nil
}

func (q Query) MapByName(name string) pkm.Map {
	// TODO
	return nil
}

func (q Query) SpeciesLearningMove(move pkm.Move) []pkm.Species {
	// TODO
	return nil
}

func (q Query) SpeciesLearningTM(tm pkm.TM) []pkm.Species {
	// TODO
	return nil
}

func (q Query) SpeciesLocations(species pkm.Species) []pkm.Map {
	// TODO
	return nil
}
