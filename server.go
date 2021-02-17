package main

import (
	"net/http"

	"muraldevice/artifact"

	"github.com/spf13/afero"
)

func main() {
	artifactHandler := artifact.New(afero.NewOsFs())

	http.HandleFunc("/artifact", artifactHandler.HandleArtifacts)
	http.HandleFunc("/", getterPoster)

	http.ListenAndServe(":8090", nil)
}

func getterPoster(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "artifact.html")
}