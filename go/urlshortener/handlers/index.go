package handlers

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/maslias/urlshortener/views"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/", ErrorBridge(h.handleIndex))
	router.HandleFunc("POST /shortenurl", ErrorBridge(h.handleShortenUrl))
	router.HandleFunc("/{shortUrl}", ErrorBridge(h.handleShortUrl))
}

func (h *IndexHandler) handleIndex(w http.ResponseWriter, r *http.Request) error {
	return views.MakeIndex().Render(r.Context(), w)
}

var urlStore = make(map[string]string)

func (h *IndexHandler) handleShortenUrl(w http.ResponseWriter, r *http.Request) error {
	toShortenUrl := r.FormValue("to-shorten-url")
	shortUrl := fmt.Sprintf("%x", md5.Sum([]byte(toShortenUrl)))[:5]

	urlStore[shortUrl] = toShortenUrl
	return views.MakeShortenUrl( r.Host +"/"+ shortUrl).Render(r.Context(), w)
}

func (h *IndexHandler) handleShortUrl(w http.ResponseWriter, r *http.Request) error {
	shortUrl := r.PathValue("shortUrl")

	url, ok := urlStore[shortUrl]
	if !ok {
        http.NotFound(w, r)
		return fmt.Errorf("url does not exist")
	} else {
        http.Redirect(w, r, url, http.StatusMovedPermanently)
    }

	return nil
}
