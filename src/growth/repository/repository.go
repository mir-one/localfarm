package repository

import (
	"github.com/gofrs/uuid"
	"github.com/mir-one/localfarm/src/growth/domain"
	"github.com/mir-one/localfarm/src/growth/storage"
)

// RepositoryResult is a struct to wrap repository result
// so its easy to use it in channel
type RepositoryResult struct {
	Result interface{}
	Error  error
}

type CropEventRepository interface {
	Save(uid uuid.UUID, latestVersion int, events []interface{}) <-chan error
}

type CropReadRepository interface {
	Save(cropRead *storage.CropRead) <-chan error
}

func NewCropBatchFromHistory(events []storage.CropEvent) *domain.Crop {
	state := &domain.Crop{}
	for _, v := range events {
		state.Transition(v.Event)
		state.Version++
	}
	return state
}

type CropActivityRepository interface {
	Save(cropActivity *storage.CropActivity, isUpdate bool) <-chan error
}
