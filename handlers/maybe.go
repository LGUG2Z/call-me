package handlers

import (
	"fmt"

	"github.com/LGUG2Z/call-me/models"
	"github.com/LGUG2Z/call-me/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	bolt "go.etcd.io/bbolt"
)

const BoltPath = "call-me.db"
var ErrKeyIsLocked = fmt.Errorf("key is locked")

func GetMaybe(params operations.GetMaybeParams, principal *models.Principal) middleware.Responder {
	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewPostMaybeInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	if err := db.Batch(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("environments"))
		if err != nil {
			return err
		}

		if bucket.Get([]byte(params.Environment)) != nil {
			return ErrKeyIsLocked
		}

		return nil
	}); err != nil {
		return operations.NewGetMaybeForbidden()
	}

	return operations.NewGetMaybeOK()
}

func DeleteMaybe(params operations.DeleteMaybeParams, principal *models.Principal) middleware.Responder {
	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewDeleteMaybeInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	if err := db.Batch(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("environments"))
		if err != nil {
			return err
		}

		return bucket.Delete([]byte(params.Environment))
	}); err != nil {
		return operations.NewPostMaybeInternalServerError().WithPayload(err.Error())
	}

	return operations.NewDeleteMaybeNoContent()
}

func PostMaybe(params operations.PostMaybeParams, principal *models.Principal) middleware.Responder {
	db, err := bolt.Open(BoltPath, 0666, nil)
	if err != nil {
		return operations.NewPostMaybeInternalServerError().WithPayload(err.Error())
	}
	defer db.Close()

	if err := db.Batch(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("environments"))
		if err != nil {
			return err
		}

		if bucket.Get([]byte(params.Environment)) != nil {
			return ErrKeyIsLocked
		}

		return bucket.Put([]byte(params.Environment), []byte("locked"))
	}); err != nil {
		switch err {
		case ErrKeyIsLocked:
			return operations.NewPostMaybeForbidden()
		default:
			return operations.NewPostMaybeInternalServerError().WithPayload(err.Error())
		}
	}

	return operations.NewPostMaybeCreated()
}
