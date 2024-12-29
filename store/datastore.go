package store

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type DataStore struct {
	Firestore *firestore.Client
	CloseFunc func()
}

func (ds *DataStore) StoreComic(ctx context.Context, collectionName string, data interface{}) error {
	doc, wr, err := ds.Firestore.Collection(collectionName).Add(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to writing data: %v", err)
	}

	fmt.Println("WriteResult: ", doc, wr)
	return nil
}

func InitStore(ctx context.Context, projectID string, accountCredentials string) (*DataStore, error) {
	opt := option.WithCredentialsFile(accountCredentials)
	client, err := firestore.NewClient(ctx, "my-personal-labs-395004", opt)
	if err != nil {
		return nil, err
	}
	closeFunc := func() {
		client.Close()
	}

	return &DataStore{Firestore: client, CloseFunc: closeFunc}, nil
}
