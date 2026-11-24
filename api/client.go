package api

import (
	"json"
)

type Ammos struct {
	id string,
	name string,
	image string,
	description string,
	type string,
	passive string,
	attackPower {name: string, amount: int},
}
type Armors struct {
	id string,
	name string,
	image string,
	description string,
	category string,
	weight int,
	dmgNegation {name: string, amount: int},
	resistance {name: string, amount: int},
}
type Ashes of War struct {
	id string,
	name string,
	image string,
	description string,
	affinity string,
	skill string,
}
type Bosses struct {
	id string,
	name string,
	image string,
	description string,
	location string,
	drops string[],
	healthPoints string,
}
type Classes struct {
	id string,
	name string,
	image string,
	description string,
	stats { level: int, vigor: int, mind: int, endurance: int, strength: int, dexterity: int, intelligence: int, faith: int, arcane: int },
}
type Creatures struct {
	id string,
	name string,
	image string,
	description string,
	location string,
	drops string[],
}
type Incantations struct {
	id string,
	name string,
	image string,
	description string,
	type string,
	cost int,
	slots int,
	effects string,
	requirements { name: string, amount: int},
}
type Items struct {
	id string,
	name string,
	image string,
	description string,
	type string,
	effect string,
}
type Locations struct {
	id string,
	name string,
	image string,
	description string,
}
type NPCs struct {
	id string,
	name string,
	image string,
	description string,
	location string,
	quote string,
}
type Shields struct {
	id string,
	name string,
	image string,
	description string,
	category string,
	weight int,
	attack {name: string, amount: int},
	defense {name: string, amount: int},
	requiredAttributes {name: string, amount: int},
	scalesWith {name: string, scaling: string},
}