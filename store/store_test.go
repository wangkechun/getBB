package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	assert := assert.New(t)

	ak := ""
	sk := ""
	bucket := "spider"
	store := Init(ak, sk, bucket)

	assert.NotNil(store)

	err := store.Save("https://www.v2ex.com/t/457637")
	assert.Nil(err)
}
