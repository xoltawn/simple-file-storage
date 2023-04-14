package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/uptrace/bunrouter"
	_http "github.com/xoltawn/simple-file-storage/delivery/http"
	_grpc "github.com/xoltawn/simple-file-storage/repository/grpc"
	"github.com/xoltawn/simple-file-storage/repository/grpc/filepb"
	"github.com/xoltawn/simple-file-storage/usecase"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cc, err := grpc.Dial(os.Getenv("FILE_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	state := cc.GetState()
	log.Println(state.String())
	log.Println("Connected to the File service...")

	client := filepb.NewFileServiceClient(cc)

	fileRepo := _grpc.NewFileGRPCRepository(client)
	fileUsecase := usecase.NewFileUsecase(fileRepo)

	maxUploadSize, err := strconv.Atoi(os.Getenv("MAX_UPLOAD_SIZE"))
	if err != nil {
		log.Fatalln(err)
	}

	router := bunrouter.New()
	_http.NewFileHTTPHandler(router, fileUsecase, int64(maxUploadSize))

	listenAddr := os.Getenv("SERVER_ADDRESS")
	server := &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	log.Println("Server is starting... ")
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("Server is shutting down... ")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	log.Println("Server is ready to handle requests at ", listenAddr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	log.Println("Server stopped")
}
