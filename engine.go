package next

import (
	"log"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// Engine .
type Engine struct {
	conf   *EngineConfig
	server atomic.Value // store *http.Server
	lock   sync.RWMutex

	*mux.Router
}

// EngineConfig engine config
type EngineConfig struct {
	Network      string
	Address      string
	Timeout      time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// NewEngine new engine
func NewEngine(conf *EngineConfig) *Engine {
	e := &Engine{Router: mux.NewRouter()}
	if err := e.SetConfig(conf); err != nil {
		panic(err)
	}
	mountPprof(e)
	return e
}

// SetConfig set engine conf
func (e *Engine) SetConfig(conf *EngineConfig) (err error) {
	if conf == nil {
		e.defaultConfig()
		return
	}
	if conf.Timeout <= 0 {
		return errors.New("go-next: config timeout must greater than 0")
	}
	if conf.Network == "" {
		conf.Network = "tcp"
	}
	if conf.Address == "" {
		conf.Address = "127.0.0.1:8080"
	}
	e.lock.Lock()
	e.conf = conf
	e.lock.Unlock()
	return
}

// defaultConfig engine default config
func (e *Engine) defaultConfig() {
	conf := &EngineConfig{
		Network: "tcp",
		Address: "127.0.0.1:8080",
		Timeout: time.Second,
	}
	e.lock.Lock()
	e.conf = conf
	e.lock.Unlock()
}

// Start start engine
func (e *Engine) Start() error {
	conf := e.conf
	l, err := net.Listen(conf.Network, conf.Address)
	if err != nil {
		errors.Wrapf(err, "go-next: listen tcp: %s", conf.Address)
		return err
	}
	log.Printf("go-next: server started on addr: %s\n", conf.Address)
	server := &http.Server{
		ReadTimeout:  time.Duration(conf.ReadTimeout),
		WriteTimeout: time.Duration(conf.WriteTimeout),
	}
	go func() {
		if err := e.Run(server, l); err != nil {
			if errors.Cause(err) == http.ErrServerClosed {
				log.Println("go-next: server closed")
				return
			}
			panic(errors.Wrapf(err, "go-next: engine.Run error(%+v, %+v)", server, l))
		}
	}()

	return nil
}

// Run .
func (e *Engine) Run(server *http.Server, l net.Listener) (err error) {
	server.Handler = e
	e.server.Store(server)
	if err = server.Serve(l); err != nil {
		err = errors.Wrapf(err, "go-next: listen server: %+v/%+v", server, l)
		return
	}
	return
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.Router.ServeHTTP(w, r)
}
