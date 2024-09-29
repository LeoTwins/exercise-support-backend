package model

type Menu struct {
	ID          *uint
	Title       Title
	Image       string
	Description Description
}

func NewMenu(id *uint, title Title, image string, description Description) (*Menu, error) {
	return &Menu{
		ID:          id,
		Title:       title,
		Image:       image,
		Description: description,
	}, nil
}

func ReCreateMenu(id *uint, title Title, image string, description Description) *Menu {
	return &Menu{
		ID:          id,
		Title:       title,
		Image:       image,
		Description: description,
	}
}
