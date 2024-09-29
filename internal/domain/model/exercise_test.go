package model_test

import (
	"testing"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewExercise_Success(t *testing.T) {
	arr := []uint{1, 2, 3}
	_, err := model.NewExercise(
		nil,
		model.Title("v"),
		"image",
		model.Description("v"),
		arr,
	)

	assert.NoError(t, err)
}

func TestNewExercise_Failure_MenuIDEmpty(t *testing.T) {
	arr := []uint{}
	_, err := model.NewExercise(
		nil,
		model.Title("v"),
		"image",
		model.Description("v"),
		arr,
	)

	assert.Error(t, err)
	assert.EqualError(t, err, "エクササイズに紐づくメニューは最低１つ選択してください")
}

func TestNewExercise_Failure_MenuID５つ以上指定している(t *testing.T) {
	arr := []uint{1, 2, 3, 4, 5, 6}
	_, err := model.NewExercise(
		nil,
		model.Title("v"),
		"image",
		model.Description("v"),
		arr,
	)

	assert.Error(t, err)
	assert.EqualError(t, err, "エクササイズに紐づくメニューは5つ以内にしてください")
}
