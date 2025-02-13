// Copyright 2019 Kaleido

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kvstore

import (
	"github.com/hyperledger/firefly-ethconnect/internal/errors"
	log "github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// ErrorNotFound signal error for not found
var ErrorNotFound = leveldb.ErrNotFound

// KVIterator interface for key value iterators
type KVIterator interface {
	Key() string
	Value() []byte
	Next() bool
	Prev() bool
	Seek(string) bool
	Last() bool
	Release()
}

// KVStore interface for key value stores
type KVStore interface {
	Put(key string, val []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	NewIterator() KVIterator
	NewIteratorWithRange(keyRange interface{}) KVIterator
	Close()
}

type levelDBKeyValueStore struct {
	path string
	db   *leveldb.DB
}

func (k *levelDBKeyValueStore) warnIfErr(op, key string, err error) {
	if err != nil && err != leveldb.ErrNotFound {
		log.Warnf("LDB %s %s '%s' failed: %s", k.path, op, key, err)
	}
}

func (k *levelDBKeyValueStore) Put(key string, val []byte) error {
	err := k.db.Put([]byte(key), val, nil)
	k.warnIfErr("Put", key, err)
	return err
}

func (k *levelDBKeyValueStore) Get(key string) ([]byte, error) {
	b, err := k.db.Get([]byte(key), nil)
	k.warnIfErr("Get", key, err)
	return b, err
}

func (k *levelDBKeyValueStore) Delete(key string) error {
	err := k.db.Delete([]byte(key), nil)
	k.warnIfErr("Delete", key, err)
	return err
}

func (k *levelDBKeyValueStore) NewIterator() KVIterator {
	return &levelDBKeyIterator{
		i: k.db.NewIterator(nil, nil),
	}
}

func (k *levelDBKeyValueStore) NewIteratorWithRange(rng interface{}) KVIterator {
	keyRange := rng.(*util.Range)

	return &levelDBKeyIterator{
		i: k.db.NewIterator(keyRange, nil),
	}
}

type levelDBKeyIterator struct {
	i iterator.Iterator
}

func (k *levelDBKeyIterator) Key() string {
	return string(k.i.Key())
}

func (k *levelDBKeyIterator) Value() []byte {
	return k.i.Value()
}

func (k *levelDBKeyIterator) Last() bool {
	return k.i.Last()
}

func (k *levelDBKeyIterator) Seek(key string) bool {
	return k.i.Seek([]byte(key))
}

func (k *levelDBKeyIterator) Next() bool {
	return k.i.Next()
}

func (k *levelDBKeyIterator) Prev() bool {
	return k.i.Prev()
}

func (k *levelDBKeyIterator) Release() {
	k.i.Next()
}

func (k *levelDBKeyValueStore) Close() {
	k.db.Close()
}

// NewLDBKeyValueStore construct a new LevelDB instance of a KV store
func NewLDBKeyValueStore(ldbPath string) (kv KVStore, err error) {
	store := &levelDBKeyValueStore{
		path: ldbPath,
	}
	if store.db, err = leveldb.OpenFile(ldbPath, nil); err != nil {
		return nil, errors.Errorf(errors.KVStoreDBLoad, ldbPath, err)
	}
	kv = store
	return
}
