package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"news/cmd/controllers"
	"news/pkg/config"
	"news/pkg/db"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	_, err = db.GetDB()
	if err != nil {
		panic(err)
	}
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/news", controllers.News)

	go func() {
		err = http.ListenAndServe(os.Getenv("app_Port"), r)
		if err != nil {
			fmt.Printf("failed ListenAndServe: %v", err)
			panic(err)
		}
	}()
	log.Print("NewsApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("NewsApp Shutting Down")

}
