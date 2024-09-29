package service

import "github.com/LeoTwins/go-clean-architecture/internal/domain/repository"

// ++++++++++++++++++++++++++++++++
// [仮にこれがなかった場合]
//
// ユースケースで、重複チェックを実装することになるが、２つ３つのユースケースで同じ重複チェックを実装することにつながる
// これを防ぐために有効なアプローチ
type IFavoriteService interface {
	// そのユーザーのエクササイズIDが重複していないかチェックする
	IsExerciseUnique(userID uint, exerciseID uint) (bool, error)
}

type favoriteService struct {
	repo repository.IFavoriteRepository
}

// IsExerciseUnique implements IFavoriteService.
func (f *favoriteService) IsExerciseUnique(userID uint, exerciseID uint) (bool, error) {
	panic("unimplemented")
}

func NewFavoriteService(repo repository.IFavoriteRepository) IFavoriteService {
	return &favoriteService{
		repo: repo,
	}
}
