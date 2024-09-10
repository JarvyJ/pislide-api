package main

import (
	"net/http"
	"os"

	"github.com/JarvyJ/pislide-api/internal/photos"
	"github.com/JarvyJ/pislide-api/internal/pislideservice"
	"github.com/JarvyJ/pislide-api/internal/slideshow"
	"github.com/JarvyJ/pislide-api/internal/util"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

func main() {
	// Create a new router & API
	router := chi.NewMux()
	config := huma.DefaultConfig("PiSlide API", "0.0.0")
	config.DocsPath = ""
	api := humachi.New(router, config)

	setupApiRef(router)

	// register the world
	pislideservice.RegisterPiSlideshowService(&api)
	slideshow.RegisterSlideshow(&api)
	photos.RegisterPhotos(&api)

	// route static requests
	fs := http.FileServer(http.Dir(getFrontendDir()))
	router.Handle("/*", fs)

	imageServer := http.FileServer(http.Dir("./slideshows"))
	router.Handle("GET /photos/*", http.StripPrefix("/photos", imageServer))

	// Start the server!
	http.ListenAndServe("0.0.0.0:"+util.GetPort(), router)
}

func setupApiRef(router *chi.Mux) {

	router.Get("/api-ref", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="referrer" content="same-origin" />
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
    <title>Docs Example reference</title>
    <!-- Embed elements Elements via Web Component -->
    <link href="/stoplight-styles.min.css" rel="stylesheet" />
    <script src="/stoplight-web-components.min.js"></script>
  </head>
  <body style="height: 100vh;">
    <elements-api
      apiDescriptionUrl="/openapi.yaml"
      router="hash"
      layout="sidebar"
      tryItCredentialsPolicy="same-origin"
    />
  </body>
</html>`))
	})
}

func isRealDir(dir string) bool {
	fileInfo, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}
	if fileInfo.IsDir() {
		return true
	}
	return false
}

func getFrontendDir() string {
	devFrontendLocation := "./frontend/build"
	if isRealDir(devFrontendLocation) {
		return devFrontendLocation
	}

	prodFrontendLocation := "/opt/pislide/frontend"
	if isRealDir(prodFrontendLocation) {
		return prodFrontendLocation
	}

	panic("Could not find the frontend anywhere. Going to panic now")
}
