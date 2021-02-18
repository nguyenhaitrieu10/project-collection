package main

import (
	"net/http"
	"router"
	"time"
)

func main() {
	router := router.SetupRouter()
	s := &http.Server{
		Addr:           "0.0.0.0:9000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	// router.Run("0.0.0.0:9000")
}
