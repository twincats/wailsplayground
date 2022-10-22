package manga

import (
	"mangav3/apps"
	"mangav3/apps/models"

	"gorm.io/gorm"
)

type Manga struct {
}

func NewManga() *Manga {
	return &Manga{}
}

func (m Manga) GetManga(id string) models.Manga {
	var manga models.Manga
	apps.DB.Order("title").First(&manga, id)
	return manga
}

func (m Manga) GetMangas() []models.Manga {
	var manga []models.Manga
	apps.DB.Order("title").Find(&manga)
	return manga
}

func (m Manga) GetMangaHome(p int, limit int) models.MangaHome {

	mangaHomeApi := []models.MangaHomeApi{}
	latest := apps.DB.Table("chapters").Select("manga_id", "MAX(CAST(chapter AS decimal)) AS chapter").Group("manga_id").Order("manga_id")
	homeApi := apps.DB.Table("mangas").
		Select("mangas.id", "mangas.title", "mdex", "status_ending", "chapters.id as chapter_id", "latest.chapter", "TO_DATE(cast(chapters.created_at as TEXT) ,'YYYY-MM-DD') as download_time").
		// Select("mangas.id, chapters.id").
		Joins("inner join chapters on mangas.id = chapters.manga_id").
		Joins("inner join (?) latest on mangas.id = latest.manga_id", latest).
		Where("latest.chapter = CAST(chapters.chapter as decimal)").
		Order("download_time desc, mangas.title") //Limit(5).

	var pageSize int64
	homeApi.Count(&pageSize)

	pg := models.MPaginationParam{
		Page:  p,
		Limit: limit,
	}.Paginate(int(pageSize))
	Paginate(homeApi, pg.CurrentPage, pg.PerPage).Find(&mangaHomeApi)
	//fill pagination
	var mangahome models.MangaHome
	mangahome.Manga = mangaHomeApi
	mangahome.Pagination = pg

	return mangahome
}

func Paginate(db *gorm.DB, page int, limit int) *gorm.DB {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * int(limit)
	return db.Offset(offset).Limit(limit)
}

// Page for Fetching chapter and Manga title
type Page struct {
	ID      int
	Chapter string
	Title   string
}

/* GetPage for Fetching list of images chapter */
func (f Manga) GetPage(id int) interface{} {
	var p Page
	apps.DB.Table("chapters").
		Select("chapters.id", "chapter", "mangas.title").
		Joins("inner join mangas on chapters.manga_id = mangas.id").
		Where("chapters.id = ?", id).
		Take(&p)
	// fmt.Println(p)
	// var page models.Page
	// page.Path = filepath.Join(p.title, strconv.Itoa(p.chapter))
	return "Title : " + p.Title
}
