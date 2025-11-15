package modules

import (
	"github.com/google/uuid"
	"github.com/yadav-shubh/base-middleware/graph/model"
	"time"
)

var (
	id  = uuid.New().String()
	now = int(time.Now().UnixMilli())
)

func CreateAppUser(input model.AppUserInput) (*model.AppUser, error) {
	return &model.AppUser{
		ID:        &id,
		Name:      input.Name,
		Username:  input.Username,
		Mobile:    input.Mobile,
		Role:      input.Role,
		IsActive:  true,
		IsDeleted: false,
		CreatedAt: &now,
		UpdatedAt: &now,
	}, nil
}

func GetAppUser(id string) (*model.AppUser, error) {
	return &model.AppUser{
		ID:        &id,
		Name:      "Shubh",
		Username:  "shubh",
		Mobile:    "1234567890",
		Role:      "admin",
		IsActive:  true,
		IsDeleted: false,
		CreatedAt: &now,
		UpdatedAt: &now,
	}, nil
}
