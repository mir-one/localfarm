package inmemory

import (
	"sort"

	"github.com/gofrs/uuid"
	"github.com/mir-one/localfarm/src/assets/query"
	"github.com/mir-one/localfarm/src/assets/storage"
)

type ReservoirEventQueryInMemory struct {
	Storage *storage.ReservoirEventStorage
}

func NewReservoirEventQueryInMemory(s *storage.ReservoirEventStorage) query.ReservoirEventQuery {
	return &ReservoirEventQueryInMemory{Storage: s}
}

func (f *ReservoirEventQueryInMemory) FindAllByID(uid uuid.UUID) <-chan query.QueryResult {
	result := make(chan query.QueryResult)

	go func() {
		f.Storage.Lock.RLock()
		defer f.Storage.Lock.RUnlock()

		events := []storage.ReservoirEvent{}
		for _, v := range f.Storage.ReservoirEvents {
			if v.ReservoirUID == uid {
				events = append(events, v)
			}
		}

		sort.Slice(events, func(i, j int) bool {
			return events[i].Version < events[j].Version
		})

		result <- query.QueryResult{Result: events}
	}()

	return result
}
