package model

import "errors"

type Favorite struct {
	UserID     uint
	ExerciseID uint
}

func NewFavorite(userID, exerciseID uint) (Favorite, error) {
	// TODO : IDという値オブジェクトかな？ UserID、ExerciseIDの別で作成したほうが良さそう
	if userID == 0 {
		return Favorite{}, errors.New("ユーザーIDは必須です")
	}

	if exerciseID == 0 {
		return Favorite{}, errors.New("エクササイズIDは必須です")
	}

	return Favorite{
		UserID:     userID,
		ExerciseID: exerciseID,
	}, nil
}

func ReCreateFavorite(userID, exerciseID uint) Favorite {
	return Favorite{
		UserID:     userID,
		ExerciseID: exerciseID,
	}
}
