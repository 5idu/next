package next

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestNewEngine(t *testing.T) {
	engine := NewEngine(nil)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			_, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			cancel()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func initRouter(e *Engine) {
	users := e.PathPrefix("/users").Subrouter()
	users.HandleFunc("", GetUser).Methods(http.MethodGet)
	users.HandleFunc("", AddUser).Methods(http.MethodPost)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "get user")
}

type User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password"`
}

// AddUser 添加用户
func AddUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := ShouldBindJSON(r, &u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Println(u)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "add user")
}
