package utils

import "net/http"

func RespondWithHeaders(w http.ResponseWriter, statusCode int, body []byte, cacheHeader string, cacheKey string) {
	w.Header().Set("X-Cache", cacheHeader)
	w.WriteHeader(statusCode)
	w.Write(body)
}
