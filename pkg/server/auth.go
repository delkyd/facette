package server

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func (server *Server) handleAuth(writer http.ResponseWriter, request *http.Request) bool {
	authorization := request.Header.Get("Authorization")

	if strings.HasPrefix(authorization, "Basic ") {
		data, err := base64.StdEncoding.DecodeString(authorization[6:])
		if err != nil {
			return false
		}

		chunks := strings.Split(string(data), ":")
		if len(chunks) != 2 {
			return false
		}

		if server.AuthHandler.Authenticate(chunks[0], chunks[1]) {
			return true
		}
	}

	writer.Header().Add("WWW-Authenticate", "Basic realm=\"Authorization Required\"")

	return false
}