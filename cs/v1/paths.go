package v1

import "net/http"

const (
	AccessFileMethod  string = http.MethodGet
	DeleteFileMethod  string = http.MethodDelete
	UploadFileMethod  string = http.MethodPost
	UploadFilesMethod string = http.MethodPost
)

func (s *service) pathAccessFd(appId, group, filename string) string {
	return s.BaseURL + "/" + appId + "/" + group + "/" + filename
}

func (s *service) pathDeleteFd(appId, group, filename string) string {
	return s.BaseURL + "/" + appId + "/" + group + "/" + filename
}

func (s *service) pathUpload() string {
	return s.BaseURL + "/upload"
}

func (s *service) pathUploads() string {
	return s.BaseURL + "/uploads"
}
