package datasaver

import (
	"fmt"
	"os"

	"github.com/AlejandroWaiz/PIA-PokemonSaver-CloudFunction/model"
	"github.com/xuri/excelize/v2"
)

type SaverImplementation struct {
	pokemonsExcelFile *excelize.File
}

type SaverInterface interface {
	SaveAllPokemons() ([]model.Pokemon, []error)
}

func GetExcelAdapterImplementation() (SaverInterface, error) {

	pokemonFile, err := excelize.OpenFile(os.Getenv("pokemon_excel_filename"))

	if err != nil {
		return &SaverImplementation{}, fmt.Errorf("Error getting %v file: %v", os.Getenv("pokemon_excel_filename"), err)
	}

	saverImplementation := SaverImplementation{pokemonsExcelFile: pokemonFile}

	return &saverImplementation, nil

}
