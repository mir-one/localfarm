package inmemory

import (
	"github.com/gofrs/uuid"
	"github.com/mir-one/localfarm/src/tasks/repository"
	"github.com/mir-one/localfarm/src/tasks/storage"
)

type TaskEventRepositoryInMemory struct {
	Storage *storage.TaskEventStorage
}

func NewTaskEventRepositoryInMemory(s *storage.TaskEventStorage) repository.TaskEventRepository {
	return &TaskEventRepositoryInMemory{Storage: s}
}

// Save is to save
func (f *TaskEventRepositoryInMemory) Save(uid uuid.UUID, latestVersion int, events []interface{}) <-chan error {
	result := make(chan error)

	go func() {
		f.Storage.Lock.Lock()
		defer f.Storage.Lock.Unlock()

		for _, v := range events {
			latestVersion++
			f.Storage.TaskEvents = append(f.Storage.TaskEvents, storage.TaskEvent{
				TaskUID: uid,
				Version: latestVersion,
				Event:   v,
			})
		}

		result <- nil

		close(result)
	}()

	return result
}
