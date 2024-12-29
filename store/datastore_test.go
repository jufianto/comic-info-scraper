package store

import (
	"context"
	"testing"

	cl "github.com/jufianto/comic-info-scraper/services"
	"github.com/stretchr/testify/assert"
)

func TestInitStore(t *testing.T) {
	ctx := context.Background()
	accountCredentials := "../cmd/config/firestore-gcp-access.json"

	clientDs, close, err := InitStore(ctx, "my-personal-labs-395004", accountCredentials)
	defer close()

	assert.NoError(t, err)
	assert.NotNil(t, clientDs)

	collectionName := "comic-scraper"
	comicScraperList := []cl.InfoComic{
		{Title: "One Piece", LastChapter: "1010"},
		{Title: "Naruto", LastChapter: "700"},
	}

	data, err := ConvertToJSON(comicScraperList)
	assert.NoError(t, err)

	err = clientDs.StoreComic(ctx, collectionName, data)
	if err != nil {
		t.Fatalf("failed to store data %v", err)
	}

	t.Log("success store data")
}
