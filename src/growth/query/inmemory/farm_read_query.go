package inmemory

import (
	"github.com/Tanibox/tania-core/src/assets/storage"
	"github.com/Tanibox/tania-core/src/growth/query"
	"github.com/gofrs/uuid"
)

type FarmReadQueryInMemory struct {
	Storage *storage.FarmReadStorage
}

func NewFarmReadQueryInMemory(s *storage.FarmReadStorage) query.FarmReadQuery {
	return FarmReadQueryInMemory{Storage: s}
}

func (s FarmReadQueryInMemory) FindByID(uid uuid.UUID) <-chan query.Result {
	result := make(chan query.Result)

	go func() {
		s.Storage.Lock.RLock()
		defer s.Storage.Lock.RUnlock()

		farm := query.CropFarmQueryResult{}

		for _, val := range s.Storage.FarmReadMap {
			if val.UID == uid {
				farm.UID = uid
				farm.Name = val.Name
			}
		}

		result <- query.Result{Result: farm}

		close(result)
	}()

	return result
}
