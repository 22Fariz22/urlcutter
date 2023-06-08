package repository

import (
	"context"
	"fmt"
	"sync"
)

type MemoryStorage struct {
	storage map[string]string
	mutex   sync.RWMutex
}

// New создание структуры для инмемори типа
func NewMemory() *MemoryStorage {
	return &MemoryStorage{
		storage: map[string]string{},
	}
}

//func New() MemoryStorage {
//	return &memoryStorage{
//		storage: map[string]entity.URL{},
//	}
//}
//

func (m *MemoryStorage) Save(ctx context.Context) {
	fmt.Println("here in-memory repo Save()")

}

func (m *MemoryStorage) Get(ctx context.Context) {
	fmt.Println("here in-memory repo Get()")
}
