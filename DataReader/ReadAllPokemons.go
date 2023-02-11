package datareader

import (
	"log"
	"os"
	"strings"

	"github.com/AlejandroWaiz/PIA-PokemonSaver-CloudFunction/model"
)

var arrayOfPokemons []model.Pokemon
var pokemonMold model.Pokemon

func (r *ReaderImplementation) ReadAllPokemons() ([]model.Pokemon, []error) {

	allPokemonExcelSheets := os.Getenv("pokemon_excel_sheetnames")

	arrayOfPokemonSheetsFromExcel := strings.Split(allPokemonExcelSheets, ",")

	var arrayOfErrors []error

	//Por cada hoja del excel de Pokemons se hará una iteración, recorriendo el documento en su totalidad.
	for _, thisSheet := range arrayOfPokemonSheetsFromExcel {

		allRows, err := r.pokemonsExcelFile.GetRows(thisSheet)

		if err != nil {
			arrayOfErrors = append(arrayOfErrors, err)
			break
		}

		for i := range allRows {

			//Se ignora la primera columna, donde va el nombre de las columnas eg: ID, Name, Type, etc...
			if i == 0 {
				continue
			}

			//Aqui se itera sobre las columnas, tomando el valor de cada una y asignandola a una struct molde, la cual se almacena
			//en un array de pokemones el cual será entregado al adaptador de salida para ser almacenado.
			for _, onlyColumns := range allRows {

				for columnPosition, columnValue := range onlyColumns {

					err := assignValueFromDatabaseToPokemonMold(pokemonMold, columnPosition, columnValue)

					if err != nil {

						log.Printf("Err assigning value %v from excel to pokemon: %v ", columnValue, err)

						arrayOfErrors = append(arrayOfErrors, err)

						break

					}

				}

			}

			arrayOfPokemons = append(arrayOfPokemons, pokemonMold)

			pokemonMold.ResetStruct()

		}

	}

	if len(arrayOfErrors) > 0 {
		log.Printf("Err reading pokemons: %v", arrayOfErrors)
		return nil, arrayOfErrors
	}

	return arrayOfPokemons, nil

}
