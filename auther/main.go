package main

import (
	"auther/api"
	"auther/api/grpc"
	autherhttp "auther/api/http"
	"auther/models"
	"auther/repository/sqlite"
	"os"
	"os/signal"
	"sync"

	"github.com/hashicorp/go-hclog"
)

func main() {
	l := hclog.Default()

	repo, err := sqlite.InitDatabase(l, "auther.db")
	if err != nil {
		l.Error("Cannot open connection to DB", err)
		os.Exit(1)
	}
	defer repo.Close()

	jw := models.NewJwtSigner("My very very secret key. Did you know I like trains?", "nocloudstream/auther", 10)

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt)
	//signal.Notify(c, os.Kill)
	//signal.Notify(c, syscall.SIGTERM)

	bp := api.NewBackpack(repo, jw, l, models.NewSigtermStruct(c, &sync.WaitGroup{}))

	bp.Sigterm.WaitGroup.Add(1)
	go autherhttp.InitHTTPServer(":8080", bp)

	bp.Sigterm.WaitGroup.Add(1)
	go grpc.InitGRPCServer(":8081", bp)

	sig := <-c
	l.Info("Got signal", sig.String())
	c <- sig
	bp.Sigterm.WaitGroup.Wait()
	return
}
