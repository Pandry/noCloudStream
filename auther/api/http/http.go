package http

import (
	"auther/api"
	"auther/api/http/handlers"
	"context"
	"net/http"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

const pathPrefix = "/api/auth"

func InitHTTPServer(socket string, bp api.Backpack) {
	logger := bp.Logger
	st := bp.Sigterm
	defer func() { logger.Info("HTTP Shutting down"); bp.Sigterm.WaitGroup.Done() }()

	router := mux.NewRouter()

	backpack := handlers.NewHttpBackpack(bp)

	pr := router.Methods(http.MethodPost).Subrouter()
	pr.HandleFunc(pathPrefix+"/signup", backpack.HandleUserSignup)
	pr.HandleFunc(pathPrefix+"/login", backpack.HandleUserLogin)
	pr.HandleFunc(pathPrefix+"/renew", backpack.HandleRenew)

	gr := router.Methods(http.MethodGet).Subrouter()
	gr.Use(backpack.NeedsAuthentication)
	gr.HandleFunc(pathPrefix+"/loggedin", backpack.HandleInfo)

	pr.Use(backpack.ValidateUser)
	pr.Use(backpack.SanitizeUser)

	//TODO: fix this - ATM used only for debug purposes
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:8080", "http://localhost:3000", "*"}))

	// create a new server
	s := http.Server{
		Addr:         socket,                                                // configure the bind address
		Handler:      ch(router),                                            // set the default handler
		ErrorLog:     logger.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  2 * time.Second,                                       // max time to read request from the client
		WriteTimeout: 2 * time.Second,                                       // max time to write response to the client
		IdleTimeout:  5 * time.Second,                                       // max time for connections using TCP Keep-Alive
	}

	go func() {
		logger.Info("HTTP server started - server listening at " + socket)

		err := s.ListenAndServe()
		if err != http.ErrServerClosed {
			logger.Error("Error starting HTTP server", "error", err)
		}
	}()

	sig := <-st.KillChannel
	st.KillChannel <- sig
	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), st.TimeOut)
	s.Shutdown(ctx)
}
