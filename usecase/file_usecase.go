package usecase

import (
	"context"

	"github.com/xoltawn/simple-file-storage/domain"
)

type fileUsecase struct {
	fileRepo domain.FileRepository
}

// NewFileUsecase is the builder function for fileUsecase
func NewFileUsecase(fileRepo domain.FileRepository) *fileUsecase {
	return &fileUsecase{
		fileRepo: fileRepo,
	}
}

func (fu *fileUsecase) DownloadFromTextFile(ctx context.Context, links []byte) (err error) {
	return fu.fileRepo.DownloadFromTextFile(ctx, links)
}

func (fu *fileUsecase) FetchFiles(ctx context.Context, limit, offset int) (files []domain.File, err error) {
	return fu.fileRepo.FetchFiles(ctx, limit, offset)
}

func (fu *fileUsecase) UploadFile(ctx context.Context, file []byte) (uploadedFile domain.File, err error) {
	return fu.fileRepo.UploadFile(ctx, file)
}
