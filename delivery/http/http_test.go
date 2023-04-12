package http_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/uptrace/bunrouter"
	_http "github.com/xoltawn/simple-file-storage/delivery/http"
	_mocks "github.com/xoltawn/simple-file-storage/domain/mocks"
)

const (
	maxFileSize = 5
)

func TestStoreFromFileHandler(t *testing.T) {
	route := fmt.Sprint(_http.ApiPath, _http.V1Path, _http.FilesPath, _http.StoreFromFilePath)
	contentType := fmt.Sprint("multipart/form-data; boundary=\"bounsdary\"")
	ctrl := gomock.NewController(t)

	t.Run("if request content type is not multipart, it throws 415 error", func(t *testing.T) {
		//arrange
		bunrouter := bunrouter.New()
		rec := httptest.NewRecorder()
		_, req, err := _http.NewFileUploadRequest(route, nil, "", "")
		assert.NoError(t, err)

		_http.NewFileHTTPHandler(bunrouter, nil, maxFileSize)

		//act
		bunrouter.ServeHTTP(rec, req)

		//assert
		assert.Equal(t, http.StatusUnsupportedMediaType, rec.Code)
	})
	t.Run("if no file is uploaded, it throws 400 error", func(t *testing.T) {
		//arrange
		bunrouter := bunrouter.New()
		rec := httptest.NewRecorder()
		_, req, err := _http.NewFileUploadRequest(route, nil, "", "")
		assert.NoError(t, err)
		req.Header.Add("Content-Type", contentType)
		_http.NewFileHTTPHandler(bunrouter, nil, maxFileSize)

		//act
		bunrouter.ServeHTTP(rec, req)

		//assert
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("if err occures in file service client, it throws 500 error", func(t *testing.T) {
		//arrange
		fileUsecase := _mocks.NewMockFileUsecase(ctrl)
		expErr := errors.New("sample error")
		fileUsecase.EXPECT().DownloadFromTextFile(context.TODO(), gomock.Any()).Return(expErr)

		bunrouter := bunrouter.New()
		rec := httptest.NewRecorder()
		writer, req, err := _http.NewFileUploadRequest(route, nil, "text_file", "../../storage/sample-links.txt")
		assert.NoError(t, err)
		req.Header.Add("Content-Type", writer.FormDataContentType())

		_http.NewFileHTTPHandler(bunrouter, fileUsecase, maxFileSize)

		//act
		bunrouter.ServeHTTP(rec, req)

		//assert
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("if no err occures, it throws 200 error", func(t *testing.T) {
		//arrange
		fileUsecase := _mocks.NewMockFileUsecase(ctrl)
		fileUsecase.EXPECT().DownloadFromTextFile(context.TODO(), gomock.Any()).Return(nil)

		bunrouter := bunrouter.New()
		rec := httptest.NewRecorder()
		writer, req, err := _http.NewFileUploadRequest(route, nil, "text_file", "../../storage/sample-links.txt")
		assert.NoError(t, err)
		req.Header.Add("Content-Type", writer.FormDataContentType())

		_http.NewFileHTTPHandler(bunrouter, fileUsecase, maxFileSize)

		//act
		bunrouter.ServeHTTP(rec, req)

		//assert
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
