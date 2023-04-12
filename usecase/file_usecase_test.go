package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xoltawn/simple-file-storage/domain"
	_mocks "github.com/xoltawn/simple-file-storage/domain/mocks"
	"github.com/xoltawn/simple-file-storage/usecase"
)

var (
	sampleRPCErr = errors.New("An rpc error from file service client")
)

func TestDownloadFromTextFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("if err occurres, the err is returned", func(t *testing.T) {
		//arrange
		fileBytes := []byte(gomock.Any().String())
		fileRepo := _mocks.NewMockFileRepository(ctrl)
		expErr := sampleRPCErr
		fileRepo.EXPECT().DownloadFromTextFile(context.TODO(), fileBytes).Return(expErr)

		//act
		sut := usecase.NewFileUsecase(fileRepo)
		err := sut.DownloadFromTextFile(context.TODO(), fileBytes)

		//assert
		assert.Error(t, err)
	})

	t.Run("if no err occurres, nil will be returned for err", func(t *testing.T) {
		//arrange
		fileBytes := []byte(gomock.Any().String())
		fileRepo := _mocks.NewMockFileRepository(ctrl)
		fileRepo.EXPECT().DownloadFromTextFile(context.TODO(), fileBytes).Return(nil)

		//act
		sut := usecase.NewFileUsecase(fileRepo)
		err := sut.DownloadFromTextFile(context.TODO(), fileBytes)

		//assert
		assert.NoError(t, err)
	})
}

func TestFetchFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	limit := 10
	offset := 0

	t.Run("if err occurres, the err is returned with files being empty slice", func(t *testing.T) {
		//arrange
		expErr := sampleRPCErr
		expFiles := []domain.File{}
		fileRepo := _mocks.NewMockFileRepository(ctrl)

		fileRepo.EXPECT().FetchFiles(context.TODO(), gomock.Any(), gomock.Any()).Return(expFiles, expErr)

		//act
		sut := usecase.NewFileUsecase(fileRepo)
		actFiles, err := sut.FetchFiles(context.TODO(), limit, offset)

		//assert
		assert.Error(t, err)
		assert.Equal(t, expFiles, actFiles)
	})

	t.Run("if no err occurres, files will be returned with err being nil", func(t *testing.T) {
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
		fileRepo := _mocks.NewMockFileRepository(ctrl)

		fileRepo.EXPECT().FetchFiles(context.TODO(), gomock.Any(), gomock.Any()).Return(expFiles, nil)

		//act
		sut := usecase.NewFileUsecase(fileRepo)
		actFiles, err := sut.FetchFiles(context.TODO(), limit, offset)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, expFiles, actFiles)
	})

}
