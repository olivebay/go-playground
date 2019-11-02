package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compare the speed of the servers, return the fastest one", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error, but got one")
		}

		if got != want {
			t.Errorf("got %q, but want %q", got, want)
		}

	})

	t.Run("returns a timeout if the response takes more than 10s", func(t *testing.T) {

		server := makeDelayedServer(21 * time.Millisecond)

		serverURL := server.URL

		defer server.Close()

		_, err := ConfigurableRacer(serverURL, serverURL, 19*time.Millisecond)

		if err == nil {
			t.Fatalf("expected an error but didin't get one")
		}

	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}
