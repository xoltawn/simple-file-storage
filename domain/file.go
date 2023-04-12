package domain

//go:generate mockgen --destination=mocks/file_repository.go . FileRepository
type FileRepository interface {
	DownloadFromTextFile(links []byte) error
}
