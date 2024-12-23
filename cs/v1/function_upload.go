package v1

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// Upload uploads a file and assigns it to the given group.
func (s *service) Upload(group string, mpfh *multipart.FileHeader,
) (url string, code int, err error) {
	// creates a new multipart writer
	var reqb bytes.Buffer
	writer := multipart.NewWriter(&reqb)

	// write the group name into the form group field
	wgroup, err := writer.CreateFormField("group")
	if err != nil {
		return "", -1, err
	}
	_, err = wgroup.Write([]byte(group))
	if err != nil {
		return "", -1, err
	}

	// open the file to read its content
	file, err := mpfh.Open()
	if err != nil {
		return "", -1, err
	}
	defer file.Close()

	// copy the file content into the form file field
	wfile, err := writer.CreateFormFile("file", mpfh.Filename)
	if err != nil {
		return "", -1, err
	}
	_, err = io.Copy(wfile, file)
	if err != nil {
		return "", -1, err
	}

	// close the writer to finalize the multipart form
	err = writer.Close()
	if err != nil {
		return "", -1, err
	}

	// create a new HTTP request to forward the multipart form data
	req, err := http.NewRequest(UploadFileMethod, s.pathUpload(), &reqb)
	if err != nil {
		return "", -1, err
	}

	// set appropriate header for the request
	req.Header.Set("X-Api-Token", s.AccessToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// send the request using http.DefaultClient
	var res *http.Response
	res, err = s.http.Do(req)
	if err != nil {
		return "", -1, err
	}
	defer res.Body.Close()

	// check the response from the external API
	if res.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("non-OK response code: %d", res.StatusCode)
		return "", res.StatusCode, errors.New(msg)
	}

	// check the response body length
	if res.ContentLength <= 0 {
		err = errors.New("upload success but no response body")
		return "", res.StatusCode, err
	}

	// bind the JSON response to the UploadResponse struct
	resb, err := io.ReadAll(res.Body)
	if err != nil {
		err = errors.New("upload success but failed read the response")
		return "", res.StatusCode, err
	}

	// success, treat response as string
	return string(resb), res.StatusCode, nil
}
