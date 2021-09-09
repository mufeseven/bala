package app

import (
	"bala/app/config"
	"bala/app/log"
	"bala/app/mid"
	"bala/app/router"
	"bala/app/util"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	_ "expvar"
	_ "net/http/pprof"
)

// Start 启动应用
func Start() {
	injector, injectorClearFunc, err := BuildInjector()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	httpClearFunc := initHttp(ctx, injector.Engine, injector.Config)
	httpsClearFunc := initHttps(ctx, injector.Engine, injector.Config)
	pprofClearFunc := initPprof(injector.Config)

	clearFunc := func() {
		injectorClearFunc()
		httpClearFunc()
		httpsClearFunc()
		pprofClearFunc()
	}

	<-util.Signal()
	log.Local().Info("prepare stop the service")
	clearFunc()
	log.Local().Info("service has been stopped")
}

// InitGin 初始化gin引擎
func InitGin(r router.IRouter, config *config.Config) (*gin.Engine, error) {
	gin.SetMode(config.Get().App.Mode)
	e := gin.New()
	// 中间件越早注册调用越晚
	/*if config.IsDev() {
		e.Use(mid.LogResponse)
	}*/
	e.Use(mid.DecodeRequestBody)
	e.Use(mid.LoggerToFile())
	e.Use(mid.Recovery(mid.RecoveryHandler))
	// Router register
	err := r.Register(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func initHttp(ctx context.Context, handler http.Handler, config *config.Config) func() {
	host := config.Get().Http.Host
	port := config.Get().Http.Port
	addr := fmt.Sprintf("%s:%d", host, port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  50 * time.Second,
		WriteTimeout: 100 * time.Second,
		IdleTimeout:  150 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Local().Printf("HTTP server is listening at %s.", addr)

	return func() {
		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			log.Local().Error(err.Error())
		}
	}
}

func initHttps(ctx context.Context, handler http.Handler, config *config.Config) func() {
	host := config.Get().Https.Host
	port := config.Get().Https.Port
	certFile := config.Get().Https.CertFile
	keyFile := config.Get().Https.KeyFile
	addr := fmt.Sprintf("%s:%d", host, port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		if certFile != "" && keyFile != "" {
			return
		}
		err := srv.ListenAndServeTLS(certFile, keyFile)
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Local().Printf("HTTPS server is listening at %s.", addr)

	return func() {
		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			log.Local().Error(err.Error())
		}
	}
}

func initPprof(config *config.Config) func() {
	if !config.Get().Pprof.Enable {
		return func() {}
	}
	host := config.Get().Pprof.Host
	port := config.Get().Pprof.Port
	addr := fmt.Sprintf("%s:%d", host, port)
	srv := &http.Server{Addr: addr}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Local().Printf("PPROF server is listening at %s.", addr)
	return func() {}
}
