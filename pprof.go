package next

import (
	"net/http"
	"net/http/pprof"
	"sync"
)

var once sync.Once

// mountPprof mount pprof
func mountPprof(e *Engine) {
	once.Do(func() {
		prefix := e.PathPrefix("/debug/pprof").Subrouter()
		prefix.HandleFunc("/", pprof.Index).Methods(http.MethodGet)
		prefix.HandleFunc("/cmdline", pprof.Cmdline).Methods(http.MethodGet)
		prefix.HandleFunc("/profile", pprof.Profile).Methods(http.MethodGet)
		prefix.HandleFunc("/symbol", pprof.Symbol).Methods(http.MethodPost)
		prefix.HandleFunc("/symbol", pprof.Symbol).Methods(http.MethodGet)
		prefix.HandleFunc("/trace", pprof.Trace).Methods(http.MethodGet)
		prefix.HandleFunc("/allocs", pprof.Handler("allocs").ServeHTTP).Methods(http.MethodGet)
		prefix.HandleFunc("/block", pprof.Handler("block").ServeHTTP).Methods(http.MethodGet)
		prefix.HandleFunc("/goroutine", pprof.Handler("goroutine").ServeHTTP).Methods(http.MethodGet)
		prefix.HandleFunc("/heap", pprof.Handler("heap").ServeHTTP).Methods(http.MethodGet)
		prefix.HandleFunc("/mutex", pprof.Handler("mutex").ServeHTTP).Methods(http.MethodGet)
		prefix.HandleFunc("/threadcreate", pprof.Handler("threadcreate").ServeHTTP).Methods(http.MethodGet)
		return
	})
}
