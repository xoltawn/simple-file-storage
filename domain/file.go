package domain

import "context"

//go:generate mockgen --destination=mocks/file_repository.go . FileRepository
type FileRepository interface {
	DownloadFromTextFile(ctx context.Context, links []byte) (err error)
}
