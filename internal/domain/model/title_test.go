package model_test

import (
	"strings"
	"testing"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewTitle_Success(t *testing.T) {
	v := "title"
	_, err := model.NewTitle(v)

	assert.NoError(t, err)
}

func TestNewTitle_Success_30文字(t *testing.T) {
	v := strings.Repeat("s", 30)
	_, err := model.NewTitle(v)

	assert.NoError(t, err)
}

func TestNewTitle_Failure_31文字(t *testing.T) {
	v := strings.Repeat("s", 31)
	_, err := model.NewTitle(v)

	assert.Error(t, err)
	assert.EqualError(t, err, "タイトルは30文字以内です")
}

func TestNewTitle_Failure_Empty(t *testing.T) {
	v := ""
	_, err := model.NewTitle(v)

	assert.Error(t, err)
	assert.EqualError(t, err, "タイトルは必須です")
}
