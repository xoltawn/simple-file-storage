package grpc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xoltawn/simple-file-storage/domain"
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
		repo := _grpc.NewFileGRPCRepository(fileClient, gomock.Any().String())
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
		repo := _grpc.NewFileGRPCRepository(fileClient, gomock.Any().String())
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
		sut := _grpc.NewFileGRPCRepository(fileClient, gomock.Any().String())
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
				CreatedAt:     "CreatedAt1",
			},
			{
				OriginalUrl:   "OriginalUrl2",
				LocalName:     "LocalName2",
				FileExtension: "FileExtension2",
				FileSize:      2,
				CreatedAt:     "CreatedAt2",
			},
		}
		fcsRes := &_filepb.FetchFilesResponse{
			Files: fcsFiles,
		}
		fileClient := _grpcmocks.NewMockFileServiceClient(ctrl)
		fileClient.EXPECT().FetchFiles(context.TODO(), gomock.Any()).Return(fcsRes, nil)

		//act
		sut := _grpc.NewFileGRPCRepository(fileClient, gomock.Any().String())
		res, err := sut.FetchFiles(context.TODO(), limit, offset)

		//assert
		assert.NoError(t, err)
		for i, file := range res {
			assert.Equal(t, file.OriginalURL, fcsFiles[i].OriginalUrl)
			assert.Equal(t, file.LocalName, fcsFiles[i].LocalName)
			assert.Equal(t, file.FileExtension, fcsFiles[i].FileExtension)
			assert.Equal(t, file.FileSize, fcsFiles[i].FileSize)
			assert.Equal(t, file.CreatedAt, fcsFiles[i].CreatedAt)
		}
	})
}

func TestUploadFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("if err occurs in fcs client, the err will be returned with empty File object", func(t *testing.T) {
		//arrange
		fscRes := &_filepb.UploadFileResponse{}
		fileClient := _grpcmocks.NewMockFileServiceClient(ctrl)
		expErr := sampleRPCErr
		fileClient.EXPECT().UploadFile(context.TODO(), gomock.Any()).Return(fscRes, expErr)

		//act
		sut := _grpc.NewFileGRPCRepository(fileClient, gomock.Any().String())
		uploadedFile, err := sut.UploadFile(context.TODO(), []byte(gomock.Any().String()))

		//assert
		assert.Error(t, err)
		assert.Equal(t, uploadedFile, domain.File{})
	})

	t.Run("if no err occurs in fcs client, newly updated will be returned", func(t *testing.T) {
		//arrange
		fileBytes := []byte(gomock.Any().String())
		pbFile := &_filepb.File{}
		fscRes := &_filepb.UploadFileResponse{
			File: pbFile,
		}
		fscReq := &_filepb.UploadFileRequest{
			File: fileBytes,
		}

		expFile := domain.File{}
		expFile.FromGRPCFile(pbFile)

		fileClient := _grpcmocks.NewMockFileServiceClient(ctrl)
		fileClient.EXPECT().UploadFile(context.TODO(), fscReq).Return(fscRes, nil)

		//act
		sut := _grpc.NewFileGRPCRepository(fileClient, gomock.Any().String())
		uploadedFile, err := sut.UploadFile(context.TODO(), fileBytes)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, uploadedFile, expFile)
	})
}
