package domain

import "context"

// File is the domain objects for stored files
type File struct {
	OriginalUrl   string `json:"original_url"`
	LocalName     string `json:"local_name"`
	FileExtension string `json:"file_extension"`
	FileSize      int64  `json:"file_size"`
	DownloadDate  string `json:"download_date"`
}

// FileRepository provides an interface regarding operations in the Repository later(CRUD)
//
//go:generate mockgen --destination=mocks/file_repository.go . FileRepository
type FileRepository interface {
	DownloadFromTextFile(ctx context.Context, links []byte) (err error)
	FetchFiles(ctx context.Context, limit, offset int) (files []File, err error)
}
