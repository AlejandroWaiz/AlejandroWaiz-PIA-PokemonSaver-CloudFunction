package datasaver

import (
	"fmt"
	"strconv"

	"github.com/AlejandroWaiz/PIA-PokemonSaver-CloudFunction/model"
)

var mapOfPokemonMovementsFromDatabase map[string]string

func assignValueFromDatabaseToPokemonMold(pokemon model.Pokemon, columnPosition int, columnValue string) error {

	if columnPosition == 0 {
		pokemonMold.ID = columnValue
	} else if columnPosition == 1 {
		pokemonMold.Name = columnValue
	} else if columnPosition == 2 {
		pokemonMold.Pasive = columnValue
	} else if columnPosition == 3 {
		columnValueAsInt, err := strconv.Atoi(columnValue)
		if err != nil {
			return fmt.Errorf("Err converting string to int from excel: %v", err)
		}
		pokemonMold.Stage = columnValueAsInt
	} else if columnPosition == 4 {
		columnValueAsInt, err := strconv.Atoi(columnValue)
		if err != nil {
			return fmt.Errorf("Err converting string to int from excel: %v", err)
		}
		pokemonMold.HP = columnValueAsInt
	} else if columnPosition == 5 {
		pokemonMold.HpLevelBonus = columnValue
	} else if columnPosition == 6 {
		columnValueAsInt, err := strconv.Atoi(columnValue)
		if err != nil {
			return fmt.Errorf("Err converting string to int from excel: %v", err)
		}
		pokemonMold.Speed = columnValueAsInt
	} else if columnPosition == 7 {
		pokemonMold.SpeedLevelBonus = columnValue
	} else if columnPosition == 8 {
		mapOfPokemonMovementsFromDatabase["Base movements"] = columnValue
	} else if columnPosition == 9 {
		mapOfPokemonMovementsFromDatabase["Level one movements"] = columnValue
	} else if columnPosition == 10 {
		mapOfPokemonMovementsFromDatabase["Level two movements"] = columnValue
	} else if columnPosition == 11 {
		mapOfPokemonMovementsFromDatabase["Level three movements"] = columnValue
	} else if columnPosition == 12 {
		mapOfPokemonMovementsFromDatabase["Level four movements"] = columnValue
	} else if columnPosition == 13 {
		mapOfPokemonMovementsFromDatabase["Level five movements"] = columnValue
	}

	return nil

}
