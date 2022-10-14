package manga

import (
	"mangav3/apps"
	"mangav3/apps/models"
)

type Manga struct {
}

func NewManga() *Manga {
	return &Manga{}
}

func (m Manga) GetManga() []models.Mangas {
	var manga []models.Mangas
	apps.DB.Order("title").Find(&manga)
	return manga
}
