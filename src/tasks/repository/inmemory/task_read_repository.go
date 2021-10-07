package inmemory

import (
	"github.com/mir-one/localfarm/src/tasks/repository"
	"github.com/mir-one/localfarm/src/tasks/storage"
)

type TaskReadRepositoryInMemory struct {
	Storage *storage.TaskReadStorage
}

func NewTaskReadRepositoryInMemory(s *storage.TaskReadStorage) repository.TaskReadRepository {
	return &TaskReadRepositoryInMemory{Storage: s}
}

// Save is to save
func (f *TaskReadRepositoryInMemory) Save(taskRead *storage.TaskRead) <-chan error {
	result := make(chan error)

	go func() {
		f.Storage.Lock.Lock()
		defer f.Storage.Lock.Unlock()

		f.Storage.TaskReadMap[taskRead.UID] = *taskRead

		result <- nil

		close(result)
	}()

	return result
}
