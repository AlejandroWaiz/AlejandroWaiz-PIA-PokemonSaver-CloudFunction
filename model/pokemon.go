package model

type Pokemon struct {
	ID              interface{}
	Name            interface{}
	Type            interface{}
	Pasive          interface{}
	Stage           interface{}
	HP              interface{}
	HpLevelBonus    interface{}
	Speed           interface{}
	SpeedLevelBonus interface{}
	Movements       []map[string]string
	Proficiencies
}

type Proficiencies struct {
	AttackProficiency, SpecialAttackProficiency, DefenseProficiency, SpecialDefenseProficiency, SpeedProficiency string
}

var emptyPokemon = &Pokemon{}

func (p *Pokemon) ResetStruct() {

	*p = *emptyPokemon

}
