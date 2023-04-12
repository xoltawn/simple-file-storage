package http

import (
	"bytes"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func NewFileUploadRequest(uri string, params map[string]string, paramName, path string) (*multipart.Writer, *http.Request, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	if paramName != "" && path != "" {
		file, err := os.Open(path)
		if err != nil {
			return nil, nil, err
		}
		fileContents, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, nil, err
		}
		fi, err := file.Stat()
		if err != nil {
			return nil, nil, err
		}
		file.Close()

		part, err := writer.CreateFormFile(paramName, fi.Name())
		if err != nil {
			return nil, nil, err
		}
		_, err = part.Write(fileContents)
		if err != nil {
			return nil, nil, err
		}
	}
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	return writer, req, err
}

func HasContentType(r *http.Request, mimetype string) bool {
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		return mimetype == "application/octet-stream"
	}

	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}
	return false
}
