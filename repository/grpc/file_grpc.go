package grpc

import (
	"context"

	"github.com/xoltawn/simple-file-storage/domain"
	_filepb "github.com/xoltawn/simple-file-storage/repository/grpc/filepb"
)

//go:generate mockgen --source=filepb/file_grpc.pb.go --destination=mocks/file_service_client.go . FileServiceClient
type fileGRPC struct {
	fileServiceClient _filepb.FileServiceClient
}

func NewFileGRPCRepository(fsc _filepb.FileServiceClient) *fileGRPC {
	return &fileGRPC{
		fileServiceClient: fsc,
	}
}

func (f *fileGRPC) DownloadFromTextFile(ctx context.Context, links []byte) (err error) {
	req := &_filepb.DownloadFromTextFileRequest{
		Links: links,
	}
	_, err = f.fileServiceClient.DownloadFromTextFile(ctx, req)
	return
}

func (f *fileGRPC) FetchFiles(ctx context.Context, limit, offset int) (files []domain.File, err error) {
	req := &_filepb.FetchFilesRequest{
		Limit:  10,
		Offset: 0,
	}
	res, err := f.fileServiceClient.FetchFiles(ctx, req)
	if err != nil {
		return
	}

	for _, f := range res.Files {
		resFile := &domain.File{}
		files = append(files, *resFile.FromGRPCFile(f))
	}
	return
}

func (f *fileGRPC) UploadFile(ctx context.Context, file []byte) (uploadedFile domain.File, err error) {
	res, err := f.fileServiceClient.UploadFile(ctx, &_filepb.UploadFileRequest{
		File: file,
	})
	if err != nil {
		return
	}

	uploadedFile.FromGRPCFile(res.File)
	return
}
