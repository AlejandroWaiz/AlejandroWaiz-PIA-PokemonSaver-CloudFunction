package datareader

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"github.com/AlejandroWaiz/PIA-PokemonSaver-CloudFunction/model"
	"github.com/xuri/excelize/v2"
)

type ReaderImplementation struct {
	pokemonsExcelFile *excelize.File
}

type ReaderInterface interface {
	ReadAllPokemons() ([]model.Pokemon, []error)
}

func GetExcelAdapterImplementation(ctx context.Context, fileName, bucketName string) (ReaderInterface, error) {

	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Printf("Error creating storage client: %v", err)
		return &ReaderImplementation{}, fmt.Errorf("Error creating storage client: %v", err)
	}

	fileReader, err := client.Bucket(bucketName).Object(fileName).NewReader(ctx)

	if err != nil {
		log.Printf("Error creating reader from file: %v", err)
		return &ReaderImplementation{}, fmt.Errorf("Error creating reader from file: %v", err)
	}

	pokemonFile, err := excelize.OpenReader(fileReader)

	if err != nil {
		log.Printf("Error getting %v file: %v", fileName, err)
		return &ReaderImplementation{}, fmt.Errorf("Error getting %v file: %v", fileName, err)
	}

	readerImplementation := ReaderImplementation{pokemonsExcelFile: pokemonFile}

	return &readerImplementation, nil

}
