package repository

import (
	"context"
	"sync"

	"github.com/22Fariz22/urlcutter/pkg/logger"

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
func (m *memoryStorage) Save(ctx context.Context, l logger.Interface, long, short string) (string, error) {

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if _, ok := m.storage[long]; ok {
		return m.storage[long], grpcerrors.ErrURLExists
	}

	m.storage[long] = short

	return short, nil
}

// Get url from in-memory
func (m *memoryStorage) Get(ctx context.Context, l logger.Interface, short string) (string, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for k, v := range m.storage {
		if v == short {
			return k, nil
		}
	}

	return "", grpcerrors.ErrDoesNotExist
}
