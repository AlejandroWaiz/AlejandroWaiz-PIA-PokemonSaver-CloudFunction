package datasaver

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"

	"google.golang.org/api/option"
)

type SaverImplementation struct {
	client *firestore.Client
}

type SaverInterface interface{}

func GetFirestoreAdapterImplementation() (SaverInterface, error) {

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("Firestore_project_id"), option.WithCredentialsFile(os.Getenv("Firestore_sa_path")))

	if err != nil {
		return nil, fmt.Errorf("Err creating firestore client: %v", err)
	}

	return &SaverImplementation{client: client}, err

}
