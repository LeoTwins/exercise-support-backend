package model

import "errors"

type Description string

func NewDescription(value string) (*Description, error) {
	if len(value) > 200 {
		return nil, errors.New("説明文は200文字以内です")
	}
	if value == "" {
		return nil, errors.New("説明文は必須です")
	}

	description := Description(value)

	return &description, nil
}

func (v Description) String() string {
	return string(v)
}
