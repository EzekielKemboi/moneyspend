package main

import (
	"log"

	"github.com/boltdb/bolt"
)

const (
	bucketName = "zest"
)

var boltdb *bolt.DB

func MustInitBolt(freshStart bool) {
	var err error

	boltdb, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatalf("open boltdb failed,err: %v", err)
	}

	if freshStart {
		boltDeleteBucket(bucketName)
	}

	boltCreateBucket(bucketName)
}

func boltDeleteBucket(bucketName string) {
	boltdb.Update(func(t *bolt.Tx) error {
		err := t.DeleteBucket([]byte(bucketName))
		if err != nil {
			log.Fatalf("delete bucket failed,err: %v", err)
		}
		return nil
	})
}

func boltCreateBucket(bucketName string) {
	boltdb.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			log.Fatalf("create bucket failed,err: %v", err)
		}
		return nil
	})
}

func BoltGet(key string) string {
	var val string
	boltdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		val = string(b.Get([]byte(key)))
		return nil
	})
	return val
}

func BoltSet(key string, val string) {
	boltdb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Put([]byte(key), []byte(val))
		if err != nil {
			log.Printf("bolt put err: %v,key: %v,val: %v", err, key, val)
		}
		return err
	})
}
