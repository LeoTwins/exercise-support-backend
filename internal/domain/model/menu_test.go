package model_test

import (
	"testing"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewMenu_Success(t *testing.T) {
	_, err := model.NewMenu(
		nil,
		model.Title("title"),
		"image",
		model.Description("description"),
	)

	assert.NoError(t, err)
}
