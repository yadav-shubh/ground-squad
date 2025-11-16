package db

import (
	"github.com/yadav-shubh/base-backend/modules/app_user/db/sqlc/generated"
	"github.com/yadav-shubh/base-backend/utils"
)

var AppUserRepository *generated.Queries

func init() {
	AppUserRepository = generated.New(utils.GetDB())
}
