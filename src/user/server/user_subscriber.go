package server

import (
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/mir-one/localfarm/src/user/domain"
	"github.com/mir-one/localfarm/src/user/storage"
)

func (s *UserServer) SaveToUserReadModel(event interface{}) error {
	userRead := &storage.UserRead{}

	switch e := event.(type) {
	case domain.PasswordChanged:
		queryResult := <-s.UserReadQuery.FindByID(e.UID)
		if queryResult.Error != nil {
			log.Error(queryResult.Error)
		}

		u, ok := queryResult.Result.(storage.UserRead)
		if !ok {
			log.Error(errors.New("internal server error. error type assertion"))
		}

		userRead = &u

		userRead.Password = e.NewPassword
		userRead.LastUpdated = e.DateChanged

	}

	err := <-s.UserReadRepo.Save(userRead)
	if err != nil {
		log.Error(err)
	}

	return nil
}
