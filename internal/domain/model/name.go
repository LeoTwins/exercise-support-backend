package model

import "errors"

type Name string

func NewName(value string) (*Name, error) {
	if len(value) > 30 {
		return nil, errors.New("氏名は30文字以内です")
	}
	if value == "" {
		return nil, errors.New("氏名は必須です")
	}
	name := Name(value)

	return &name, nil
}

func (n Name) String() string {
	return string(n)
}
