package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/meatballhat/negroni-logrus"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "success!\n")
	})

	n := negroni.New()

	nl := negronilogrus.NewMiddleware()
	// override the default Before
	nl.Before = customBefore
	// wrap the default after, here replacing dots with underscores in keys
	nl.After = makeNoDotAfter(nl.After)

	n.Use(nl)
	n.UseHandler(r)

	n.Run(":9999")
}

func customBefore(entry *logrus.Entry, _ *http.Request, remoteAddr string) *logrus.Entry {
	return entry.WithFields(logrus.Fields{
		"REMOTE_ADDR": remoteAddr,
		"YELLING":     true,
	})
}

func makeNoDotAfter(after negronilogrus.AfterFunc) negronilogrus.AfterFunc {
	return func(entry *logrus.Entry, res negroni.ResponseWriter, latency time.Duration, name string) *logrus.Entry {
		return negronilogrus.EntryKeysReplace(after(entry, res, latency, name), ".", "_").WithField("ALL_DONE", true)
	}
}
