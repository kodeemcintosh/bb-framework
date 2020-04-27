package main

import (
	"context"
	"crypto/tls"
	// "flag"
	"fmt"
	"os"
	// "io"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

const (
	httpPort          = "127.0.0.1:8080"
	redirectHTTPToTLS = true
)

var m *autocert.Manager

// ServeTLS ... Used to serve the app up in the main function
func (a App) ServeTLS() {
	hostPolicy := func(ctx context.Context, host string) error {
		// Note: change to your real host
		allowedHost := os.Getenv("ALLOWED_HOST")
		if host == allowedHost {
			return nil
		}
		return fmt.Errorf("acme/autocert: only %s host is allowed", allowedHost)
	}

	certDir := "."
	m = &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: hostPolicy,
		Cache:      autocert.DirCache(certDir),
	}

	a.Server.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
	a.Server.Handler = Router()
	a.Server.Addr = ":443"

	go func() {
		fmt.Printf("Starting HTTPS server on %s\n", a.Server.Addr)
		err := a.Server.ListenAndServeTLS("", "")
		if err != nil {
			log.Fatalf("https.ListendAndServeTLS() failed with %s", err)
		}
	}()
}

func (a App) ServeHTTP() {
	r := Router()

	if redirectHTTPToTLS {
		r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			newURI := "https://" + req.Host + req.URL.String()
			http.Redirect(w, req, newURI, http.StatusFound)
		})
	}

	// allow autocert handle Let's Encrypt callbacks over http
	if m != nil {
		a.Server.Handler = m.HTTPHandler(r)
	}

	a.Server.Addr = "localhost:8080"
	a.Server.Handler = r

	go func() {
		fmt.Printf("Starting HTTP server on %s\n", httpPort)
		err := a.Server.ListenAndServe()
		if err != nil {
			log.Fatalf("httpServe.ListenAndServe() failed with %s", err)
		}
	}()
}
