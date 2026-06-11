package client

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func GetDownloadLink(token string, filePath string) string {
	return fmt.Sprintf("%s/file/bot%s/%s", baseUrl, token, filePath)
}

// This object represents the contents of a file to be uploaded. Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
type InputFile interface {
	Close()
	GetReader() io.Reader
	Name() string
	IsStream() bool
}

// If the file is already stored somewhere on the Telegram servers, you don't need to reupload it: each file object has a file_id field, simply pass this file_id as a parameter instead of uploading. There are no limits for files sent this way.
type FileIdInputFile struct {
	FileId string
}

func NewFileIdInputFile(fileId string) *FileIdInputFile {
	return &FileIdInputFile{
		FileId: fileId,
	}
}

func (FileIdInputFile) Close() {}

func (inputFile *FileIdInputFile) GetReader() io.Reader {
	return strings.NewReader(inputFile.FileId)
}

func (inputFile *FileIdInputFile) Name() string {
	return inputFile.FileId
}

func (inputFile *FileIdInputFile) IsStream() bool {
	return false
}

// Provide Telegram with an HTTP URL for the file to be sent. Telegram will download and send the file. 5 MB max size for photos and 20 MB max for other types of content.
type UrlInputFile struct {
	Url string
}

func NewUrlInputFile(url string) *UrlInputFile {
	return &UrlInputFile{
		Url: url,
	}
}

func (UrlInputFile) Close() {}

func (inputFile *UrlInputFile) GetReader() io.Reader {
	return strings.NewReader(inputFile.Url)
}

func (inputFile *UrlInputFile) Name() string {
	return inputFile.Url
}

func (inputFile *UrlInputFile) IsStream() bool {
	return false
}

// Post the file using multipart/form-data in the usual way that files are uploaded via the browser. 10 MB max size for photos, 50 MB for other files.
type FileInputFile struct {
	Reader   io.Reader
	FileName string
}

func NewFileInputFile(path string) (*FileInputFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return &FileInputFile{
		Reader:   file,
		FileName: stat.Name(),
	}, nil
}

func (inputFile *FileInputFile) Close() {
	rc, ok := inputFile.Reader.(io.ReadCloser)
	if ok {
		rc.Close()
	}
}

func (inputFile *FileInputFile) GetReader() io.Reader {
	return inputFile.Reader
}

func (inputFile *FileInputFile) Name() string {
	return inputFile.FileName
}

func (inputFile *FileInputFile) IsStream() bool {
	return true
}

// This object represents the content of a media message to be sent. It should be one of InputMediaAnimation, InputMediaDocument, InputMediaAudio, InputMediaPhoto, InputMediaVideo
type InputMedia interface{}

// InputMediaPhoto and InputMediaVideo
type InputMediaGroup interface{}
