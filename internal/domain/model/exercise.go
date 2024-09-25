package model

import "errors"

type Exercise struct {
	ID          *uint
	Title       Title
	Image       string
	Description Description
	Menus       []uint
}

func NewExercise(id *uint, title Title, image string, description Description, menus []uint) (*Exercise, error) {
	if image == "" {
		return nil, errors.New("エクササイズイメージ画像は必須です")
	}

	if len(menus) == 0 {
		return nil, errors.New("エクササイズに紐づくメニューは最低１つ選択してください")
	}

	if len(menus) > 5 {
		return nil, errors.New("エクササイズに紐づくメニューは5つ以内にしてください")
	}

	return &Exercise{
		ID:          id,
		Title:       title,
		Image:       image,
		Description: description,
		Menus:       menus,
	}, nil
}

func ReCreateExercise(id *uint, title Title, image string, description Description, menus []uint) *Exercise {
	return &Exercise{
		ID:          id,
		Title:       title,
		Image:       image,
		Description: description,
		Menus:       menus,
	}
}
