package utils

import "net/http"

func RespondWithHeaders(w http.ResponseWriter, r http.Response, body []byte, cacheHeader string, cacheKey string) {
	w.Header().Set("X-Cache", cacheHeader)
	w.WriteHeader(r.StatusCode)
	for k, v := range r.Header {
		w.Header()[k] = v
	}
	w.Write(body)
}
