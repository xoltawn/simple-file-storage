package grpc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	_grpc "github.com/xoltawn/simple-file-storage/repository/grpc"
	_filepb "github.com/xoltawn/simple-file-storage/repository/grpc/filepb"
	_grpcmocks "github.com/xoltawn/simple-file-storage/repository/grpc/mocks"
)

func TestDownloadFromTextFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	dummylinkBytes := []byte{}

	t.Run("success", func(t *testing.T) {
		//arrange
		res := &_filepb.DownloadFromTextFileResponse{}
		fileClient := _grpcmocks.NewMockFileServiceClient(ctrl)
		fileClient.EXPECT().DownloadFromTextFile(context.TODO(), gomock.Any()).Return(res, nil)

		//act
		repo := _grpc.NewFileGRPCRepository(fileClient)
		err := repo.DownloadFromTextFile(context.TODO(), dummylinkBytes)

		//assert
		assert.NoError(t, err)
	})

	t.Run("handle error from file service client", func(t *testing.T) {
		//arrange
		res := &_filepb.DownloadFromTextFileResponse{}
		expErr := errors.New("An rpc error from file service client")
		fileClient := _grpcmocks.NewMockFileServiceClient(ctrl)
		fileClient.EXPECT().DownloadFromTextFile(context.TODO(), gomock.Any()).Return(res, expErr)

		//act
		repo := _grpc.NewFileGRPCRepository(fileClient)
		err := repo.DownloadFromTextFile(context.TODO(), dummylinkBytes)

		//assert
		assert.Error(t, err)
		assert.Equal(t, expErr, err)
	})
}
