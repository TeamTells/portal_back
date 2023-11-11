package network

import (
	"net/http"
	"strconv"
	"strings"
)

func GetAccessTokenFromHeader(r *http.Request) string {
	tokenValue := r.Header.Get(AccessHeader)
	token, _ := strings.CutPrefix(tokenValue, Barier)
	return token
}

func GetCompanyIdFromHeader(r *http.Request) (int, error) {
	return strconv.Atoi(r.Header.Get(CompanyIdHeader))
}

func GetUserIdFromHeader(r *http.Request) (int, error) {
	return strconv.Atoi(r.Header.Get(UserIdHeader))
}
