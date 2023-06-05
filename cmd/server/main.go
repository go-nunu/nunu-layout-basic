package main

import (
	"fmt"
	"github.com/go-nunu/nunu-layout-base/cmd/server/wire"
	"github.com/go-nunu/nunu-layout-base/pkg/config"
	"github.com/go-nunu/nunu-layout-base/pkg/http"
	"github.com/go-nunu/nunu-layout-base/pkg/log"
	"go.uber.org/zap"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	app, cleanup, err := wire.NewApp(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
}
