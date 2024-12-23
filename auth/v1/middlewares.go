package v1

import "net/http"

const (
	DefaultMockUserID   string = "00000000"
	DefaultMockUserName string = "Mock User"
)

type AuthPass struct {
	UserID   string
	UserName string
}

func (a AuthPass) ServeHTTP(_ http.ResponseWriter, r *http.Request) {
	r.Header.Set(HTTPHeaderUserID, a.UserID)
	r.Header.Set(HTTPHeaderUserName, a.UserName)
}

// DefaultMockAuthPass construct AuthPass with default mock data.
func DefaultMockAuthPass() AuthPass {
	return AuthPass{
		UserID:   DefaultMockUserID,
		UserName: DefaultMockUserName,
	}
}

// InjectMockAuthPass simulates production behavior by injecting mock
// authentication headers into requests. Indeployment, these headers are added
// by the Auth API middleware upon successful checks.
//
// Default mock data is used to feed the headers.
func InjectMockAuthPass(next http.Handler) http.Handler {
	return InjectMockAuthPassWith(next, DefaultMockAuthPass())
}

// InjectMockAuthPassWith simulates production behavior by injecting mock
// authentication headers into requests. Indeployment, these headers are added
// by the Auth API middleware upon successful checks.
//
// The provided data is used to feed the headers.
func InjectMockAuthPassWith(next http.Handler, data AuthPass) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data.ServeHTTP(w, r)
		next.ServeHTTP(w, r)
	})
}
