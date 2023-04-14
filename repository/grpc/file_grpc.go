package grpc

import (
	"context"
	"fmt"

	"github.com/xoltawn/simple-file-storage/domain"
	_filepb "github.com/xoltawn/simple-file-storage/repository/grpc/filepb"
)

//go:generate mockgen --source=filepb/file_grpc.pb.go --destination=mocks/file_service_client.go . FileServiceClient
type fileGRPCRepository struct {
	fileServiceClient _filepb.FileServiceClient
	gateWayAddress    string
}

// NewFileGRPCRepository is the builder function for fileGRPC
func NewFileGRPCRepository(fsc _filepb.FileServiceClient, gateWayAddress string) *fileGRPCRepository {
	return &fileGRPCRepository{
		fileServiceClient: fsc,
		gateWayAddress:    gateWayAddress,
	}
}

func (f *fileGRPCRepository) DownloadFromTextFile(ctx context.Context, links []byte) (err error) {
	req := &_filepb.DownloadFromTextFileRequest{
		Links: links,
	}
	_, err = f.fileServiceClient.DownloadFromTextFile(ctx, req)
	return
}

func (f *fileGRPCRepository) FetchFiles(ctx context.Context, limit, offset int) (files []domain.File, err error) {
	req := &_filepb.FetchFilesRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	res, err := f.fileServiceClient.FetchFiles(ctx, req)
	if err != nil {
		return
	}

	for _, file := range res.Files {
		resFile := &domain.File{}
		resFile.URL = fmt.Sprint(f.gateWayAddress, "/", file.FileLocation, "/", file.LocalName, ".", file.FileExtension)
		files = append(files, *resFile.FromGRPCFile(file))
	}
	return
}

func (f *fileGRPCRepository) UploadFile(ctx context.Context, file []byte) (uploadedFile domain.File, err error) {
	res, err := f.fileServiceClient.UploadFile(ctx, &_filepb.UploadFileRequest{
		File: file,
	})
	if err != nil {
		return
	}

	uploadedFile.FromGRPCFile(res.File)
	return
}
