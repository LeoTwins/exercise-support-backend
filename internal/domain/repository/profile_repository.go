package repository

import (
	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
)

type IProfileRepository interface {
	Find() ([]*model.Profile, error)
	FindByUserID(userID uint) (*model.Profile, error)
}
