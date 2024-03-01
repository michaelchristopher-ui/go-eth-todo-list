package main

import (
	"fmt"
	"net/http"
	"time"

	httpapi "go-eth-todo-list/api/http"

	"github.com/labstack/echo"
)

type Config struct {
	ReadTimeout  int
	WriteTimeout int
}

func main() {
	e := echo.New()
	//todo Zap SugaredLogger when start server.
	config := Config{}

	//Register APIs
	httpapi.CreateApi(e)
	httpapi.ListApi(e)
	httpapi.UpdateApi(e)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", "8008"),
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
	}
	e.StartServer(s)
}
