package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"nucleus/models"
	"nucleus/modules/logger"
	"os"
	"strings"
)

func Server(routes *[]models.Route) {

	port := os.Getenv("PORT")
	log.Println("- Proxy server started at:", port)

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger.Log(r)

		if len(*routes) == 0 {
			http.ServeFile(w, r, "modules/proxy/index.html")
			return
		}

		for _, route := range *routes {

			if strings.HasPrefix(r.URL.Path, route.Prefix) && methodAllowed(route.Methods, r.Method) {

				targetURL, err := url.Parse(route.Target)

				if err != nil {
					http.Error(w, "Invalid target URL", http.StatusInternalServerError)
					return
				}

				proxy := httputil.NewSingleHostReverseProxy(targetURL)

				originalPath := r.URL.Path
				r.URL.Path = strings.TrimPrefix(originalPath, route.Prefix)
				r.URL.Path = "/" + strings.TrimPrefix(r.URL.Path, "/")
				r.Host = targetURL.Host

				proxy.ServeHTTP(w, r)
				return
			}
		}

		http.NotFound(w, r)

	})

	log.Println(http.ListenAndServe(":"+port, handler))

}

func methodAllowed(methods []string, method string) bool {
	for _, m := range methods {
		if strings.EqualFold(m, method) {
			return true
		}
	}
	return false
}
