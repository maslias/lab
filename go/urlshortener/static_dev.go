//+build dev
//go:build dev
// +build dev

package root

import (
	"fmt"
	"net/http"
	"os"
)

func Public() http.Handler {
    // router.Handle("/public/", http.FileServer(http.FS(static)))
    fmt.Println("building static files for dev, hope so")
    return http.StripPrefix("/public", http.FileServerFS(os.DirFS("public")))
}
