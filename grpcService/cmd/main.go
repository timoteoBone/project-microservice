package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc/reflection"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	_ "github.com/go-sql-driver/mysql"

	"google.golang.org/grpc"

	pb "github.com/timoteoBone/project-microservice/grpcService/pkg/pb"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/user"
)

func main() {

	var db *sql.DB

	db, err := sql.Open("mysql", "root:PewDiePie8!!@tcp(127.0.0.1:3306)/test?parseTime=true")

	if err != nil {
		fmt.Println("sd")
	}

	defer db.Close()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	repo := user.NewSQL(db, logger)
	srv := user.NewService(logger, repo)

	end := user.MakeEndpoint(srv)
	grpcSv := user.NewGrpcServer(end)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":50000")

	if err != nil {
		level.Error(logger).Log("exit", err)
		panic(err)
	}

	defer grpcListener.Close()

	go func() {
		baseServer := grpc.NewServer()
		reflection.Register(baseServer)
		pb.RegisterUserServiceServer(baseServer, grpcSv)
		level.Info(logger).Log("msg", "Server started")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)

}
