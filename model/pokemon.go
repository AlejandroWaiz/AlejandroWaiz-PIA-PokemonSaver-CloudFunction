package model

type Pokemon struct {
	ID              string
	Name            string
	Pasive          string
	Stage           int
	HP              int
	HpLevelBonus    string
	Speed           int
	SpeedLevelBonus string
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
