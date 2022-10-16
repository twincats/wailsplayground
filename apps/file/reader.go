package file

import "os"

type FileReader struct {
}

func NewFileReader() *FileReader {
	return &FileReader{}
}

func (f FileReader) GetFile(path string) []string {
	return []string{}
}

func (f FileReader) GetFileManga() []string {
	dirs, _ := os.ReadDir(MANGA_PATH)
	var manga []string
	for _, i := range dirs {
		manga = append(manga, i.Name())
	}
	return manga
}
