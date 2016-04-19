package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/hersshel/hersshel/config"
	"github.com/hersshel/hersshel/router"
	"github.com/hersshel/hersshel/router/middleware"
)

func main() {
	var cfg *config.Config

	cfg = config.LoadConfigFromEnv()
	if cfg.Server != nil && cfg.Server.GOMAXPROCS != nil {
		runtime.GOMAXPROCS(*cfg.Server.GOMAXPROCS)
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	handler := router.Load(
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true),
		middleware.Store(cfg.PostgreSQL),
	)

	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.HTTPPort), handler))
}
