package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
    t.Run("compares speed of the servers, returns the url of the faster one", func(t *testing.T) {
        slowServer := makeDelayedServer(20 * time.Millisecond)
        fastServer := makeDelayedServer(0 * time.Millisecond) 

        defer slowServer.Close()
        defer fastServer.Close()

        slowURL := slowServer.URL
        fastURL := fastServer.URL

        want := fastURL
        got,_ := Racer(slowURL, fastURL)

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
        
        slowServer.Close()
        fastServer.Close()
    })
    t.Run("returns an error if the server doesn't return within 10s", func(t *testing.T){
        server := makeDelayedServer(15 * time.Millisecond)
        defer server.Close()
        
        _,err := ConfigurableRacer(server.URL, server.URL, 10 * time.Millisecond)

        if err == nil {
            t.Errorf("expected an error but didn't get one")
        }
    })
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}
