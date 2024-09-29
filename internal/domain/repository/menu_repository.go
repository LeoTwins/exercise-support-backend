package repository

import "github.com/LeoTwins/go-clean-architecture/internal/domain/model"

type IMenuRepository interface {
	Find()
	FindByIDs(menuIDs []uint)
	Save(item *model.Menu) error
	Update(item *model.Menu) error
	Delete(ID uint) error
}
