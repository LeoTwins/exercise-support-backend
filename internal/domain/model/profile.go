package model

import "errors"

type Profile struct {
	UserID                *uint
	Name                  Name
	Image                 string
	ImplementationSeconds uint
}

func NewProfile(userID *uint, name Name, image string, implementationSeconds uint) (*Profile, error) {
	if image == "" {
		return nil, errors.New("画像イメージは必須です")
	}

	var sec uint
	if implementationSeconds == 0 {
		sec = 60
	} else {
		sec = implementationSeconds
	}

	return &Profile{
		UserID:                userID,
		Name:                  name,
		Image:                 image,
		ImplementationSeconds: sec,
	}, nil
}

func ReCreateProfile(userID *uint, name Name, image string, implementationSeconds uint) *Profile {
	return &Profile{
		UserID:                userID,
		Name:                  name,
		Image:                 image,
		ImplementationSeconds: implementationSeconds,
	}
}
