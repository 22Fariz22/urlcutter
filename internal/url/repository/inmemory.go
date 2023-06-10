package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/22Fariz22/urlcutter/pkg/grpcerrors"
)

type memoryStorage struct {
	storage map[string]string
	mutex   sync.RWMutex
}

// NewMemory создание структуры для инмемори типа
func NewMemory() *memoryStorage {
	return &memoryStorage{
		storage: map[string]string{},
	}
}

// Save url to in-memory
func (m *memoryStorage) Save(ctx context.Context, long, short string) (string, error) {

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if _, ok := m.storage[long]; ok {
		return m.storage[long], grpcerrors.ErrURLExists
	}

	m.storage[long] = short

	return short, nil
}

// Get url from in-memory
func (m *memoryStorage) Get(ctx context.Context, short string) (string, error) {
	fmt.Println("here in-memory repo Get()")
	fmt.Println("storage:", m.storage)

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for k, v := range m.storage {
		if v == short {
			return k, nil
		}
		fmt.Println(k, v)
	}

	return "", grpcerrors.ErrDoesNotExist
}
