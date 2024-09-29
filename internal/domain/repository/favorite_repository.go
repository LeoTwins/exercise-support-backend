package repository

import "github.com/LeoTwins/go-clean-architecture/internal/domain/model"

type IFavoriteRepository interface {
	FindByUserID(userID uint) ([]*model.Favorite, error)
	Save(userID uint, exerciseID uint) error
	Delete(userID uint, exerciseID uint) error
}
