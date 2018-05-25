package store

import (
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// Store spider db
type Store struct {
	accessKey string
	secretKey string
	bucket    string
	*storage.BucketManager
}

// Init a store
func Init(ak, sk, bucket string) *Store {
	store := &Store{
		accessKey: ak,
		secretKey: sk,
		bucket:    bucket,
	}

	mac := qbox.NewMac(store.accessKey, store.secretKey)
	cfg := storage.Config{
		UseHTTPS: false,
	}

	store.BucketManager = storage.NewBucketManager(mac, &cfg)
	return store
}

// Save page
func (store *Store) Save(url string) error {
	_, err := store.Fetch(url, store.bucket, url)
	return err
}

// IfExists check exists
func (store *Store) IfExists(url string) {
	mac := qbox.NewMac(store.accessKey, store.secretKey)
	cfg := storage.Config{
		UseHTTPS: false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	fileInfo, sErr := bucketManager.Stat(store.bucket, url)
	if sErr != nil {
		fmt.Println(sErr)
		return
	}
	fmt.Println(fileInfo.String())
	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
}
