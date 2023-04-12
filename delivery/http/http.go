package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
	"github.com/xoltawn/simple-file-storage/domain"
)

const (
	ApiPath           = "/api"
	V1Path            = "/v1"
	FilesPath         = "/files"
	StoreFromFilePath = "/store_from_text"
)

type fileHTTPHandler struct {
	fileUsecase   domain.FileUsecase
	maxFileSizeMB int64
}

func NewFileHTTPHandler(router *bunrouter.Router, fileUsecase domain.FileUsecase, maxFileSizeMB int64) {
	fileHander := fileHTTPHandler{
		fileUsecase:   fileUsecase,
		maxFileSizeMB: maxFileSizeMB,
	}
	router.WithGroup(ApiPath, func(apirouter *bunrouter.Group) {
		apirouter.WithGroup(V1Path, func(apirouter *bunrouter.Group) {
			apirouter.WithGroup(FilesPath, func(filesrouter *bunrouter.Group) {
				filesrouter.POST(fmt.Sprint(StoreFromFilePath), fileHander.storeFromFileHandler)
			})
		})
	})
}

func (h *fileHTTPHandler) storeFromFileHandler(w http.ResponseWriter, req bunrouter.Request) (err error) {

	//check content type
	if ok := HasContentType(req.Request, "multipart/form-data"); !ok {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return bunrouter.JSON(w, "unaccepted content type")
	}

	//check if the text file containing links is sent and get file content
	err = req.Request.ParseMultipartForm(h.maxFileSizeMB * 1024 * 1024)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, "error reading file")
	}

	_, _, err = req.Request.FormFile("text_file")
	if err != nil {
		log.Println(err)
		if err.Error() == "multipart: NextPart: EOF" {
			w.WriteHeader(http.StatusBadRequest)
			return bunrouter.JSON(w, "text_file is not specified")
		}
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, "internal server error")
	}

	w.WriteHeader(http.StatusOK)
	return bunrouter.JSON(w, "done")
}
