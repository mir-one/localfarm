package service

import (
	"github.com/gofrs/uuid"
	"github.com/mir-one/localfarm/src/assets/domain"
	"github.com/mir-one/localfarm/src/assets/query"
	"github.com/mir-one/localfarm/src/assets/storage"
)

type ReservoirServiceInMemory struct {
	FarmReadQuery query.FarmReadQuery
}

func (s ReservoirServiceInMemory) FindFarmByID(uid uuid.UUID) (domain.ReservoirFarmServiceResult, error) {
	result := <-s.FarmReadQuery.FindByID(uid)

	if result.Error != nil {
		return domain.ReservoirFarmServiceResult{}, result.Error
	}

	farm, ok := result.Result.(storage.FarmRead)

	if !ok {
		return domain.ReservoirFarmServiceResult{}, domain.ReservoirError{Code: domain.ReservoirErrorFarmNotFound}
	}

	if farm == (storage.FarmRead{}) {
		return domain.ReservoirFarmServiceResult{}, domain.ReservoirError{Code: domain.ReservoirErrorFarmNotFound}
	}

	return domain.ReservoirFarmServiceResult{
		UID:  farm.UID,
		Name: farm.Name,
	}, nil
}
