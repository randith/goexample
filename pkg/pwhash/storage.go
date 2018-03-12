package pwhash

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("not found")
)

type Store interface {
	Get(key string) (string, error)
	// key is generated and returned
	Set(val string) (string, error)
}

type inMemoryStore struct {
	data map[string]string
	lck  sync.RWMutex
}

func NewInMemoryStore() Store {
	return &inMemoryStore{data: make(map[string]string)}
}

func (store *inMemoryStore) Get(key string) (string, error) {
	store.lck.RLock()
	defer store.lck.RUnlock()
	v, ok := store.data[key]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

func (store *inMemoryStore) Set(val string) (string, error) {
	store.lck.Lock()
	defer store.lck.Unlock()
	key, err := genId()
	if err != nil {
		return "", err
	}
	store.data[key] = val
	return key, nil
}

/**
 * TODO replace this simple uuid generator with something more robust such has https://github.com/satori/go.uuid
 */
func genId() (string, error) {
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		return "", err
	}
	u[8] = (u[8] | 0x80) & 0xBF
	u[6] = (u[6] | 0x40) & 0x4F
	return hex.EncodeToString(u), nil
}
