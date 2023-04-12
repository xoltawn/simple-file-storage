package domain

import (
	"context"

	_filepb "github.com/xoltawn/simple-file-storage/repository/grpc/filepb"
)

// File is the domain objects for stored files
type File struct {
	OriginalUrl   string `json:"original_url"`
	LocalName     string `json:"local_name"`
	FileExtension string `json:"file_extension"`
	FileSize      int64  `json:"file_size"`
	CreatedAt     string `json:"created_at"`
}

// FromGRPCFile converts data from *_filepb.File (used in grpc) and converts it to File
func (f *File) FromGRPCFile(in *_filepb.File) *File {
	f.OriginalUrl = in.OriginalUrl
	f.LocalName = in.LocalName
	f.FileExtension = in.FileExtension
	f.FileSize = in.FileSize
	f.CreatedAt = in.CreatedAt
	return f
}

// FileRepository provides an interface regarding operations in the Repository later(CRUD)
//
//go:generate mockgen --destination=mocks/file_repository.go . FileRepository
type FileRepository interface {
	//DownloadFromTextFile takes a text files containg links of files to be downloaded and download them
	DownloadFromTextFile(ctx context.Context, links []byte) (err error)
	//FetchFiles fetches stored files
	FetchFiles(ctx context.Context, limit, offset int) (files []File, err error)
	//UploadFile takes the file to be uploaded and stores the file
	UploadFile(ctx context.Context, file []byte) (uploadedFile File, err error)
}

// FileUsecase
//
//go:generate mockgen --destination=mocks/file_repository.go . FileUsecase
type FileUsecase interface {
	DownloadFromTextFile(ctx context.Context, links []byte) (err error)
	FetchFiles(ctx context.Context, limit, offset int) (files []File, err error)
	UploadFile(ctx context.Context, file []byte) (uploadedFile File, err error)
}
