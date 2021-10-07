package server

import (
	"github.com/labstack/gommon/log"
	"github.com/mir-one/localfarm/src/user/domain"
	"github.com/mir-one/localfarm/src/user/storage"
)

func (s *AuthServer) SaveToUserReadModel(event interface{}) error {
	userRead := &storage.UserRead{}

	switch e := event.(type) {
	case domain.UserCreated:
		userRead.UID = e.UID
		userRead.Username = e.Username
		userRead.Password = e.Password
		userRead.CreatedDate = e.CreatedDate
		userRead.LastUpdated = e.LastUpdated

	}

	err := <-s.UserReadRepo.Save(userRead)
	if err != nil {
		log.Error(err)
	}

	return nil
}
