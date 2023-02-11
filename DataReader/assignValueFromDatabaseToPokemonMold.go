package datareader

import (
	"github.com/AlejandroWaiz/PIA-PokemonSaver-CloudFunction/model"
)

var mapOfPokemonMovementsFromDatabase map[string]string

func assignValueFromDatabaseToPokemonMold(pokemon model.Pokemon, columnPosition int, columnValue string) error {

	if mapOfPokemonMovementsFromDatabase == nil {
		mapOfPokemonMovementsFromDatabase = make(map[string]string)
	}

	if columnPosition == 0 {
		pokemonMold.ID = columnValue
	} else if columnPosition == 1 {
		pokemonMold.Name = columnValue
	} else if columnPosition == 2 {
		pokemonMold.Type = columnValue
	} else if columnPosition == 3 {
		pokemonMold.Pasive = columnValue
	} else if columnPosition == 4 {
		pokemonMold.Stage = columnValue
	} else if columnPosition == 5 {
		pokemonMold.HP = columnValue
	} else if columnPosition == 6 {
		pokemonMold.HpLevelBonus = columnValue
	} else if columnPosition == 7 {
		pokemonMold.Speed = columnValue
	} else if columnPosition == 8 {
		pokemonMold.SpeedLevelBonus = columnValue
	} else if columnPosition == 9 {
		mapOfPokemonMovementsFromDatabase["Base movements"] = string(columnValue)
	} else if columnPosition == 10 {
		mapOfPokemonMovementsFromDatabase["Level one movements"] = string(columnValue)
	} else if columnPosition == 11 {
		mapOfPokemonMovementsFromDatabase["Level two movements"] = string(columnValue)
	} else if columnPosition == 12 {
		mapOfPokemonMovementsFromDatabase["Level three movements"] = string(columnValue)
	} else if columnPosition == 13 {
		mapOfPokemonMovementsFromDatabase["Level four movements"] = string(columnValue)
	} else if columnPosition == 14 {
		mapOfPokemonMovementsFromDatabase["Level five movements"] = string(columnValue)
	}

	return nil

}
