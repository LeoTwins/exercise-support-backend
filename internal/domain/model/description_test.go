package model_test

import (
	"strings"
	"testing"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewDescription_Success(t *testing.T) {
	_, err := model.NewDescription("description")

	assert.NoError(t, err)
}

func TestNewDescription_Failure_Empty(t *testing.T) {
	_, err := model.NewDescription("")

	assert.Error(t, err)
	assert.EqualError(t, err, "説明文は必須です")
}

func TestNewDescription_Success_200文字(t *testing.T) {
	v := strings.Repeat("s", 200)
	_, err := model.NewDescription(v)

	assert.NoError(t, err)
}
func TestNewDescription_Failure_201文字(t *testing.T) {
	v := strings.Repeat("s", 201)
	_, err := model.NewDescription(v)

	assert.Error(t, err)
	assert.EqualError(t, err, "説明文は200文字以内です")
}
