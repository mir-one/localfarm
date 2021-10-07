package inmemory

import (
	"github.com/mir-one/localfarm/src/assets/repository"
	"github.com/mir-one/localfarm/src/assets/storage"
)

type ReservoirReadRepositoryInMemory struct {
	Storage *storage.ReservoirReadStorage
}

func NewReservoirReadRepositoryInMemory(s *storage.ReservoirReadStorage) repository.ReservoirReadRepository {
	return &ReservoirReadRepositoryInMemory{Storage: s}
}

func (f *ReservoirReadRepositoryInMemory) Save(reservoirRead *storage.ReservoirRead) <-chan error {
	result := make(chan error)

	go func() {
		f.Storage.Lock.Lock()
		defer f.Storage.Lock.Unlock()

		f.Storage.ReservoirReadMap[reservoirRead.UID] = *reservoirRead

		result <- nil

		close(result)
	}()

	return result
}
