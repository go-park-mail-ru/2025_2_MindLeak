package swagger

import (
	"net/http"
	"path/filepath"

	httpSwagger "github.com/swaggo/http-swagger"
)

func SwaggerHandler(w http.ResponseWriter, r *http.Request) {
	if filepath.Base(r.URL.Path) == "swagger.json" {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "swagger/swagger.json")
		return
	}

	httpSwagger.Handler(
		httpSwagger.URL("/api/swagger/swagger.json"),
	).ServeHTTP(w, r)
}
