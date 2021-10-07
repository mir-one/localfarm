package inmemory

import (
	"github.com/gofrs/uuid"
	"github.com/mir-one/localfarm/src/assets/repository"
	"github.com/mir-one/localfarm/src/assets/storage"
)

type ReservoirEventRepositoryInMemory struct {
	Storage *storage.ReservoirEventStorage
}

func NewReservoirEventRepositoryInMemory(s *storage.ReservoirEventStorage) repository.ReservoirEventRepository {
	return &ReservoirEventRepositoryInMemory{Storage: s}
}

func (f *ReservoirEventRepositoryInMemory) Save(uid uuid.UUID, latestVersion int, events []interface{}) <-chan error {
	result := make(chan error)

	go func() {
		f.Storage.Lock.Lock()
		defer f.Storage.Lock.Unlock()

		for _, v := range events {
			latestVersion++
			f.Storage.ReservoirEvents = append(f.Storage.ReservoirEvents, storage.ReservoirEvent{
				ReservoirUID: uid,
				Version:      latestVersion,
				Event:        v,
			})
		}

		result <- nil

		close(result)
	}()

	return result
}
