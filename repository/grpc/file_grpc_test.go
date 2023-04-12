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

var (
	sampleRPCErr = errors.New("An rpc error from file service client")
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
		expErr := sampleRPCErr
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

func TestFetchFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	limit := 10
	offset := 0
	t.Run("if error occurs in file client, it FetchFiles should returns err", func(t *testing.T) {
		//arrange
		fcsRes := &_filepb.FetchFilesResponse{}
		expErr := sampleRPCErr
		fileClient := _grpcmocks.NewMockFileServiceClient(ctrl)
		fileClient.EXPECT().FetchFiles(context.TODO(), gomock.Any()).Return(fcsRes, expErr)

		//act
		sut := _grpc.NewFileGRPCRepository(fileClient)
		res, err := sut.FetchFiles(context.TODO(), limit, offset)

		//assert
		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("if no error occurs in file client, FetchFiles should returns fetch Files", func(t *testing.T) {
		//arrange
		fcsFiles := []*_filepb.File{
			{
				OriginalUrl:   "OriginalUrl1",
				LocalName:     "LocalName1",
				FileExtension: "FileExtension1",
				FileSize:      1,
				DownloadDate:  "DownloadDate1",
			},
			{
				OriginalUrl:   "OriginalUrl2",
				LocalName:     "LocalName2",
				FileExtension: "FileExtension2",
				FileSize:      2,
				DownloadDate:  "DownloadDate2",
			},
		}
		fcsRes := &_filepb.FetchFilesResponse{
			Files: fcsFiles,
		}
		fileClient := _grpcmocks.NewMockFileServiceClient(ctrl)
		fileClient.EXPECT().FetchFiles(context.TODO(), gomock.Any()).Return(fcsRes, nil)

		//act
		sut := _grpc.NewFileGRPCRepository(fileClient)
		res, err := sut.FetchFiles(context.TODO(), limit, offset)

		//assert
		assert.NoError(t, err)
		for i, file := range res {
			assert.Equal(t, file.OriginalUrl, fcsFiles[i].OriginalUrl)
			assert.Equal(t, file.LocalName, fcsFiles[i].LocalName)
			assert.Equal(t, file.FileExtension, fcsFiles[i].FileExtension)
			assert.Equal(t, file.FileSize, fcsFiles[i].FileSize)
			assert.Equal(t, file.DownloadDate, fcsFiles[i].DownloadDate)
		}
	})
}
