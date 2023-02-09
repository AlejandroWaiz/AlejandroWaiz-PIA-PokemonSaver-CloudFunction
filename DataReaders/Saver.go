package datareader

import (
	"fmt"
	"os"

	"github.com/AlejandroWaiz/PIA-PokemonSaver-CloudFunction/model"
	"github.com/xuri/excelize/v2"
)

type ReaderImplementation struct {
	pokemonsExcelFile *excelize.File
}

type ReaderInterface interface {
	ReadAllPokemons() ([]model.Pokemon, []error)
}

func GetExcelAdapterImplementation() (ReaderInterface, error) {

	pokemonFile, err := excelize.OpenFile(os.Getenv("pokemon_excel_filename"))

	if err != nil {
		return &ReaderImplementation{}, fmt.Errorf("Error getting %v file: %v", os.Getenv("pokemon_excel_filename"), err)
	}

	readerImplementation := ReaderImplementation{pokemonsExcelFile: pokemonFile}

	return &readerImplementation, nil

}
