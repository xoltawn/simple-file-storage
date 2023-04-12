package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
}
