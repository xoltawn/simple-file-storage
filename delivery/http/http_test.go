package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/uptrace/bunrouter"
	_http "github.com/xoltawn/simple-file-storage/delivery/http"
)

func TestStoreFromFileHandler(t *testing.T) {
	route := fmt.Sprint(_http.ApiPath, _http.V1Path, _http.FilesPath, _http.StoreFromFilePath)
	t.Run("if no file is updated, it throws 400 error", func(t *testing.T) {
		//arrange
		bunrouter := bunrouter.New()
		rec := httptest.NewRecorder()
		_, req, err := _http.NewFileUploadRequest(route, nil, "", "")
		assert.NoError(t, err)

		_http.NewFileHTTPHandler(bunrouter)

		//act
		bunrouter.ServeHTTP(rec, req)

		//assert
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
