package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Aris-haryanto/Best-Way-To-Structuring-Golang-Code/adapters"
	"github.com/Aris-haryanto/Best-Way-To-Structuring-Golang-Code/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gin-gonic/gin"

	pb "github.com/Aris-haryanto/Best-Way-To-Structuring-Golang-Code/proto"
)

func sqlConn() *adapters.SqlDB {
	//define sql connection
	sqlMasterConn := adapters.SqlConnection("root", "root", "localhost", "3308")
	setSqlMasterConn := &adapters.SqlDB{SqlConn: sqlMasterConn}
	return setSqlMasterConn
}

func main() {
	// get config
	getConfig, errConfig := getConfig("./config.yml")
	if errConfig != nil {
		log.Fatalln(errConfig)
	}

	//set config to global variable on service package
	services.InitConfig(getConfig)

	// initial service
	hello := &services.Hello{}

	// set DB management
	sqlConn := sqlConn()
	hello.SetDB(sqlConn)

	// register http server
	srvRest := &services.RestServer{}
	srvRest.RegisterHello(hello)

	// register grpc server
	srvGrpc := &services.GrpcServer{}
	srvGrpc.RegisterHello(hello)

	// ==========================

	rest := runHttpServer("8081", srvRest)
	grpc := runGrpcServer("8082", srvGrpc)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// close everything in here
	// like a connection, channel, or anything
	defer func() {
		cancel()
		sqlConn.CloseSql()
	}()

	grpc.GracefulStop()
	if err := rest.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting!")

}

func addRoutes(r *gin.Engine, srvRest *services.RestServer) {
	r.GET("/hello", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()

		// call service method
		srvRest.SayHello(ctx, c.Writer, c.Request)
	})
}

func runHttpServer(port string, srvRest *services.RestServer) *http.Server {

	// define gin
	r := gin.Default()
	r.Use(gin.Recovery())

	// add routes to gin
	addRoutes(r, srvRest)

	// listen and serve on 0.0.0.0:8080
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv

}

func runGrpcServer(port string, srvGrpc *services.GrpcServer) *grpc.Server {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterGrpcServerServer(srv, srvGrpc)

	reflection.Register(srv)

	log.Printf("server listening at %v", listen.Addr())

	go func() {
		if err := srv.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return srv
}
