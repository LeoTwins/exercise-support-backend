package repository

import "github.com/LeoTwins/go-clean-architecture/internal/domain/model"

type IExerciseRepository interface {
	Find() ([]*model.Exercise, error)
	// TODO : ExerciseIDという値オブジェクトがいいね
	FindByID(ID uint) (*model.Exercise, error)
	Save(item *model.Exercise) error
	Update(item *model.Exercise) error
	// TODO : ExerciseIDという値オブジェクトがいいね
	Delete(ID uint) error
}
