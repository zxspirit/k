package build

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 装饰器模式

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)

		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(now))
	}
	return http.HandlerFunc(fn)
}

func HelloWorld(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))

}
func HowAreYou(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I'm fine, thank you!"))

}

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", HelloWorld)
	mux.HandleFunc("GET /howareyou", HowAreYou)
	srv := http.Server{
		Addr:    ":8080",
		Handler: Logger(mux),
	}
	fmt.Println("Server is running on port ", srv.Addr)
	srv.ListenAndServe()
}
