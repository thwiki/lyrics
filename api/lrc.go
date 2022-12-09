package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/thwiki/lyrics/lyrics"
)

var (
	responseMaxAge               = os.Getenv("RESPONSE_MAX_AGE")
	responseSMaxAge              = os.Getenv("RESPONSE_S_MAX_AGE")
	responseStaleWhileRevalidate = os.Getenv("RESPONSE_STALE_WHILE_REVALIDATE")
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")

	header := w.Header()
	header.Set("Content-Type", "text/plain; charset=utf-8")
	header.Set("Cache-Control", "max-age="+responseMaxAge+", s-maxage="+responseSMaxAge+", stale-while-revalidate="+responseStaleWhileRevalidate+", public")

	text, err := lyrics.GetLyrics(name)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(text))
}

type ErrorResponse struct {
	Error string `json:"error"`
}
