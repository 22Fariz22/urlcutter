package repository

import "sync"

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

func (m *MemoryStorage) Save() {

}

func (m *MemoryStorage) Get() {

}
