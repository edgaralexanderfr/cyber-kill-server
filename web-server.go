package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/edgaralexanderfr/cyber-kill-server/pkg/config"
)

func main() {
	configFactory := &config.ConfigFactory{}
	config := configFactory.NewConfig().GetConfig(true)

	fs := http.FileServer(http.Dir(config.Server.Web.Dir))
	http.Handle("/", fs)

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Web.Port), nil)

	if err != nil {
		log.Fatal(err)
	}
}
