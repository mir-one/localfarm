package service

import (
	"errors"

	"github.com/mir-one/localfarm/src/user/domain"

	"github.com/mir-one/localfarm/src/user/query"
	"github.com/mir-one/localfarm/src/user/storage"
)

type UserServiceImpl struct {
	UserReadQuery query.UserReadQuery
}

func (s UserServiceImpl) FindUserByUsername(username string) (domain.UserServiceResult, error) {
	result := <-s.UserReadQuery.FindByUsername(username)

	if result.Error != nil {
		return domain.UserServiceResult{}, result.Error
	}

	user, ok := result.Result.(storage.UserRead)
	if !ok {
		return domain.UserServiceResult{}, errors.New("error type assertion")
	}

	return domain.UserServiceResult{
		UID:      user.UID,
		Username: user.Username,
	}, nil
}
