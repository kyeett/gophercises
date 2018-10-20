package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	dbName     = "my.db"
	bucketName = "urls"
)

func main() {

	db, err := bolt.Open(dbName, 0777, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return errors.Wrap(err, "create bucket failed")
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	//Read all data
	log.Println("Read all data:")
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()

		cnt := 0
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
			cnt++
		}

		return nil
	})

	log.Println("Create new entry")
	uuid := []byte(uuid.New().String())
	val := []byte("10")

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Put(uuid, val)
		if err != nil {
			return err
		}
		return nil
	})

	//Write some data
}
