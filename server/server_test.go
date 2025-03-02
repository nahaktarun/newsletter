package server_test

import (
	"canvas/integrationtest"
	"testing"
	"github.com/matryer/is"
	"net/http"
)

func TestServer_Start(t *testing.T) {
	integrationtest.SkipIfShort(t)

	t.Run("start the server and listens for requests", func(t *testing.T){
		is := is.New(t)

		cleanup := integrationtest.CreateServer()
		defer cleanup()

		resp, err := http.Get("http://localhost:8081/")
		is.NoErr(err)
		is.Equal(http.StatusNotFound, resp.StatusCode)
	})
}
