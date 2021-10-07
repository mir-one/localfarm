package server

import (
	"github.com/mir-one/localfarm/src/user/domain"
	"github.com/mir-one/localfarm/src/user/storage"
)

func MapToUserRead(user *domain.User) storage.UserRead {
	userRead := storage.UserRead{}
	userRead.UID = user.UID
	userRead.Username = user.Username
	userRead.CreatedDate = user.CreatedDate
	userRead.LastUpdated = user.LastUpdated

	return userRead
}
