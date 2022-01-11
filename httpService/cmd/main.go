package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/timoteoBone/final-project-microservice/http-service/endpoints"
	"github.com/timoteoBone/final-project-microservice/http-service/repository"
	"github.com/timoteoBone/final-project-microservice/http-service/service"
	tr "github.com/timoteoBone/final-project-microservice/http-service/transport"
	"google.golang.org/grpc"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "httpService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	var (
		httpAddr = flag.String("http.addr", ":8000", "HTTP address")
	)
	var (
		grpcServerAddress = flag.String("addr", "localhost:50000", "grpcSvAddres")
	)

	flag.Parse()

	var err error
	var grpcServerConnection *grpc.ClientConn
	{
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())
		grpcServerConnection, err = grpc.Dial(*grpcServerAddress, opts...)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	if err != nil {
		level.Error(logger).Log(err)
	}

	repo := repository.NewgRPClient(logger, grpcServerConnection)

	srvc := service.NewService(repo, logger)

	endpoint := endpoints.MakeEndpoints(srvc)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		httpHandler := tr.NewHTTPSrv(*endpoint)
		level.Info(logger).Log("Listening to", httpAddr)
		errs <- http.ListenAndServe(*httpAddr, httpHandler)
	}()

	level.Error(logger).Log("exit", <-errs)

}
