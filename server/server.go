package server

import (
	"context"
	"fmt"
	"github.com/Ad3bay0c/ContactSaver/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	DB	db.DB
	ErrorLog	*log.Logger
	InfoLog	*log.Logger
}

func (s *Server) ApplicationSetup() {
	file, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("Error: %v", err.Error())
	}
	ErrorLog := log.New(os.Stdout, "Error:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog.SetOutput(file)

	InfoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLog.SetOutput(file)

	mongoDb := &db.MongoDB{}
	mongoDb.Init()
	s.DB = mongoDb
	s.InfoLog = InfoLog
	s.ErrorLog = ErrorLog
}

func (s *Server) Start() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if port == ":" {
		port += "2500"
	}
	router := gin.Default()

	s.ApplicationSetup()
	s.Routes(router)
	server := &http.Server{
		Handler: router,
		Addr: port,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen %s\n", err)
		}
	}()

	log.Printf("Server Started at localhost%s", port)

	quit := make(chan os.Signal)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// The code always stays until it gets an interrupt or terminate signals, before it proceed to the next line
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Serever Forced to Exit: %v", err.Error())
	}
	log.Println("Server Shut Down")
}
