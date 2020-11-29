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

	loglevel := logrus.InfoLevel

	if quiet {
		loglevel = logrus.ErrorLevel
	}

	n := negroni.New()
	n.Use(negronilogrus.NewCustomMiddleware(loglevel, &logrus.JSONFormatter{}, "web"))
	n.UseHandler(r)

	n.Run(":9999")
}
