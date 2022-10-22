package file

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var MANGA_PATH string = `D:\DATA\Manga`

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	requestedFilename := strings.TrimPrefix(req.URL.Path, "/")
	if strings.Split(requestedFilename, "/")[0] == "file" {
		requestedFilename = MANGA_PATH + requestedFilename[4:]
		// println("Requesting file:", requestedFilename)
		fileData, err := os.ReadFile(requestedFilename)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
		}

		res.Write(fileData)
	}

}
