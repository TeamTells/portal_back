package network

import (
	"net/http"
	"strings"
)

func GetAccessTokenFromHeader(r *http.Request) string {
	tokenValue := r.Header.Get(Access)
	token, _ := strings.CutPrefix(tokenValue, Barier)
	return token
}
