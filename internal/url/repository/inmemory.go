package repository

import (
	"context"
	"fmt"
	"github.com/22Fariz22/urlcutter/pkg/grpcerrors"
	"sync"
)

type MemoryStorage struct {
	storage map[string]string
	mutex   sync.RWMutex
}

// NewMemory создание структуры для инмемори типа
func NewMemory() *MemoryStorage {
	return &MemoryStorage{
		storage: map[string]string{},
	}
}

func (m *MemoryStorage) Save(ctx context.Context, long, short string) (string, error) {
	fmt.Println("here in-memory repo Save()")

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if val, ok := m.storage[long]; ok {
		fmt.Println("already exist in mermory:", val)
		return m.storage[long], grpcerrors.ErrURLExists
	}

	m.storage[long] = short

	return short, nil
}

func (m *MemoryStorage) Get(ctx context.Context, short string) (string, error) {
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
