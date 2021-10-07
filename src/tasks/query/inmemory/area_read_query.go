package inmemory

import (
	"github.com/gofrs/uuid"
	"github.com/mir-one/localfarm/src/assets/storage"
	"github.com/mir-one/localfarm/src/tasks/query"
)

type AreaQueryInMemory struct {
	Storage *storage.AreaReadStorage
}

func NewAreaQueryInMemory(s *storage.AreaReadStorage) query.AreaQuery {
	return AreaQueryInMemory{Storage: s}
}

func (s AreaQueryInMemory) FindByID(uid uuid.UUID) <-chan query.QueryResult {
	result := make(chan query.QueryResult)

	go func() {
		s.Storage.Lock.RLock()
		defer s.Storage.Lock.RUnlock()

		area := query.TaskAreaQueryResult{}
		for _, val := range s.Storage.AreaReadMap {
			if val.UID == uid {
				area.UID = uid
				area.Name = val.Name
			}
		}

		result <- query.QueryResult{Result: area}

		close(result)
	}()

	return result
}
