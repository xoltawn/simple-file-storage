package http

import (
	"fmt"
	"net/http"

	"github.com/uptrace/bunrouter"
)

const (
	ApiPath           = "/api"
	V1Path            = "/v1"
	FilesPath         = "/files"
	StoreFromFilePath = "/store_from_text"
)

type fileHTTPHandler struct {
}

func NewFileHTTPHandler(router *bunrouter.Router) {
	fileHander := fileHTTPHandler{}
	router.WithGroup(ApiPath, func(apirouter *bunrouter.Group) {
		apirouter.WithGroup(V1Path, func(apirouter *bunrouter.Group) {
			apirouter.WithGroup(FilesPath, func(filesrouter *bunrouter.Group) {
				filesrouter.POST(fmt.Sprint(StoreFromFilePath), fileHander.storeFromFileHandler)
			})
		})
	})
}

func (h *fileHTTPHandler) storeFromFileHandler(w http.ResponseWriter, req bunrouter.Request) (err error) {

	if ok := HasContentType(req.Request, "multipart/form-data"); !ok {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return bunrouter.JSON(w, "unaccepted content type")
	}

	//check if the text file containing links is sent and get file content
	_, _, err = req.Request.FormFile("text_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, "text_file is not specified")
	}

	return
}
