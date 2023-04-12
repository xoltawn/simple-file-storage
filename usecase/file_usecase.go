package usecase

import (
	"context"

	"github.com/xoltawn/simple-file-storage/domain"
)

type fileUsecase struct {
	fileRepo domain.FileRepository
}

func NewFileUsecase(fileRepo domain.FileRepository) *fileUsecase {
	return &fileUsecase{
		fileRepo: fileRepo,
	}
}

func (fu *fileUsecase) DownloadFromTextFile(ctx context.Context, links []byte) (err error) {
	return fu.fileRepo.DownloadFromTextFile(ctx, links)
}
