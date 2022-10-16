package models

import (
	"encoding/json"

	"github.com/SamuelTissot/sqltime"
)

type Manga struct {
	ID           int          `json:"id"`
	Title        string       `json:"title"`
	StatusEnding bool         `json:"status_end"`
	Mdex         *int         `orm:"null" json:"mdex"`
	CreatedAt    sqltime.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	UpdatedAt    sqltime.Time `orm:"auto_now;type(datetime)" json:"-"`
}

type MangaHomeApi struct {
	ID           int          `json:"id"`
	Title        string       `json:"title"`
	StatusEnding bool         `json:"status_end"`
	Mdex         *int         `orm:"null" json:"mdex"`
	ChapterID    int          `json:"chapter_id"`
	Chapter      json.Number  `json:"chapter"`
	DownloadTime sqltime.Time `json:"download_time"`
}

type MangaHome struct {
	Manga      []MangaHomeApi `json:"manga"`
	Pagination *Pagination    `json:"pagination"`
}

type MPaginationParam struct {
	Page  int
	Limit int
}

func (p MPaginationParam) Paginate(totalPage int) *Pagination {
	//set to
	to := p.Page * p.Limit
	if to > totalPage {
		to = totalPage
	}
	// set last
	last := totalPage / p.Limit
	if totalPage%p.Limit > 0 {
		last += 1
	}

	//limit page to available page
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Page > last {
		p.Page = last
	}
	//return pagination output
	return &Pagination{
		CurrentPage: p.Page,
		From:        p.Limit*p.Page - (p.Limit - 1),
		To:          to,
		Total:       totalPage,
		Last_page:   last,
		PerPage:     p.Limit,
	}
}

type Pagination struct {
	CurrentPage int
	From        int
	To          int
	Total       int
	Last_page   int
	PerPage     int
}
