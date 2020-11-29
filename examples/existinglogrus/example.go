package main

import (
	"flag"
	"fmt"
	"net/http"

	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func main() {
	quiet := false

	flag.BoolVar(&quiet, "quiet", quiet, "if set, hide messages from the logger")
	flag.Parse()
	r := http.NewServeMux()
	r.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "success!\n")
	})

	log := logrus.New()
	log.Level = logrus.InfoLevel

	n := negroni.New()
	n.Use(negronilogrus.NewMiddlewareFromLogger(log, "web"))
	n.UseHandler(r)

	n.Run(":9999")
}
