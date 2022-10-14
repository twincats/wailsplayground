package models

import (
	"github.com/SamuelTissot/sqltime"
)

type Mangas struct {
	ID         int          `json:"id"`
	Title      string       `json:"title"`
	StatusEnd  bool         `json:"status_end"`
	Mdex       *int         `orm:"null" json:"mdex"`
	Created_at sqltime.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	Updated_at sqltime.Time `orm:"auto_now;type(datetime)" json:"-"`
}
