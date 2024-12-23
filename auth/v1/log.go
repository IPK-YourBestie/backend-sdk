package v1

import "fmt"

const (
	logPrefix string = "sdk auth(v1)"
	logFormat string = "%s: '%s' is not present in the header"
)

// fmtMissingHeader format the log message for missing header.
func fmtMissingHeader(key string) string {
	return fmt.Sprintf(logFormat, logPrefix, key)
}
