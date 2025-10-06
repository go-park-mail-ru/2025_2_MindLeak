package swagger

import (
	"net/http"
	"path/filepath"

	httpSwagger "github.com/swaggo/http-swagger"
)

func SwaggerHandler(w http.ResponseWriter, r *http.Request) {
	if filepath.Base(r.URL.Path) == "swagger.json" {
		http.ServeFile(w, r, "/api/swagger/swagger.json")
		return
	}

	httpSwagger.Handler(
		httpSwagger.URL("/swagger/swagger.json"),
	).ServeHTTP(w, r)
}
