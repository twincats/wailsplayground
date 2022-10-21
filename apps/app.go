package apps

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

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

}

func (f App) GetENV() runtime.EnvironmentInfo {
	return runtime.Environment(f.ctx)
}

func (f App) GetUserDataPath() string {
	userDir := "Appdata"

	exePath, _ := os.Executable()
	exeDir, exeName := filepath.Split(exePath)
	baseExe := strings.Split(exeName, "-")

	if len(baseExe) > 1 {
		// userDir = filepath.Join("%APPDATA%", exeName)
		userDir += "-dev"
	}

	return filepath.Join(exeDir, userDir)

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
