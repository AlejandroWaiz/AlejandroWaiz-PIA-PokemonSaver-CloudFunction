package main

import (
	"context"
	"log"

	datareader "github.com/AlejandroWaiz/PIA-PokemonSaver-CloudFunction/DataReader"
	datasaver "github.com/AlejandroWaiz/PIA-PokemonSaver-CloudFunction/DataSaver"
	"github.com/joho/godotenv"
)

/* type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
} */

func main() {

	godotenv.Load(".env")

	ctx := context.Background()
	fileName := "Pokemon.xlsx"
	bucketName := "bucket-for-save-pokemon"

	reader, err := datareader.GetExcelAdapterImplementation(ctx, fileName, bucketName)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	saver, err := datasaver.GetFirestoreAdapterImplementation()

	if err != nil {
		log.Println(err)
		panic(err)
	}

	pokemons, arrOfErr := reader.ReadAllPokemons()

	if err != nil {
		log.Println(arrOfErr)
	}

	notSavedPokemons, err := saver.SaveManyPokemon(pokemons)

	if err != nil {
		log.Println(err)
	}

	if notSavedPokemons != nil {
		log.Printf("not saved pokemons: %v", notSavedPokemons)
	}

}
