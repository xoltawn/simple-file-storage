package grpc

import (
	"context"

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

func (f *fileGRPC) FetchFiles(ctx context.Context, limit, offset int) (err error) {
	return
}
