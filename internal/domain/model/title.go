package model

import "errors"

type Title string

func NewTitle(value string) (*Title, error) {
	if len(value) > 30 {
		return nil, errors.New("タイトルは30文字以内です")
	}
	if value == "" {
		return nil, errors.New("タイトルは必須です")
	}

	title := Title(value)

	return &title, nil
}

func (t Title) String() string {
	return string(t)
}
