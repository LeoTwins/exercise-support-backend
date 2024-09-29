package model_test

import (
	"testing"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewProfile_Success_実施秒数の指定が０(t *testing.T) {
	v, err := model.NewProfile(
		nil,
		model.Name("v"),
		"image",
		0,
	)

	assert.NoError(t, err)
	assert.Equal(t, v.ImplementationSeconds, uint(60))
}

func TestNewProfile_Success_実施秒数指定(t *testing.T) {
	v, err := model.NewProfile(
		nil,
		model.Name("v"),
		"image",
		120,
	)

	assert.NoError(t, err)
	assert.Equal(t, v.ImplementationSeconds, uint(120))
}

func TestNewProfile_Failure(t *testing.T) {
	_, err := model.NewProfile(
		nil,
		model.Name("v"),
		"",
		0,
	)

	assert.Error(t, err)
	assert.EqualError(t, err, "画像イメージは必須です")
}
