package inmemory

import (
	"sort"

	"github.com/gofrs/uuid"
	"github.com/mir-one/localfarm/src/tasks/query"
	"github.com/mir-one/localfarm/src/tasks/storage"
)

type TaskEventQueryInMemory struct {
	Storage *storage.TaskEventStorage
}

func NewTaskEventQueryInMemory(s *storage.TaskEventStorage) query.TaskEventQuery {
	return &TaskEventQueryInMemory{Storage: s}
}

func (f *TaskEventQueryInMemory) FindAllByTaskID(uid uuid.UUID) <-chan query.QueryResult {
	result := make(chan query.QueryResult)

	go func() {
		f.Storage.Lock.RLock()
		defer f.Storage.Lock.RUnlock()

		events := []storage.TaskEvent{}
		for _, v := range f.Storage.TaskEvents {
			if v.TaskUID == uid {
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
