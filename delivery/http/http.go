package http

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/uptrace/bunrouter"
	"github.com/xoltawn/simple-file-storage/domain"
)

const (
	//APIPath ...
	APIPath = "/api"
	//V1Path ...
	V1Path = "/v1"
	//FilesPath ...
	FilesPath = "/files"
	//StoreFromFilePath ...
	StoreFromFilePath = "/store_from_text"
)

type fileHTTPHandler struct {
	fileUsecase   domain.FileUsecase
	maxFileSizeMB int64
}

// NewFileHTTPHandler is the builder for fileHTTPHandler
func NewFileHTTPHandler(router *bunrouter.Router, fileUsecase domain.FileUsecase, maxFileSizeMB int64) {
	fileHander := fileHTTPHandler{
		fileUsecase:   fileUsecase,
		maxFileSizeMB: maxFileSizeMB,
	}
	router.WithGroup(APIPath, func(apirouter *bunrouter.Group) {
		apirouter.WithGroup(V1Path, func(apirouter *bunrouter.Group) {
			apirouter.WithGroup(FilesPath, func(filesrouter *bunrouter.Group) {
				// /api/v1/files/store_from_text
				filesrouter.POST(fmt.Sprint(StoreFromFilePath), fileHander.storeFromFileHandler)
				// /api/v1/files
				filesrouter.GET("", fileHander.fetchFilesHandler)
				// /api/v1/files
				filesrouter.POST("", fileHander.uploadFileHandler)
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
	err = req.ParseMultipartForm(h.maxFileSizeMB * 1024 * 1024)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, "error reading file")
	}

	multipartFile, _, err := req.Request.FormFile("text_file")
	defer multipartFile.Close()
	if err != nil {
		if err.Error() == "multipart: NextPart: EOF" {
			w.WriteHeader(http.StatusBadRequest)
			return bunrouter.JSON(w, "text_file is not specified")
		}
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, "internal server error")
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, multipartFile); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, "internal server error")
	}

	err = h.fileUsecase.DownloadFromTextFile(req.Context(), buf.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, "internal server error")
	}

	w.WriteHeader(http.StatusOK)
	return bunrouter.JSON(w, "done")
}

func (h *fileHTTPHandler) fetchFilesHandler(w http.ResponseWriter, req bunrouter.Request) (err error) {
	offset := 0
	limit := 10

	keys, ok := req.URL.Query()["offset"]
	if ok && len(keys[0]) > 0 {
		l, err := strconv.Atoi(keys[0])
		if err == nil {
			offset = l
		}
	}

	keys, ok = req.URL.Query()["limit"]
	if ok && len(keys[0]) > 0 {
		l, err := strconv.Atoi(keys[0])
		if err == nil {
			limit = l
		}
	}
	files, err := h.fileUsecase.FetchFiles(req.Context(), limit, offset)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, "internal server error")
	}

	w.WriteHeader(http.StatusOK)
	return bunrouter.JSON(w, files)
}

func (h *fileHTTPHandler) uploadFileHandler(w http.ResponseWriter, req bunrouter.Request) (err error) {
	//check content type
	if ok := HasContentType(req.Request, "multipart/form-data"); !ok {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return bunrouter.JSON(w, "unaccepted content type")
	}

	err = req.ParseMultipartForm(h.maxFileSizeMB * 1024 * 1024)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, "error reading file")
	}

	multipartFile, _, err := req.Request.FormFile("file")
	defer multipartFile.Close()
	if err != nil {
		if err.Error() == "multipart: NextPart: EOF" {
			w.WriteHeader(http.StatusBadRequest)
			return bunrouter.JSON(w, "text_file is not specified")
		}
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, "internal server error")
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, multipartFile); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, "internal server error")
	}

	uploadedFile, err := h.fileUsecase.UploadFile(req.Context(), buf.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, "internal server error")
	}

	w.WriteHeader(http.StatusOK)
	return bunrouter.JSON(w, uploadedFile)
}
