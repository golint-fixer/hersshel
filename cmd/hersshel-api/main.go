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

// TODO: find this function a real home
func feedCron(url string, ticker *time.Ticker) {
	client := DefaultPooledClient()
	for t := range ticker.C {
		logrus.WithField("at", t).Info("start refreshing feed")
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			logrus.WithError(err).Info("error creating request")
		}
		res, err := client.Do(req)
		if err != nil {
			logrus.WithError(err).Info("error refreshing feed")
			continue
		}
		logrus.WithField("response", res).Info("feed refreshed")
	}
}

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
		middleware.Engine(),
	)

	refreshURL := fmt.Sprintf("http://localhost:%d/v1/items", cfg.Server.HTTPPort)
	go feedCron(refreshURL, time.NewTicker(15*time.Minute))
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.HTTPPort), handler))
}
