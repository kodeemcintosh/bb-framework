package main

import (
	"flag"
	"net/http"
	"os"
	"time"
	// "fmt"
	// "log"
)

var (
	flgProduction          = false
	flgRedirectHTTPToHTTPS = false
)

var app App

func main() {
	parseFlags()

	// TODO: remove and replace with env variables in dockerfile or yaml
	os.Setenv("ALLOWED_HOST", "localhost")
	os.Setenv("PG_PORT", ":2345")
	os.Setenv("PG_HOST", "localhost")
	os.Setenv("PG_USER", "kodee")
	os.Setenv("PG_PASS", "hunter12")
	os.Setenv("PG_NAME", "bb_db")

	app = App{
		Server: http.Server{
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}

	if flgProduction {
		app.Env = Production

		app.ServeTLS()
		app.Initialize()
		defer app.DB.Close()
	}

	app.Env = Develoment
	app.Initialize()
	defer app.DB.Close()

	app.ServeHTTP()
}

func parseFlags() {
	flag.BoolVar(&flgProduction, "production", false, "if true, we start HTTPS server")
	flag.BoolVar(&flgRedirectHTTPToHTTPS, "redirect-to-https", false, "if true, we redirect HTTP to HTTPS")
	flag.Parse()
}
