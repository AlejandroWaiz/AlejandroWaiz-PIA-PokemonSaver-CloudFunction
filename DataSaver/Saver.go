package datasaver

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/AlejandroWaiz/PIA-PokemonSaver-CloudFunction/model"
	"google.golang.org/api/option"
)

type SaverImplementation struct {
	client *firestore.Client
}

type SaverInterface interface {
	SaveManyPokemon(Pokemons []model.Pokemon) (map[string]string, error)
}

func GetFirestoreAdapterImplementation() (SaverInterface, error) {

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("Firestore_project_id"), option.WithCredentialsFile("service-account.json"))

	if err != nil {
		log.Printf("Err converting string to int from excel: %v", err)
		return nil, fmt.Errorf("Err creating firestore client: %v", err)
	}

	return &SaverImplementation{client: client}, err

}
