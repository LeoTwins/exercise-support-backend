package model_test

import (
	"strings"
	"testing"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewName_Success(t *testing.T) {
	v := "test"
	_, err := model.NewName(v)
	assert.NoError(t, err)
}

func TestNewName_Failure_Empty(t *testing.T) {
	v := ""
	_, err := model.NewName(v)
	assert.Error(t, err)
	assert.EqualError(t, err, "氏名は必須です")
}

func TestNewName_Success_30文字(t *testing.T) {
	v := strings.Repeat("s", 30)
	_, err := model.NewName(v)
	assert.NoError(t, err)
}
func TestNewName_Failure_31文字(t *testing.T) {
	v := strings.Repeat("s", 31)
	_, err := model.NewName(v)
	assert.Error(t, err)
	assert.EqualError(t, err, "氏名は30文字以内です")
}
