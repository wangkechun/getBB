package store

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	assert := assert.New(t)

	ak := os.Getenv("ACCESS_KEY")
	sk := os.Getenv("SECRET_KEY")
	bucket := os.Getenv("BUCKET")
	store := Init(ak, sk, bucket)

	assert.NotNil(store)

	url := "https://www.v2ex.com/t/457637"

	err := store.Save(url)
	assert.Nil(err)

	exists := store.IfExists(url)
	assert.True(exists)

	exists = store.IfExists(url + "abc")
	assert.False(exists)
}
