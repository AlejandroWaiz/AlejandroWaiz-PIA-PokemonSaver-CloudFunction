package datasaver

import (
	"context"
	"errors"
	"os"

	"github.com/AlejandroWaiz/PIA-PokemonSaver-CloudFunction/model"
)

func (s *SaverImplementation) SaveManyPokemon(Pokemons []model.Pokemon) (map[string]string, error) {

	ctx := context.Background()

	var notSavedPokemons = make(map[string]string)

	for _, pokemon := range Pokemons {

		_, _, err := s.client.Collection(os.Getenv("collectionName_fromFirestore_OfPokemons")).Add(ctx, pokemon)

		if err != nil {
			notSavedPokemons[pokemon.Name.(string)] = err.Error()
		}

	}

	if len(notSavedPokemons) > 0 {

		return notSavedPokemons, errors.New("Error saving some pokemons to database")

	}

	return nil, nil

}
