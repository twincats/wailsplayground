package apps

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/climech/naturalsort"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
	db  *gorm.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	runtime.WindowSetTitle(ctx, "MangaV3.5")
	// tit, _ := syscall.UTF16PtrFromString("waislplayground")
	// hwnd := win.FindWindow(nil, tit)
	// win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.GetWindowLong(hwnd, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (app *App) Find(root string, ext []string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		for _, i := range ext {
			if filepath.Ext(d.Name()) == i {
				a = append(a, s)
			}
		}
		return nil
	})
	naturalsort.Sort(a)
	return a
}

//internal
func (a *App) ConnectDatabaseSqlite(dbname string) {
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		panic("failed to connect database " + dbname)
	}
	a.db = db
}

func (a *App) ConnectDatabasePostgres(dbname string) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", "127.0.0.1", "achul", "1234", dbname, 5432)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database " + dbname)
	}
	a.db = db
}

type Manga struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	StatusEnd bool   `json:"status_end"`
	Mdex      string `json:"mdex"`
	CreatedAt string `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt string `gorm:"type:timestamp" json:"updated_at"`
}

func (a *App) GetManga() []Manga {
	var mangas []Manga
	a.db.Find(&mangas)
	return mangas
}
