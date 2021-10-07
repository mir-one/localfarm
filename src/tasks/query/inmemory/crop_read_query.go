package inmemory

import (
	"github.com/gofrs/uuid"
	"github.com/mir-one/localfarm/src/growth/storage"
	"github.com/mir-one/localfarm/src/tasks/query"
)

type CropQueryInMemory struct {
	Storage *storage.CropReadStorage
}

func NewCropQueryInMemory(s *storage.CropReadStorage) query.CropQuery {
	return CropQueryInMemory{Storage: s}
}

func (s CropQueryInMemory) FindCropByID(uid uuid.UUID) <-chan query.QueryResult {
	result := make(chan query.QueryResult)

	go func() {
		s.Storage.Lock.RLock()
		defer s.Storage.Lock.RUnlock()

		crop := query.TaskCropQueryResult{}

		for _, val := range s.Storage.CropReadMap {
			if val.UID == uid {
				crop.UID = uid
				crop.BatchID = val.BatchID
			}
		}
		result <- query.QueryResult{Result: crop}

		close(result)
	}()

	return result
}
