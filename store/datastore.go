package store

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type DataStore struct {
	firestore *firestore.Client
}

func (ds *DataStore) StoreComic(ctx context.Context, collectionName string, data interface{}) error {
	doc, wr, err := ds.firestore.Collection(collectionName).Add(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to writing data: %v", err)
	}

	fmt.Println("WriteResult: ", doc, wr)
	return nil
}

func InitStore(ctx context.Context, projectID string, accountCredentials string) (*DataStore, func(), error) {

	if projectID == "" || accountCredentials == "" {
		return nil, func() {}, fmt.Errorf("projectID and accountCredentials must be provided")
	}

	opt := option.WithCredentialsFile(accountCredentials)
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		return nil, func() {}, err
	}
	closeFunc := func() {
		client.Close()
	}

	return &DataStore{firestore: client}, closeFunc, nil
}
