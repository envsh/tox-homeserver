package couchcache

import (
	"sync"
)

const (
	maxTTLInSec   = 60 * 60 * 24 * 30
	maxSizeInByte = 20 * 1024 * 1024
	maxKeyLength  = 250
)

type couchbaseDatastore struct {
	sync.Map
}

func newDatastore() (ds *couchbaseDatastore, err error) {
	return &couchbaseDatastore{}, nil
}

func (ds *couchbaseDatastore) get(k string) []byte {
	if v, ok := ds.Load(k); ok {
		return v.([]byte)
	}
	return nil
}

func (ds *couchbaseDatastore) set(k string, v []byte, ttl int) error {
	ds.Store(k, v)
	rcp.changed(k)
	return nil
}

func (ds *couchbaseDatastore) delete(k string) error {
	ds.Delete(k)
	rcp.changed(k)
	return nil
}

func (ds *couchbaseDatastore) append(k string, v []byte) error {
	if vx, ok := ds.Load(k); ok {
		ds.Store(k, append(vx.([]byte), v...))
	} else {
		ds.Store(k, v)
	}
	rcp.changed(k)
	return nil
}

func (ds *couchbaseDatastore) validKey(key string) error {
	return nil
}

func (ds *couchbaseDatastore) validValue(v []byte) error {
	return nil
}
