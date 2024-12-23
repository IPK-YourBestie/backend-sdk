package v1

import (
	"errors"
	"fmt"
	"net/http"
)

// Delete deletes a file from the given URL.
func (s *service) Delete(url string) (int, error) {
	// Create a new HTTP DELETE request
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return -1, err
	}

	// Set appropriate headers for the request
	req.Header.Set("X-Api-Token", s.AccessToken)

	// Send the request using http.DefaultClient
	var res *http.Response
	res, err = s.http.Do(req)
	if err != nil {
		return -1, err
	}
	res.Body.Close()

	// Check the response from the external API
	if res.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("non-OK response code: %d", res.StatusCode)
		return res.StatusCode, errors.New(msg)
	}

	return res.StatusCode, nil
}

// DeleteFd deletes a file from the specified parameters.
func (s *service) DeleteFd(appId, group, filename string) (int, error) {
	url := s.pathDeleteFd(appId, group, filename)
	return s.Delete(url)
}
