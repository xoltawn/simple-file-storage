package domain

import "context"

// FileRepository provides an interface regarding operations in the Repository later(CRUD)
//
//go:generate mockgen --destination=mocks/file_repository.go . FileRepository
type FileRepository interface {
	DownloadFromTextFile(ctx context.Context, links []byte) (err error)
	FetchFiles(ctx context.Context, limit, offset int) (err error)
}
