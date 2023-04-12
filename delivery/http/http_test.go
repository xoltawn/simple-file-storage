package http_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/uptrace/bunrouter"
	_http "github.com/xoltawn/simple-file-storage/delivery/http"
	"github.com/xoltawn/simple-file-storage/domain"
	_mocks "github.com/xoltawn/simple-file-storage/domain/mocks"
)

const (
	maxFileSize = 5
)

var (
	sampleErr = errors.New("sample error")
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
		expErr := sampleErr
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

func TestFetchFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	route := fmt.Sprint(_http.ApiPath, _http.V1Path, _http.FilesPath)

	t.Run("if err occures in file service client, it throws 500 error", func(t *testing.T) {
		//arrange
		fileUsecase := _mocks.NewMockFileUsecase(ctrl)
		expErr := sampleErr
		fileUsecase.EXPECT().FetchFiles(context.TODO(), gomock.Any(), gomock.Any()).Return([]domain.File{}, expErr)
		bunrouter := bunrouter.New()
		req := httptest.NewRequest("GET", route, nil)
		rec := httptest.NewRecorder()
		_http.NewFileHTTPHandler(bunrouter, fileUsecase, maxFileSize)

		//act
		bunrouter.ServeHTTP(rec, req)

		//assert
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("if no err occures in file service client, it throws 200 code with files", func(t *testing.T) {
		//arrange
		expFiles := []domain.File{
			{
				OriginalUrl:   "OriginalUrl1",
				LocalName:     "LocalName1",
				FileExtension: "FileExtension1",
				FileSize:      1,
				CreatedAt:     "CreatedAt1",
			},
		}
		fileUsecase := _mocks.NewMockFileUsecase(ctrl)
		fileUsecase.EXPECT().FetchFiles(context.TODO(), gomock.Any(), gomock.Any()).Return(expFiles, nil)

		bunrouter := bunrouter.New()
		req := httptest.NewRequest("GET", route, nil)
		rec := httptest.NewRecorder()
		rec.Header().Set("Content-Type", "application/json")
		_http.NewFileHTTPHandler(bunrouter, fileUsecase, maxFileSize)

		//act
		bunrouter.ServeHTTP(rec, req)

		//assert
		assert.Equal(t, http.StatusOK, rec.Code)
		wantResp, err := json.Marshal(expFiles)
		if err != nil {
			log.Fatal(err)
		}
		assert.Contains(t, rec.Body.String(), string(wantResp))
	})
}

func TestUploadFile(t *testing.T) {
	route := fmt.Sprint(_http.ApiPath, _http.V1Path, _http.FilesPath)
	defaultContentType := fmt.Sprint("multipart/form-data; boundary=\"bounsdary\"")
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
		req.Header.Add("Content-Type", defaultContentType)
		_http.NewFileHTTPHandler(bunrouter, nil, maxFileSize)

		//act
		bunrouter.ServeHTTP(rec, req)

		//assert
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("if err occures in usecase, it throws 500 error", func(t *testing.T) {
		//arrange
		fileUsecase := _mocks.NewMockFileUsecase(ctrl)
		expErr := sampleErr
		expFile := domain.File{}
		fileUsecase.EXPECT().UploadFile(context.TODO(), gomock.Any()).Return(expFile, expErr)

		bunrouter := bunrouter.New()
		rec := httptest.NewRecorder()
		writer, req, err := _http.NewFileUploadRequest(route, nil, "file", "../../storage/sample-links.txt")
		assert.NoError(t, err)
		req.Header.Add("Content-Type", writer.FormDataContentType())

		_http.NewFileHTTPHandler(bunrouter, fileUsecase, maxFileSize)

		//act
		bunrouter.ServeHTTP(rec, req)

		//assert
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

}
