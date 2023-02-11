package datareader

import (
	"fmt"
	"log"

	"github.com/AlejandroWaiz/PIA-PokemonSaver-CloudFunction/model"
	"github.com/xuri/excelize/v2"
)

type ReaderImplementation struct {
	pokemonsExcelFile *excelize.File
}

type ReaderInterface interface {
	ReadAllPokemons() ([]model.Pokemon, []error)
}

func GetExcelAdapterImplementation(fileName string) (ReaderInterface, error) {

	pokemonFile, err := excelize.OpenFile(fileName)

	if err != nil {
		log.Printf("Error getting %v file: %v", fileName, err)
		return &ReaderImplementation{}, fmt.Errorf("Error getting %v file: %v", fileName, err)
	}

	readerImplementation := ReaderImplementation{pokemonsExcelFile: pokemonFile}

	return &readerImplementation, nil

}
