package v1

import "net/http"

// GetUserID returns the user id from the request header. If not present,
// returns an empty string ("").
func GetUserID(header http.Header) string {
	return header.Get(HTTPHeaderUserID)
}

// GetUserName returns the user name from the request header. If not present,
// returns an empty string ("").
func GetUserName(header http.Header) string {
	return header.Get(HTTPHeaderUserName)
}

// MustGetUserID returns existing user id from the request header or panic.
func MustGetUserID(header http.Header) string {
	if s := GetUserID(header); s != "" {
		return s
	}
	msg := fmtMissingHeader(HTTPHeaderUserID)
	panic(msg)
}

// MustGetUserName returns existing user name from the request header or panic.
func MustGetUserName(header http.Header) string {
	if s := GetUserName(header); s != "" {
		return s
	}
	msg := fmtMissingHeader(HTTPHeaderUserName)
	panic(msg)
}
